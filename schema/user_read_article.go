package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type UserReadArticle struct {
	UserID       string
	ArticleID    string
	UpdateNanoTS types.NanoTS
}
