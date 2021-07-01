# Big5 <=> Utf8

因為歷史因素. Ptt 的 Big5 => Utf8 是使用 Big5-UAO encoding. \
而不是國際上常見的 Big5(CP950) => Utf8 encoding (golang 的 traditionalchinsese package).

Big5-UAO 還包含了 Big5-ext 裡關於日文字和非常用字/符號的支援. (CP950 有大約 1.3 萬字. Big5 有大約 1.9 萬字) \
Ptt 的 ASCII 需要 Big5-UAO 才能顯示相對應的字.

參考 c-pttbbs ([cmsys/utf8](https://github.com/ptt/pttbbs/blob/master/common/sys/utf8.c#L6), [mbbsd/convert](https://github.com/ptt/pttbbs/blob/master/mbbsd/convert.c#L9)), \
Implement 在 [go-pttbbs/types/big5](https://github.com/Ptt-official-app/go-pttbbs/blob/main/types/big5.go#L154).

