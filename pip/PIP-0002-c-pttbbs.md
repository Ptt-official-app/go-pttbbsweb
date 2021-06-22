# c-pttbbs

c-pttbbs 大致上是每個 user 的 connection 是獨立 fork 出來的 process 與 user 做 interaction. \
使用 shm 和 file 來做 cross-process 的 communication.


## User

使用單一 file (/home/bbs/.PASSWDS) 記錄著所有 users 的重要資訊. 這個 file 有每日備份. \
每個 user 在 .PASSWDS 裡佔 512 bytes. \
並且也會以 hash 的方式存在 shm 裡.

使用者相關的資訊儲存在 /home/bbs/home/[1st-char]/[username] \
其中包括:
* `.fav` 我的最愛
* `.brc3` 使用者看過一年內文章的資訊.
* `USERLOG` 各種 log
* `logins.recent` 最近上站紀錄
* `money.recent` 最近錢的紀錄


## 板

使用單一 file (/home/bbs/.BRD) 記錄著板的 list. 這個 file 有每日備份. \
這個 file 裡的一部分也會被放到 shm 裡. \
並且 shm 裡有包含著"分類-based sorted" 和"title-based sorted" 的結果.

within 板的 file 資訊儲存在 /home/bbs/boards/[1st-char]/[board-name] \
其中包括:

* `.DIR` 所有文章的 list
* `.DIR.bottom` 置底文章的 list.
* `.timecap` 文章被 update 時的各種版本. 會固定時間 recycle.
* `SR.` 各種 search 的暫時結果.


## 文章

文章的 content 以 file 的方式儲存在 /home/bbs/boards/[1st-char]/[board-name]/M.[create-time].A.[XYZ]


## 推文

推文是以 append 的方式直接加在 file 的最後一行. 這也是為什麼推文在 c-pttbbs 無法容易被 index/search.


## 丟水球

根據對方的 process 的 pid:
1. 塞進 shm 裡相對應的 queue.
2. 對 pid 做 kill -s USR1 來丟水球.
