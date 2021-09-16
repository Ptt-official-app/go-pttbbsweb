package mockhttp

import (
	"bytes"

	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func EditArticleDetail3(params *api.EditArticleParams) (ret *api.EditArticleResult) {
	content := bytes.Join(params.Content, []byte{'\n'})

	ret = &api.EditArticleResult{
		MTime:     types.Time4(1623511858),
		Content:   content,
		RealTitle: []byte("this is a test"),
		Class:     []byte("\xb4\xfa\xb8\xd5"),
		FullTitle: []byte("[\xb4\xfa\xb8\xd5] this is a test"),
	}

	return ret
}
