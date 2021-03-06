# syntax = docker/dockerfile:experimental

FROM l.gcr.io/google/ubuntu1804:latest AS builder

ENV DEBIAN_FRONTEND=noninteractive \
	OSXCROSS_PATH=/usr/local/osxcross \
	OSXCROSS_BIN=/usr/local/osxcross/target/bin

ARG MACOS_SDK_VERSION
ARG XCODE_VERSION
ARG XCODE_PRODUCT_BUILD_VERSION

RUN set -eux && \
	apt-get update && \
	apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		\
		git \
		patch \
		python3-dev \
		libssl-dev \
		lzma-dev \
		libxml2-dev \
		bash \
		\
		# for ld64 -random_uuid support
		uuid-dev \
		\
		tar \
		xz-utils \
		\
		gnupg \
		\
		sudo \
		&& \
	rm -rf /var/lib/apt/lists/*

# Import public key required for verifying signature of cmake download
RUN set -eux && \
	gpg --keyserver hkp://keyserver.insect.com --recv 0x2D2CEF1034921684

ARG CMAKE_VERSION_MAJOR
ARG CMAKE_VERSION_MINOR
ARG CMAKE_VERSION_PATCH
ENV CMAKE_VERSION=${CMAKE_VERSION_MAJOR}.${CMAKE_VERSION_MINOR}.${CMAKE_VERSION_PATCH}
RUN set -eux && \
	curl -SL "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/cmake-${CMAKE_VERSION}-Linux-x86_64.tar.gz" | tar xzf - -C /usr/local --strip-components=1

ARG LLVM_VERSION
RUN set -eux && \
	curl -SL https://github.com/llvm/llvm-project/releases/download/llvmorg-${LLVM_VERSION}/clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-18.04.tar.xz | tar xf - -J --strip-components=1 -C /usr

ARG OSXCROSS_COMMIT
RUN set -eux && \
	git clone --depth 1 https://github.com/tpoechtrager/osxcross.git ${OSXCROSS_PATH} && \
	git -C ${OSXCROSS_PATH} checkout -q "${OSXCROSS_COMMIT}" && \
	rm -rf ${OSXCROSS_PATH}/.git

RUN set -eux && \
	curl -o ${OSXCROSS_PATH}/tarballs/MacOSX${MACOS_SDK_VERSION}.sdk.tar.xz -SL https://storage.googleapis.com/darwin-sdk/MacOSX${MACOS_SDK_VERSION}-${XCODE_VERSION}-${XCODE_PRODUCT_BUILD_VERSION}.sdk.tar.xz

WORKDIR $OSXCROSS_PATH
# debug for osxcross
ARG OCDEBUG=0
RUN set -eux && \
	CC=/usr/bin/clang CXX=/usr/bin/clang++ UNATTENDED=1 SDK_VERSION=${MACOS_SDK_VERSION} OSX_VERSION_MIN=${MACOS_SDK_VERSION} MACOSX_DEPLOYMENT_TARGET=${MACOS_SDK_VERSION} JOBS="$(nproc)" OCDEBUG=${OCDEBUG} ./build.sh

ARG LIBTOOL_VERSION
RUN set -eux && \
	mkdir -p "${OSXCROSS_PATH}/tarballs/MacOSX${MACOS_SDK_VERSION}.sdk/usr/" && \
	case ${MACOS_SDK_VERSION} in \
		10.13) MACOS_CODENAME='high_sierra' ;; \
		10.14) MACOS_CODENAME='mojave' ;; \
		10.15) MACOS_CODENAME='catalina' ;; \
		10.16) MACOS_CODENAME='catalina' ;; \
		11.0) MACOS_CODENAME='catalina' ;; \
	esac && \
	\
	curl -fsSL "https://homebrew.bintray.com/bottles/libtool-${LIBTOOL_VERSION}.${MACOS_CODENAME}.bottle.tar.gz" | gzip -dc | \
		tar xf - -C "${OSXCROSS_PATH}/target/SDK/MacOSX${MACOS_SDK_VERSION}.sdk/usr/" --strip-components=2 "libtool/${LIBTOOL_VERSION}/include/" "libtool/${LIBTOOL_VERSION}/lib/"

RUN set -eux && \
	groupadd wheel && \
	groupadd admin && \
	usermod -G wheel root

WORKDIR /

ENV OSXCROSS_BIN=/usr/local/osxcross/target/bin \
        PATH=/usr/local/osxcross/target/bin:$PATH \
	CC=${OSXCROSS_BIN}/o64-clang \
	CXX=${OSXCROSS_BIN}/o64-clang++ \
	MACOSX_DEPLOYMENT_TARGET=${MACOS_SDK_VERSION} \
	OSXCROSS_GCC_NO_STATIC_RUNTIME=1

LABEL org.opencontainers.image.authors="The containerz Authors"
LABEL org.opencontainers.image.url="https://github.com/containerz-dev/xx"
LABEL org.opencontainers.image.documentation="https://github.com/containerz-dev/xx"
LABEL org.opencontainers.image.source="git@github.com:containerz-dev/xx"
LABEL org.opencontainers.image.version="${XCODE_VERSION}-${XCODE_PRODUCT_BUILD_VERSION}"
LABEL org.opencontainers.image.vendor="containerz"
LABEL org.opencontainers.image.licenses="BSD-3-Clause"
LABEL org.opencontainers.image.title="macOS Cross toolchain for Linux and *BSD"
LABEL org.opencontainers.image.description="Linux container based cross-compiling macOS Mach-O executable binary"

CMD ["/bin/bash"]
