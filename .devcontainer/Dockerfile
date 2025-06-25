FROM mcr.microsoft.com/devcontainers/base:noble

RUN apt-get update && apt-get upgrade -y

ARG GO_VERSION=1.24.0
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz | tar -C /usr/local -zx

USER vscode

ENV PATH=${PATH}:/usr/local/go/bin
ENV PATH=${PATH}:/home/vscode/go/bin

RUN go install github.com/766b/go-outliner@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

WORKDIR /home/vscode/github.com/si-arakaki/account-manager
