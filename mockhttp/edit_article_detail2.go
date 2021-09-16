package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func EditArticleDetail2(params *api.EditArticleParams) (ret *api.EditArticleResult) {
	ret = &api.EditArticleResult{
		MTime:     types.Time4(1583511858),
		Content:   []byte("\xb4\xfa\xb8\xd50\r\n\xb4\xfa\xb8\xd51\r\n\n--\n\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0 docker(pttdocker.test), \xa8\xd3\xa6\xdb: 172.22.0.1\n\xa1\xb0 \xbds\xbf\xe8: SYSOP (127.0.0.1), 03/07/2020 00:24:18\r\n"),
		RealTitle: []byte("this is a test"),
		Class:     []byte("\xb4\xfa\xb8\xd5"),
		FullTitle: []byte("[\xb4\xfa\xb8\xd5] this is a test"),
	}

	return ret
}
