ARG GO_VERSION=1.24
ARG KUBECTL_VERSION=v1.31.12
ARG GOLANGCILINT_VERSION=v1.60.3

FROM golang:${GO_VERSION}-bookworm

ARG TARGET_ARCH=$(dpkg --print-architecture)

RUN apt-get update && apt-get install -y \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCILINT_VERSION}

# install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -sL https://dl.k8s.io/release/stable.txt)/bin/linux/${TARGET_ARCH}/kubectl" \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl
