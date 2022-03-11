package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var ManContentBlock_c *db.Collection

func GetManContentBlock(boardID bbs.BBoardID, articleID types.ManArticleID, contentID types.ContentID, idx int) (contentBlock *ContentBlock, err error) {
	return getContentBlock(boardID, bbs.ArticleID(articleID), contentID, idx, ManContentBlock_c)
}

func GetManContentBlocks(boardID bbs.BBoardID, articleID types.ManArticleID, contentID types.ContentID, startIdx int, n int64) (contentBlocks []*ContentBlock, err error) {
	return getContentBlocks(boardID, bbs.ArticleID(articleID), contentID, startIdx, n, ManContentBlock_c)
}

func GetAllManContentBlocks(boardID bbs.BBoardID, articleID types.ManArticleID, contentID types.ContentID) (contentBlocks []*ContentBlock, err error) {
	return GetManContentBlocks(boardID, articleID, contentID, 0, 0)
}

func UpdateManContentBlocks(contentBlocks []*ContentBlock, updateNanoTS types.NanoTS) (err error) {
	return updateContentBlocks(contentBlocks, updateNanoTS, ManContentBlock_c)
}
