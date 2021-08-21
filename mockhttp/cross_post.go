package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func CrossPost(params *api.CrossPostParams) (ret *api.CrossPostResult) {
	articleSummary := &bbs.ArticleSummary{
		BBoardID:    bbs.BBoardID("10_WhoAmI"),
		ArticleID:   bbs.ArticleID("1VrooM21"),
		IsDeleted:   false,
		Filename:    "M.1607937174.A.081",
		CreateTime:  types.Time4(1607937174),
		MTime:       types.Time4(1607937100),
		Recommend:   0,
		Owner:       bbs.UUserID("SYSOP"),
		Class:       []byte{0xc2, 0xe0}, // 轉
		FullTitle:   []byte("Fw: [\xb4\xfa\xb8\xd5]this is a test"),
		Money:       0,
		Filemode:    0,
		Read:        false,
		RealTitle:   []byte("this is a test"),
		SubjectType: ptttype.SUBJECT_FORWARD,
	}

	ret = &api.CrossPostResult{
		ArticleSummary: articleSummary,
		Comment:        []byte("\xa1\xb0 \x1b[1;32mSYSOP\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO WhoAmI\x1b[m                                         12/14 04:12\r"),
		CommentMTime:   1607937174,
	}

	return ret
}
