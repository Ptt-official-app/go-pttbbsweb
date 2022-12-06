package mockhttp

import (
	"os"

	"github.com/Ptt-official-app/go-pttbbs/api"
)

func GetFavorites(params *api.GetFavoritesParams) (ret *api.GetFavoritesResult) {
	filename := "./testcase/home1/t/testUser2/.fav"
	content, _ := os.ReadFile(filename)
	ret = &api.GetFavoritesResult{
		MTime:   1234567890,
		Content: content,
	}

	return ret
}
