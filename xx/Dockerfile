# syntax = docker/dockerfile:experimental
FROM launcher.gcr.io/google/ubuntu1804:latest AS builder

ENV DEBIAN_FRONTEND=noninteractive \
	OSXCROSS_PATH=/osxcross

ARG MACOS_SDK_VERSION=10.15
ARG XCODE_VERSION=11.5
ARG XCODE_PRODUCT_BUILD_VERSION=11N605c

RUN set -eux && \
	apt-get update && \
	apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		\
		# cmake \
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
		# bzip2 \
		# cpio \
		# gzip \
		# sed \
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

ARG CMAKE_VERSION=3.17.2
RUN set -eux && \
	mkdir /tmp/cmake-install && \
	cd /tmp/cmake-install	&& \
	\
	curl -o /tmp/cmake-install/cmake-${CMAKE_VERSION}-SHA-256.txt.asc -SL "https://cmake.org/files/v${CMAKE_VERSION%.*}/cmake-${CMAKE_VERSION}-SHA-256.txt.asc" && \
	curl -o /tmp/cmake-install/cmake-${CMAKE_VERSION}-SHA-256.txt -SL "https://cmake.org/files/v${CMAKE_VERSION%.*}/cmake-${CMAKE_VERSION}-SHA-256.txt" && \
	gpg --verbose --verify /tmp/cmake-install/cmake-${CMAKE_VERSION}-SHA-256.txt.asc /tmp/cmake-install/cmake-${CMAKE_VERSION}-SHA-256.txt && \
	\
	curl -o /tmp/cmake-install/cmake-${CMAKE_VERSION}-Linux-x86_64.tar.gz -SL "https://cmake.org/files/v${CMAKE_VERSION%.*}/cmake-${CMAKE_VERSION}-Linux-x86_64.tar.gz" && \
	( grep "cmake-${CMAKE_VERSION}-Linux-x86_64.tar.gz" "cmake-${CMAKE_VERSION}-SHA-256.txt" | sha256sum -c - ) && \
	\
	tar xzf "cmake-${CMAKE_VERSION}-Linux-x86_64.tar.gz" -C /usr/local --strip-components=1 && \
	cd / && \
	rm -rf /tmp/cmake-install

ARG LLVM_VERSION=10.0.0
RUN set -eux && \
	curl -SL https://github.com/llvm/llvm-project/releases/download/llvmorg-${LLVM_VERSION}/clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-18.04.tar.xz | tar xf - -J --strip-components=1 -C /usr

ARG OSXCROSS_COMMIT=master
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

ARG LIBTOOL_VERSION=2.4.6_1
RUN set -eux && \
	mkdir -p "${OSXCROSS_PATH}/tarballs/MacOSX${MACOS_SDK_VERSION}.sdk/usr/" && \
	case ${MACOS_SDK_VERSION} in \
		10.13) MACOS_CODENAME='high_sierra' ;; \
		10.14) MACOS_CODENAME='mojave' ;; \
		10.15) MACOS_CODENAME='catalina' ;; \
	esac && \
	\
	curl -fsSL "https://homebrew.bintray.com/bottles/libtool-${LIBTOOL_VERSION}.${MACOS_CODENAME}.bottle.tar.gz" | gzip -dc | tar xf - -C "${OSXCROSS_PATH}/target/SDK/MacOSX${MACOS_SDK_VERSION}.sdk/usr/" \
			--strip-components=2 \
			"libtool/${LIBTOOL_VERSION}/include/" \
			"libtool/${LIBTOOL_VERSION}/lib/"

RUN set -eux && \
	groupadd wheel && \
	groupadd admin && \
	usermod -G wheel root

WORKDIR /

ENV PATH=${OSXCROSS_PATH}/target/bin:$PATH \
	MACOSX_DEPLOYMENT_TARGET=${MACOS_SDK_VERSION} \
	OSXCROSS_GCC_NO_STATIC_RUNTIME=1

CMD ["/bin/bash"]
