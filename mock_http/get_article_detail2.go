package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func GetArticleDetail2(params *api.GetArticleParams) (ret *api.GetArticleResult) {
	ret = &api.GetArticleResult{
		MTime:   types.Time4(1234567890),
		Content: []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI\n\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xa9\xd2\xa5H\xafS\xae\xed\xa6r\xafu\xaa\xba\xacO\xa6\xb3\xba\xf1\xa6\xe2\xaa\xba\xa1\xe3\n\xae\xc9\xb6\xa1: Sat Dec 19 22:35:04 2020\n\n\xb5M\xab\xe1 \\n \xa4\xa3\xb7|\xa6b big5 \xb5\xb2\xa7\xc0. \xa5i\xa5H\xa9\xf1\xa4\xdf\xaa\xbd\xb1\xb5\xa5\xce \\n \xc2_\xa6\xe6.\n\xa7\xda\xacO\xb3\\\xa5\\\xbb\\\n\n--\n\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0 docker(pttdocker.test), \xa8\xd3\xa6\xdb: 172.22.0.1\n\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0                                                       \x1b[m 12/19 22:35\n\x1b[1;37m\xb1\xc0 \x1b[33mchhsiao123\x1b[m\x1b[33m:\xb1\xc0                                                  \x1b[m 12/19 22:36\n\x1b[1;31m\xbcN \x1b[33mchhsiao123\x1b[m\x1b[33m:\xbcN\xa1\xe3                                                \x1b[m 12/19 22:37\n"),
	}

	return ret
}
