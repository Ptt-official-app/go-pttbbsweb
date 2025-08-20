package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

var ContentBlock_c *db.Collection

type ContentBlock struct {
	BBoardID  bbs.BBoardID    `bson:"bid"` //
	ArticleID bbs.ArticleID   `bson:"aid"` //
	ContentID types.ContentID `bson:"Cid"`
	Idx       int             `bson:"idx"`
	IsDeleted bool            `bson:"deleted,omitempty"` //

	Content [][]*types.Rune `bson:"content"` //

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_CONTENT_BLOCK = &ContentBlock{}

var (
	CONTENT_BLOCK_BBOARD_ID_b  = getBSONName(EMPTY_CONTENT_BLOCK, "BBoardID")
	CONTENT_BLOCK_ARTICLE_ID_b = getBSONName(EMPTY_CONTENT_BLOCK, "ArticleID")
	CONTENT_BLOCK_CONTENT_ID_b = getBSONName(EMPTY_CONTENT_BLOCK, "ContentID")
	CONTENT_BLOCK_IDX_b        = getBSONName(EMPTY_CONTENT_BLOCK, "Idx")
	CONTENT_BLOCK_IS_DELETED_b = getBSONName(EMPTY_CONTENT_BLOCK, "IsDeleted")

	CONTENT_BLOCK_CONTENT_b        = getBSONName(EMPTY_CONTENT_BLOCK, "Content")
	CONTENT_BLOCK_UPDATE_NANO_TS_b = getBSONName(EMPTY_CONTENT_BLOCK, "UpdateNanoTS")
)

type ContentBlockQuery struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	ArticleID bbs.ArticleID   `bson:"aid"`
	ContentID types.ContentID `bson:"Cid"`
	Idx       int             `bson:"idx"`
}

var EMPTY_CONTENT_BLOCK_QUERY = &ContentBlockQuery{}

func assertContentBlockFields() error {
	if err := assertFields(EMPTY_CONTENT_BLOCK, EMPTY_CONTENT_BLOCK_QUERY); err != nil {
		return err
	}

	return nil
}

func GetContentBlock(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID, idx int) (contentBlock *ContentBlock, err error) {
	return getContentBlock(boardID, articleID, contentID, idx, ContentBlock_c)
}

func GetContentBlocks(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID, startIdx int, n int64) (contentBlocks []*ContentBlock, err error) {
	return getContentBlocks(boardID, articleID, contentID, startIdx, n, ContentBlock_c)
}

func GetAllContentBlocks(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID) (contentBlocks []*ContentBlock, err error) {
	return GetContentBlocks(boardID, articleID, contentID, 0, 0)
}

func UpdateContentBlocks(contentBlocks []*ContentBlock, updateNanoTS types.NanoTS) (err error) {
	return updateContentBlocks(contentBlocks, updateNanoTS, ContentBlock_c)
}
