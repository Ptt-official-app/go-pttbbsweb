# Introduction

Ptt ([Web 版](https://term.ptt.cc/), [Publicly Accessible 版](https://www.ptt.cc/bbs/index.html)) \
為台灣最知名的 bbs 站. 2019 年底時八卦板的一篇文章幫助了台灣及早因應 COVID-19 的危機而讓台灣平安度過了 2020.

但是 Ptt 的架構為大約 20 年前時使用 C-lang, 並大量使用 shared-mem / specialized file-structure 來增進單一機器的效率. \
現在有著許多 horizontally scalable 的方式. 希望這次的改動可以讓 Ptt 可以 horizontally scalable, 進而讓 Ptt 可以 self-sustainable.

# 系統架構

在一次的討論會議中, 決定新的 ptt 的架構會是[這張圖](https://github.com/Ptt-official-app/ptt_official_app_wanted#%E7%B3%BB%E7%B5%B1%E6%9E%B6%E6%A7%8B):

* 會有一段過渡期. 過渡期時 c-pttbbs 還是會需要存在原本的機器運作. \
  並不希望新的 ptt program 過於影響到既有的 c-pttbbs 在原本的機器運作.
* 新的架構會逐漸地將 c-pttbbs 的功能移到中介應用層 + DB.
* 新的架構會有一個 go-program (中介底層) 直接做跟 c-pttbbs 一樣的事情 + 根據中介應用層的需求提供 api. \
  目的是類似中介應用層的 file-based db.
* computationally intensive 的部分 (DBCS <=> utf8, parse 文章) 會在中介應用層做. 減少 c-pttbbs 的機器的負擔.
* 這個新架構可以成為以後 horizontally scale-out 的基本架構.

# Implementation

根據系統架構:

* [go-openbbsmiddleware](https://github.com/Ptt-official-app/go-openbbsmiddleware) implements [這張圖](https://github.com/Ptt-official-app/ptt_official_app_wanted#%E7%B3%BB%E7%B5%B1%E6%9E%B6%E6%A7%8B) 的中介應用層.
* [go-pttbbs](https://github.com/Ptt-official-app/go-pttbbs) implements [這張圖](https://github.com/Ptt-official-app/ptt_official_app_wanted#%E7%B3%BB%E7%B5%B1%E6%9E%B6%E6%A7%8B) 的中介底層.
* 使用 [MongoDB](https://www.mongodb.com/) 為相對應的 DB.
* [demo-pttbbs](https://github.com/Ptt-official-app/demo-pttbbs) 展示 [go-openbbsmiddleware](https://github.com/Ptt-official-app/go-openbbsmiddleware) 的使用範例.
