# 使用 MongoDB

MongoDB 有著以下的 features:

* 以 json 來表示 data.
* SQL 人眼裡的 NoSQL (不需要先訂 schema, document 裡可以隨時加新的 columns.)
* 使用 B-tree 做 indexing. 所以可以有 $lt / $gte 這樣子的 operators (hash-based DB 無法使用 $lt/$gte). (https://docs.mongodb.com/manual/indexes/)
* find 的 return 是 iterator. 所以可以再拿 iterator 再做 within-db / cross-collection 的 join. 然後最後再 eval. \
  如果 find 的過程都是有合適的 indexing 的話. mongo 可以利用如此方式達到 O(nlgn) 的 join.
* 容易 horizontally-scalable. 包含著 replica / sharding 的機制. replica 使用 raft 選 leader.
* indexing key 可以拿來做 auto-sharding. auto-sharding 的機制還蠻清楚明暸. \
  就是利用把一整個 key-space 分成很多個 (lower-key/upper-key) 的 chunk. \
  然後 given 一個 indexing-key, 可以利用 $gte lower-key / $lt upper-key 知道是在哪個 chunk.
* 可能用不到, 不過 4.0 以後增加了 transaction 的機制.
