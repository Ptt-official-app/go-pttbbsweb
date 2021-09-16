package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func GetArticleDetail3(params *api.GetArticleParams) (ret *api.GetArticleResult) {
	ret = &api.GetArticleResult{
		MTime:   types.Time4(1608388624),
		Content: []byte("\xb4\xfa\xb8\xd50\r\n\xb4\xfa\xb8\xd51\r\n\n--\n\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0 docker(pttdocker.test), \xa8\xd3\xa6\xdb: 172.22.0.1\n\xa1\xb0 \xbds\xbf\xe8: SYSOP (127.0.0.1), 03/07/2020 00:24:18\r\n\x1b[1;37m\xb1\xc0 \x1b[33mSYSOP\x1b[m\x1b[33m: test123                                                  \x1b[m 12/14 17:15\n"),
	}

	return ret
}
