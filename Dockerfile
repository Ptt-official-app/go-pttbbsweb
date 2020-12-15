FROM golang:1.14-buster

ENV GOROOT=/usr/local/go
ENV PATH=${PATH}:/usr/local/go/bin

# go-pttbbs
COPY . /srv/go-openbbsmiddleware

WORKDIR /srv/go-openbbsmiddleware
RUN mkdir -p /etc/go-openbbsmiddleware && cp 01-config.docker.ini /etc/go-openbbsmiddleware/production.ini

WORKDIR /srv/go-openbbsmiddleware
RUN go build

# cmd
WORKDIR /srv/go-openbbsmiddleware
CMD ["/srv/go-openbbsmiddleware/go-openbbsmiddleware", "-ini", "production.ini"]

EXPOSE 3457
