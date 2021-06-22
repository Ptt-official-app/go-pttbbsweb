# 程式架構說明

* `api/` api 主要的 functions.
* `apitypes/` api 所需要使用到的 types.
* `schema/` 存在 db 裡的各種 data 的 schema.
* `types/` 基本的 types.
* `utils/` 基本的 utils.
* `queue/` 把要處理的 comments 丟到 queue 裡相關的 module.
* `dbcs/` 處理 dbcs (double-byte-color-system) 相關的 module.
* `db/` db 的基本 wrapping functions.
* `mock/` go-openbbsmiddleware 給 frontend 的假資料.
* `mock_http/` 測試時的 go-pttbbs 的假資料.
