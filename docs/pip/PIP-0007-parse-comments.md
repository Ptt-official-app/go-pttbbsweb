#Parse Comments

c-pttbbs 推文區目前已知的種類包含著:
1. 推
2. 噓
3. ->
4. reply (作者 edit 時在推噓文下方的回應, 唯一沒有特殊 prefix 的 comments)
5. 轉錄
6. 編輯時間
7. 刪除

由於會有 "盡量完整還原 c-pttbbs 文章" 的需求, 所以需要保留轉錄/編輯/刪除的結果.

c-pttbbs 的推文只有 mm/dd HH:MM. 並沒有年份和秒/sub-seconds.
目前的 strategy 是根據 create-time / mtime 和每個推文所提供的月日時分做夾擊來 inference 年份.

comments 有可能會因為作者的 edit 而有所增減.
並且 comments 有可能會很多 (> 10k) 而無法在短時間內 parse 全部.
目前使用的 [strategy](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/api/get_article_detail.go#L250) 是會先 parse 前面一部分的 comments.
然後把所有的 comments 丟進 queue 裡重新 parse.

comments 另外有[舊版](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/testcase/temp8) comments 需要考慮

Parse comments 分成[兩個階段](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/queue/comment_queue.go#L77)
1. 根據 commentDBCS parse 成 comments.
2. 將 parse 的 comments 和既有的 comments integrate.


## [根據 commentDBCS parse 成 comments](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/comment.go#L34)

1. 估計可能會有多少個 comments.
2. 特別處理推噓文之前的 reply.
3. for-loop:
    * 找到最近的推/噓/->/轉/編輯/刪除
        * 處理相對應的 comments.
        * 嘗試找相對應的 reply.
    * 找下一個推/噓/->/轉/編輯/刪除


## [Integrate comments](https://github.com/Ptt-official-app/go-openbbsmiddleware/blob/main/dbcs/integrate_comments.go)

1. 拿到所有的既有的 comments 的 MD5
2. 根據 edit-distance 找到相同 MD5 的 comments 和相應的 chunkts
3. 對於每個 edit-distance chunk, inference 出相對應的 timestamp:
    * 分成 sort-timestamp 和 create-timestamp. create-timestamp 跟月日時分完全 align. sort-timestamp 會因為 sort-order 而有可能跟 create-timestamp 不一樣.
    * 做 forward-inference timestamp, 直到 timestamp >= article-create-time + 365 天.
    * 根據 end-timestamp 做 backward inference timestamp.
4. 設定 new-commments 和 to-delete-comments.
