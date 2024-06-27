package apitypes

import (
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func ToURL(fbboardID FBoardID, farticleID FArticleID) string {
	if fbboardID == "" || farticleID == "" {
		return ""
	}
	return types.URL_PREFIX + "/" + "board" + "/" + string(fbboardID) + "/" + "article" + "/" + string(farticleID)
}
