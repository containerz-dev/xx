# syntax = docker/dockerfile:experimental
ARG XX_VERSION
FROM gcr.io/containerz/xx/xx:${XX_VERSION} AS builder

ENV GOPATH=/go \
	PATH=/go/bin:/usr/local/go/bin:$PATH \
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=1 \
	\
	GOROOT_BOOTSTRAP=/go1.4

RUN set -eux && \
	curl -fsSL "https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz" -o go1.4-bootstrap.tar.gz && \
	echo "f4ff5b5eb3a3cae1c993723f3eab519c5bae18866b5e5f96fe1102f0cb5c3e52  go1.4-bootstrap.tar.gz" | sha256sum -c - && \
	tar -C / -xzf go1.4-bootstrap.tar.gz && \
	rm go1.4-bootstrap.tar.gz && \
	mv /go "${GOROOT_BOOTSTRAP}" && \
	\
	cd ${GOROOT_BOOTSTRAP}/src && \
	CC=/usr/bin/clang CXX=/usr/bin/clang++ CGO_ENABLED=0 ./make.bash -a

ARG GOLANG_VERSION
RUN set -eux && \
	curl -fsSL "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" -o golang.tar.gz && \
	tar -C /usr/local -xzf golang.tar.gz && \
	rm golang.tar.gz && \
	cd /usr/local/go/src && \
	\
	CC=/usr/bin/clang CXX=/usr/bin/clang++ CC_FOR_TARGET=${OSXCROSS_PATH}/target/bin/o64-clang CXX_FOR_TARGET=${OSXCROSS_PATH}/target/bin/o64-clang++ \
		./make.bash -a=true -d=true -v=1 -no-banner && \
	\
	rm -rf "${GOROOT_BOOTSTRAP}" && \
	mkdir -p "${GOPATH}/src" "${GOPATH}/bin" && chmod -R 777 "${GOPATH}"

WORKDIR "${GOPATH}"

LABEL org.opencontainers.image.authors="The containerz Authors"
LABEL org.opencontainers.image.url="https://github.com/containerz-dev/golang"
LABEL org.opencontainers.image.documentation="https://github.com/containerz-dev/golang"
LABEL org.opencontainers.image.source="git@github.com:containerz-dev/golang"
LABEL org.opencontainers.image.version="${GOLANG_VERSION}-${XX_VERSION}"
LABEL org.opencontainers.image.vendor="containerz"
LABEL org.opencontainers.image.licenses="BSD-3-Clause"
LABEL org.opencontainers.image.title="Go runtime with macOS Cross toolchain for Linux and *BSD"
LABEL org.opencontainers.image.description="Linux container based Go cross-compiling macOS Mach-O executable binary"
