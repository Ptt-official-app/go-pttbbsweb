package mockhttp

import (
	"os"

	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/sirupsen/logrus"
)

func GetFavorites(params *api.GetFavoritesParams) (ret *api.GetFavoritesResult) {
	filename := "./testcase/home1/t/testUser2/.fav"
	if FAVORITES_VERSION == 1 {
		filename = "./testcase/home2/t/testUser2/.fav"
	}

	content, _ := os.ReadFile(filename)

	logrus.Infof("GetFavorites: content: %v", content)

	mtime := types.Time4(1234567890)
	if FAVORITES_VERSION == 1 {
		mtime = types.Time4(1234567892)
	}
	ret = &api.GetFavoritesResult{
		MTime:   mtime,
		Content: content,
	}

	return ret
}
