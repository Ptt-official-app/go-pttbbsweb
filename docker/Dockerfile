FROM golang:1.22-bullseye

ENV GOROOT=/usr/local/go
ENV PATH=${PATH}:/usr/local/go/bin

ARG GO_PTTBBS_VERSION
ARG TAG=docker

# go-pttbbs
COPY . /srv/go-pttbbsweb

RUN echo "GO_PTTBBS_VERSION: ${GO_PTTBBS_VERSION} TAG: ${TAG}"

WORKDIR /srv/go-pttbbsweb
RUN mkdir -p /etc/go-pttbbsweb && cp docs/config/01-config.docker.ini /etc/go-pttbbsweb/production.ini

WORKDIR /srv/go-pttbbsweb
RUN go mod download
RUN cp /srv/go-pttbbsweb/docs/config/00-config-production.go.template /go/pkg/mod/github.com/\!ptt-official-app/go-pttbbs@${GO_PTTBBS_VERSION}/ptttype/00-config-production.go
RUN go build -ldflags "-X github.com/Ptt-official-app/go-pttbbsweb/types.GIT_VERSION=`git rev-parse --short HEAD` -X github.com/Ptt-official-app/go-pttbbsweb/types.VERSION=`git describe --tags`" -tags ${TAG}

RUN mkdir -p /static

# cmd
WORKDIR /srv/go-pttbbsweb
CMD ["/srv/go-pttbbsweb/go-pttbbsweb", "-ini", "production.ini"]

EXPOSE 3457
