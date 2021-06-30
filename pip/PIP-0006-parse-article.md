# Parse 文章

c-pttbbs 的文章包含著本文 (content), 站方簽名檔 (signature), 推文區 (comments).

parse 文章被認為是 computationally intensive 的 task. 中介底層只提供 raw-file-data 給中介應用層. 由中介應用層負責做文章的 parsing.


## Split 文章的 Algorithm

在轉錄/Re 時會再次增加站方簽名檔.

推文區目前已知的種類包含著:
1. 推
2. 噓
3. ->
4. reply (作者 edit 時在推噓文下方的回應)
5. 轉錄
6. 編輯時間
7. 刪除

根據以上規則. 有著以下的 [split 文章 algorithm](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/split_article_signature_comments_dbcs.go)

1. 找到所有 match '\n--' 的部分為 match_list.
2. 在 match-list 裡找到最後一個 valid match.
3. 如果有 valid match 的話:
    * valid-match 之前為本文.
    * 決定站方簽名檔的部分.
    * 剩餘的為推文區.
4. 如果沒有任何 match 的話:
    * 根據 pattern 規則找第一個推文.
    * 如果沒有找到第一個推文: 全部都當成是本文.
    * 如果有找到第一個推文: 之前的是本文. 之後的是推文區.

找到的具代表性的文章都放在 [dbcs/testcase/](https://github.com/Ptt-official-app/go-openbbsmiddleware/tree/main/dbcs/testcase) 裡.
