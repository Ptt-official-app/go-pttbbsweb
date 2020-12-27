package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
)

func ToURL(bboardID bbs.BBoardID, articleID bbs.ArticleID) string {
	filenameRaw := articleID.ToFilename()
	filename := pttbbstypes.CstrToString(filenameRaw[:])
	return types.URL_PREFIX + "/" + string(bboardID) + "/" + filename + ".html"
}
