FROM mcr.microsoft.com/devcontainers/base:ubuntu-24.04

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
        curl \
        wget \
        git \
        make \
        jq \
        unzip \
        ca-certificates \
        apt-transport-https \
        gnupg \
        lsb-release \
    && rm -rf /var/lib/apt/lists/*
