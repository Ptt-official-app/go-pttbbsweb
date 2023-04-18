# go-openbbsmiddleware

[![Go](https://github.com/Ptt-official-app/go-openbbsmiddleware/actions/workflows/go.yml/badge.svg)](https://github.com/Ptt-official-app/go-openbbsmiddleware/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/Ptt-official-app/go-openbbsmiddleware?status.svg)](https://pkg.go.dev/github.com/Ptt-official-app/go-openbbsmiddleware?tab=doc)
[![codecov](https://codecov.io/gh/Ptt-official-app/go-openbbsmiddleware/branch/main/graph/badge.svg)](https://codecov.io/gh/Ptt-official-app/go-openbbsmiddleware)

## README Translation

* [English](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/README.en.md)
* [正體中文](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/README.zh-TW.md)

## Overview

go implementation of [openbbs-middleware](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96).

With [Ptt-official-app pttbbs](https://github.com/ptt-official-app/go-pttbbs), go-openbbsmiddleware intends to be the web-based bbs.

## Demo Site

* [dev](https://www.devptt.dev)
* [term (PttChrome)](https://term.devptt.dev)

## Getting Started

You can try the [1-script setup by @tingyuchang](https://github.com/tingyuchang/demo-bbs-docker)

You can start with the [swagger api](https://doc.devptt.dev) and try the api.

You can copy the curl command from the link if you encounter CORS issue.

You can go to [https://term.devptt.dev](https://term.devptt.dev) and check how the api affects the existing pttbbs in [www.devptt.dev](https://www.devptt.dev).

The test data (/home/bbs, adopted from [OCF PttID Data](https://ocf.tw/p/pttid/)) can be accessed [here](https://drive.google.com/file/d/1lHuqOYpPDmKayYAaf3UIiLRV1wCjF6bc/view?usp=sharing).
Please setup the following config in pttbbs.conf to use the test data:

```sh
MAX_USERS = 200000 /* 最高註冊人數 */
MAX_BOARD = 8192 /* 最大開板個數 */
```

## Coding Convention

We use the following libraries for coding convention:

* [gotests](https://github.com/cweill/gotests) for test-generation
* [gofumpt](https://github.com/mvdan/gofumpt) for formatting

## docker-compose

You can do the following to start with docker-compose:

* copy `docs/etc/` to some etc directory (ex: `/etc/go-openbbsmiddleware`).
* copy `docs/config/01-config.docker.ini` to the etc directory as production.ini (ex: `cp 01-config.docker.ini /etc/go-openbbsmiddleware/production.ini`).
* copy `docker/docker_compose.env.template` to `docker/docker_compose.env` and modify the settings.
* `./scripts/docker_initbbs.sh [BBSHOME] pttofficialapps/go-pttbbs:latest`
* `docker-compose --env-file docker/docker_compose.env -f docker/docker-compose.yaml up -d`
* register at `http://localhost:3457/account/register`
* login at `http://localhost:3457/account/login`
* `telnet localhost 8888` and use the account that you registered.
* register SYSOP and guest.

## Discussing / Reviewing / Questioning the code

Besides creating issues, you can do the following
to discuss / review / question the code:

* `git clone` the repo
* create a review-[topic] branch
* commenting at the specific code-region.
* pull-request
* start discussion.
* close the pr with comments with only the link of the pr in the code-base.

## Develop

You can start developing by forking this repository.

## Unit-Test

You can do unit-test with:

* `./scripts/test.sh`

You can check coverage with:

* `./scripts/coverage.sh`

## Swagger

You can run swagger with:

* setup python virtualenv.
* `cd apidoc; pip install . && pip uninstall apidoc -y && python setup.py develop; cd ..`
* `./scripts/swagger.sh [host]`
* go to `http://localhost:5000`


## Schema definition

* `https://github.com/Ptt-official-app/go-openbbsmiddleware/tree/main/schema`

## Repository Naming

The reason why this repo is called go-openbbsmiddleware is because previously the [.NET ASP](https://github.com/Ptt-official-app/AspCoreOpenBBSMiddleware) developers envisioned that the scope of this [middleware](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96) can include other versions of bbs (Maple/中山之島). The naming of this repo followed the naming convention at that time.
