FROM mcr.microsoft.com/vscode/devcontainers/go:1.17

RUN apt-get update

RUN yes | apt-get install imagemagick