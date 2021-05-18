package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func GetArticleDetail(params *api.GetArticleParams) (ret *api.GetArticleResult) {
	ret = &api.GetArticleResult{
		MTime:   types.Time4(1608386280),
		Content: []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI\n\xbc\xd0\xc3D: [\xa4\xdf\xb1o] \xb4\xfa\xb8\xd5\xa4@\xa4U\xafS\xae\xed\xa6r\xa1\xe3\n\xae\xc9\xb6\xa1: Sat Dec 19 21:57:58 2020\n\n\xa1\xb0\xb3o\xbc\xcb\xa4l\xa6\xb3\xba\xf1\xa6\xe2\xb6\xdc\xa1H\xa1\xe3\n\xa1\xb0 \xb5o\xabH\xaf\xb8\n\n--\n\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0 docker(pttdocker.test), \xa8\xd3\xa6\xdb: 172.22.0.1\n"),
	}

	return ret
}
