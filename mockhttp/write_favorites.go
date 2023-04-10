package mockhttp

import (
	"os"

	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/sirupsen/logrus"
)

func WriteFavorites(params *api.WriteFavoritesParams) (ret *api.WriteFavoritesResult) {
	filename := "./testcase/home2/t/testUser2/.fav"
	_ = os.WriteFile(filename, params.Content, ptttype.DEFAULT_FILE_CREATE_PERM)

	logrus.Infof("WriteFavorites: content: %v", params.Content)

	FAVORITES_VERSION = 1
	ret = &api.WriteFavoritesResult{
		MTime: 1234567892,
	}

	return ret
}
