# Parse DBCS

在 c-pttbbs 裡. 使用的 encoding 是 Big5 + [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code). \
另外利用 Big5 除了 0~127 以外, 都是 double-byte 的特性, 可以達到一字雙色的效果.

c-pttbbs 裡有預期通常一行是 80 chars. 以 '\n' 或是 '\r\n' 做為分行依據. 另外有特殊起始字會讓整行呈現綠色.

c-pttbbs 另外有 BBSMovie. 在這個 PIP 裡先暫時不處理 BBSMovie 的情形.

這個年代的通用 encoding 是 utf8. 需要設計相對應的 data structure 來將 c-pttbbs 裡的文章轉為 utf8-based content.

MongoDB 有著每個 document 32MB 的上限. c-pttbbs 裡有可能會出現超長文章. 尤其是 BBSMovie 似乎以前曾經出現過 1G+ 的文章. \
不過似乎是後來會被站長群關切. 目前先假設所有本文都 < 1M. 可以被裝進 MongoDB 的單一 document. \
先不需要考慮文章過長而必須要放在 multiple MongoDB documents 裡.

由於 Parse DBCS 需要對於每個 bytes 作轉換. 目前被認為是 computationally intensive 的 task. \
所以在讀取文章時, 中介底層只提供 raw-file-data 給中介應用層. 而由中介應用層負責做 DBCS => utf8 的 Parse.


## Data Structure

在 c-pttbbs 裡, 每一篇文章會分成本文 (content) / 站方簽名 (signature) / 推文群 (comments) \
本文 (content) 包含著以 '\n' 或 '\r\n' 分開的數個行. \
每行裡由一個或多個不同顏色屬性的 bytes 所組成.

目前定義 [rune](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/types/rune.go) 如下:

```
type Rune struct {
    Utf8   string `json:"text" bson:"utf8"` //utf8-string
    Big5   []byte `json:"-" bson:"big5"`    //big5-bytes, stored in db for debugging.
    DBCS   []byte `json:"-" bson:"dbcs"`    //dbcs-bytes, stored in db for concat and debugging.
    Color0 Color  `json:"color0" bson:"color0"`
    Color1 Color  `json:"color1" bson:"color1"`
}
```

其中:

* DBCS 是 c-pttbbs 裡包含顏色屬性的 bytes.
* Big5 是去掉顏色屬性的純文字 bytes.
* Utf8 是 Big5 => utf8 的 string.
* Color0 是一字雙色的第 1 個 byte 的顏色.
* Color1 是一字雙色的第 2 個 byte 的顏色.

所以每行就會是 \[\]\*Rune, 每個本文就會是 \[\]\[\]\*Rune


## [DBCS => Big5](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/dbcs.go#L148)

對於每行的 DBCS => big5 大致如下:

1. state 的定義包含:
    * none: 沒有進入到 double-byte 裡.
    * lead: double-byte 的起始.
    * tail: double-byte 的結尾.

1. 檢查是否尾端是 '\r', 如果是的話. 則捨棄 '\r'
2. estimate 最多需要的 big5 blocks.
3. init
    * state 為 none
    * start-idx 為 0
    * color1 為 invalid-color
    * dbcs0pos 為 -1 (一字雙色起始位置)
4. for-loop 每個 byte:
    * 如果 ch 不是 '\x1b', 如果目前(到前一個 byte) 的 state 為:
        * lead: state 變成 tail.
        * none: 如果 ch >= '0x80': 設定 state 為 lead, 設定 dbcs0pos, reset color1.
        * tail: 把目前的 bytes-block 加到 rune list 裡. 重新設定 start-idx 和 color0. \
                並且如果 ch >= '0x80', 則設定 state 為 lead. 否則設定 state 為 none.
    * 如果 ch 是 '\x1b' (要準備換顏色):
        1. 如果目前(到前一個 byte) 的 state 為:
            * lead (假設要換顏色為 1 字雙色): 把目前的 bytes-block 加到 rune list 裡. 重新設定 start-idx 為 lead.
            * none (假設要換顏色): 把目前的 bytes-block 加到 rune-list 裡. 重新設定 start-idx 為目前的 idx.
            * tail (假設要換顏色): 把目前的 bytes-block 加到 rune-list 裡. 重新設定 start-idx 為目前的 idx. state 為 none.
        2. parse color. 設定 color 和新的 idx.
5. 處理 for-loop 結束後的 state.


## [Big5 => utf8](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/dbcs.go#L102)

對於每個 Rune. 直接將 Big5 轉為 utf8.


## [Utf8 => DBCS](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/dbcs.go#L40)

1. 如果已經存在 DBCS. 則直接 return DBCS.
2. parse color0bytes, color1bytes, big5bytes. 將這 3 個組合起來成為 DBCS.
