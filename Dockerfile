FROM ubuntu:jammy-20220801 AS base
# Dockerの --mount=type=cache を使用するためaptのパッケージキャッシュを有効化
RUN \
  rm -f /etc/apt/apt.conf.d/docker-clean; \
  echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache

#
# dev: 開発環境用イメージ
#
FROM base AS dev

# OSパッケージの更新とgcc,curl等のインストール
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
  export DEBIAN_FRONTEND=noninteractive && \
  apt-get update && \
  apt-get upgrade -y && \
  apt-get install -y --no-install-recommends \
    curl ca-certificates \
    git unzip build-essential

# goのインストール
ARG GO_VERSION=1.21.0
RUN \
  ARCH="$(uname -m)"; \
  case "$ARCH" in \
    'x86_64') ARCH="amd64" ;; \
    'aarch64') ARCH="arm64" ;; \
  esac; \
  curl -sSL https://dl.google.com/go/go${GO_VERSION}.linux-${ARCH}.tar.gz \
    | tar -xzC /usr/local
ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin

# go generate用ツールのインストール
WORKDIR /app
ADD go.mod ./
RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install golang.org/x/tools/gopls@latest && \
  go install golang.org/x/tools/cmd/goimports@latest
ADD . ./

# NOTE: go buildのVCS timestampのため別ユーザーのgitリポジトリ参照を許可
RUN \
  git config --global --add safe.directory /app

CMD ["go", "run", "./cmd", "serve"]

