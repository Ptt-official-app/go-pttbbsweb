package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func CreateComment(params *api.CreateCommentParams) (ret *api.CreateCommentResult) {
	ret = &api.CreateCommentResult{
		Content: []byte("\x1b[1;37m\xb1\xc0 \x1b[33mSYSOP\x1b[m\x1b[33m: test123                                                  \x1b[m 12/14 17:15\n"),
		MTime:   1607937324,
	}

	return ret
}
