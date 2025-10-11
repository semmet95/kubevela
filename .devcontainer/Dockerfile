ARG GO_VERSION=1.24
ARG KIND_VERSION=v0.30.0
ARG KUBECTL_VERSION=v1.31.12
ARG KIND_NODE_IMAGE=kindest/node:${KUBECTL_VERSION}
ARG KIND_CLUSTER_NAME=vela

FROM golang:${GO_VERSION}-alpine

ARG TARGET_ARCH
RUN if [ ${TARGET_ARCH} != "amd64" || ${TARGET_ARCH} != "arm64"]; then \
        echo "Unsupported architecture: ${TARGET_ARCH}"; \
        exit 1; \
    fi

RUN apk add --no-cache bash curl git iptables ca-certificates && \
    update-ca-certificates

# Install kubectl
RUN curl -L -o /usr/local/bin/kubectl https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/${TARGET_ARCH}/kubectl && \
    chmod +x /usr/local/bin/kubectl && \
    kubectl version --client=true

# Install kind
RUN curl -L -o /usr/local/bin/kind https://kind.sigs.k8s.io/dl/${KIND_VERSION}/kind-linux-${TARGET_ARCH} && \
    chmod +x /usr/local/bin/kind && \
    kind --version