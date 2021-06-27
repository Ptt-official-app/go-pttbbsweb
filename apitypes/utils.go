package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func ToURL(fbboardID FBoardID, farticleID FArticleID) string {
	return types.URL_PREFIX + "/" + string(fbboardID) + "/" + string(farticleID) + ".html"
}
