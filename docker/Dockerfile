FROM golang:1.19-bullseye

ENV GOROOT=/usr/local/go
ENV PATH=${PATH}:/usr/local/go/bin

ARG TAG=docker

# go-pttbbs
COPY . /srv/go-openbbsmiddleware

WORKDIR /srv/go-openbbsmiddleware
RUN mkdir -p /etc/go-openbbsmiddleware && cp docs/config/01-config.docker.ini /etc/go-openbbsmiddleware/production.ini

WORKDIR /srv/go-openbbsmiddleware
RUN go mod download
RUN cp /srv/go-openbbsmiddleware/docs/config/00-config-production.go.template /go/pkg/mod/github.com/\!ptt-official-app/go-pttbbs*/ptttype/
RUN go build -ldflags "-X github.com/Ptt-official-app/go-openbbsmiddleware/types.GIT_VERSION=`git rev-parse --short HEAD` -X github.com/Ptt-official-app/go-openbbsmiddleware/types.VERSION=`git describe --tags`" -tags ${TAG}

RUN mkdir -p /static

# cmd
WORKDIR /srv/go-openbbsmiddleware
CMD ["/srv/go-openbbsmiddleware/go-openbbsmiddleware", "-ini", "production.ini"]

EXPOSE 3457
