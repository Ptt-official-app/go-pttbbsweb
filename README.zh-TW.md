# go-pttbbsweb

[![Go](https://github.com/Ptt-official-app/go-pttbbsweb/actions/workflows/go.yml/badge.svg)](https://github.com/Ptt-official-app/go-pttbbsweb/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/Ptt-official-app/go-pttbbsweb?status.svg)](https://pkg.go.dev/github.com/Ptt-official-app/go-pttbbsweb?tab=doc)
[![codecov](https://codecov.io/gh/Ptt-official-app/go-pttbbsweb/branch/main/graph/badge.svg)](https://codecov.io/gh/Ptt-official-app/go-pttbbsweb)

## README Translation

* [English](https://github.com/Ptt-official-app/go-pttbbsweb/blob/main/README.en.md)
* [正體中文](https://github.com/Ptt-official-app/go-pttbbsweb/blob/main/README.zh-TW.md)

## 概觀

這裡是使用 golang 來達成 [中介應用層](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96).

與 [Ptt-official-app pttbbs](https://github.com/ptt-official-app/go-pttbbs) 一起成為 web-based BBS.

## Demo Site

* [dev](https://www.devptt.dev)
* [term (PttChrome)](https://term.devptt.dev)

## Getting Started

您可以試著使用 [@tingyuchang 所開發的 1-script setup](https://github.com/tingyuchang/demo-bbs-docker)

您可以到 [swagger api](https://doc.devptt.dev) 並且試著使用 api.

如果您在 swagger 網頁裡遇到 CORS 的問題. 你可以在網頁裡 copy
curl 指令測試.

您可以到 [https://term.devptt.dev](https://term.devptt.dev) 確認 api 如何影響在 [www.devptt.dev](https://www.devptt.dev) 裡既有的 pttbbs.

使用的測試資料 (/home/bbs, 從 [OCF PttID Data](https://ocf.tw/p/pttid/) 更改而來) 在[這裡](https://drive.google.com/file/d/1lHuqOYpPDmKayYAaf3UIiLRV1wCjF6bc/view?usp=sharing).
當使用測試資料時, 請在 pttbbs.conf 做以下的設定:

```sh
MAX_USERS = 200000 /* 最高註冊人數 */
MAX_BOARD = 8192 /* 最大開板個數 */
```

## Coding Convention

我們使用以下 library 幫助 coding convention:

* [gotests](https://github.com/cweill/gotests) for test-generation
* [gofumpt](https://github.com/mvdan/gofumpt) for formatting

## docker-compose

您可以使用以下方式來使用 docker-compose:

* 將 `docs/etc/` copy 到你自訂的 etc directory (ex: `/etc/go-pttbbsweb`).
* 將 `docs/config/01-config.docker.ini` copy 到你自訂的 etc directory 為 production.ini (ex: `cp 01-config.docker.ini /etc/go-pttbbsweb/production.ini`).
* 將 `docker/docker/docker_compose.env.template` copy 到 `docker/docker_compose.env` 並且更改相對應的設定.
* `./scripts/docker_initbbs.sh [BBSHOME] pttofficialapps/go-pttbbs:latest`
* `docker-compose --env-file docker/docker_compose.env -f docker/docker-compose.yaml up -d`
* 在 `http://localhost:3457/account/register` 做 register
* 在 `http://localhost:3457/account/login` 做 login
* `telnet localhost 8888` 並且使用您剛剛登錄的帳號使用.
* 第一次使用時. 須先將 SYSOP 和 guest 建立起來.

## Discussing / Reviewing / Questioning the code

除了開 issues 以外, 您還可以做以下的事情來對於 code 做討論 / review / 提出問題.

* `git clone` 這個 repo.
* 開一個 review-[topic] 的 branch.
* 對於想要討論的部分在 code 裡寫 comments.
* pull-request
* 對於 PR 進行討論.
* 當 PR 關掉時, comments 會留下關於這個 pr 討論的 link.

## Develop

您可以使用 fork 來一起開發.

## Unit-Test

你可以做以下的事情來進行 unit-test:

* `./scripts/test.sh`

您可以做以下的事情來進行 coverage-check:

* `./scripts/coverage.sh`

## Swagger

You can run swagger with:

* 設定 python virtualenv.
* `cd apidoc; pip install . && pip uninstall apidoc -y && python setup.py develop; cd ..`
* `./scripts/swagger.sh [host]`
* go to `http://localhost:5000`

## Schema definition

* `https://github.com/Ptt-official-app/go-pttbbsweb/tree/main/schema`

## Repository Naming

* 2024-06-26: 這個 repo 主要是以 pttbbs 的機制為主, 所以改成命名為 go-pttbbsweb.
* 2022-12-16: 這個 repo 之所以會被稱為 go-openbbsmiddleware, 是因為古早的 [.NET ASP](https://github.com/Ptt-official-app/AspCoreOpenBBSMiddleware) 的開發者希望 [中介應用層](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96) 可以擴及其他版本的 bbs (Maple/中山之島). 這個 repo 就 follow 當時的 naming convention 命名為 go-openbbsmiddleware.
