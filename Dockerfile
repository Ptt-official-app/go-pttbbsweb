FROM golang:1.16-buster

ENV GOROOT=/usr/local/go
ENV PATH=${PATH}:/usr/local/go/bin

# go-pttbbs
COPY . /srv/go-openbbsmiddleware

WORKDIR /srv/go-openbbsmiddleware
RUN mkdir -p /etc/go-openbbsmiddleware && cp 01-config.docker.ini /etc/go-openbbsmiddleware/production.ini

WORKDIR /srv/go-openbbsmiddleware
RUN go build -ldflags "-X github.com/Ptt-official-app/go-openbbsmiddleware/types.GIT_VERSION=`git rev-parse --short HEAD` -X github.com/Ptt-official-app/go-openbbsmiddleware/types.VERSION=`git describe --tags`"

RUN mkdir -p /static

# cmd
WORKDIR /srv/go-openbbsmiddleware
CMD ["/srv/go-openbbsmiddleware/go-openbbsmiddleware", "-ini", "production.ini"]

EXPOSE 3457
