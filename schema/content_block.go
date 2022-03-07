package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
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
	query := bson.M{
		CONTENT_BLOCK_BBOARD_ID_b:  boardID,
		CONTENT_BLOCK_ARTICLE_ID_b: articleID,
		CONTENT_BLOCK_CONTENT_ID_b: contentID,
		CONTENT_BLOCK_IDX_b:        idx,
	}
	contentBlock = &ContentBlock{}
	err = ContentBlock_c.FindOne(query, &contentBlock, nil)
	if err != nil {
		return nil, err
	}

	return contentBlock, nil
}

func GetContentBlocks(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID, startIdx int, n int64) (contentBlocks []*ContentBlock, err error) {
	query := bson.M{
		CONTENT_BLOCK_BBOARD_ID_b:  boardID,
		CONTENT_BLOCK_ARTICLE_ID_b: articleID,
		CONTENT_BLOCK_CONTENT_ID_b: contentID,
		CONTENT_BLOCK_IDX_b: bson.M{
			"$gte": startIdx,
		},
	}

	sort := bson.D{
		{Key: CONTENT_BLOCK_IDX_b, Value: 1},
	}

	err = ContentBlock_c.Find(query, n, &contentBlocks, nil, sort)
	if err != nil {
		return nil, err
	}

	return contentBlocks, nil
}

func GetAllContentBlocks(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID) (contentBlocks []*ContentBlock, err error) {
	return GetContentBlocks(boardID, articleID, contentID, 0, 0)
}

func UpdateContentBlocks(contentBlocks []*ContentBlock, updateNanoTS types.NanoTS) (err error) {
	if len(contentBlocks) == 0 {
		return nil
	}

	p_contentBlocks := contentBlocks

	var first []*ContentBlock

	for block := getBlock(len(p_contentBlocks), MAX_CONTENT_BLOCK); len(p_contentBlocks) > 0; block = getBlock(len(p_contentBlocks), MAX_CONTENT_BLOCK) {
		first, p_contentBlocks = p_contentBlocks[:block], p_contentBlocks[block:]

		err = updateContentBlocksCore(first, updateNanoTS)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateContentBlocksCore(contentBlocks []*ContentBlock, updateNanoTS types.NanoTS) (err error) {
	theList := make([]*db.UpdatePair, len(contentBlocks))
	for idx, each := range contentBlocks {
		filter := &ContentBlockQuery{
			BBoardID:  each.BBoardID,
			ArticleID: each.ArticleID,
			ContentID: each.ContentID,
			Idx:       each.Idx,
		}

		theList[idx] = &db.UpdatePair{
			Filter: filter,
			Update: each,
		}
	}
	r, err := ContentBlock_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(contentBlocks)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateContentBlocks := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*ContentBlockQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					CONTENT_BLOCK_BBOARD_ID_b:  origFilter.BBoardID,
					CONTENT_BLOCK_ARTICLE_ID_b: origFilter.ArticleID,
					CONTENT_BLOCK_CONTENT_ID_b: origFilter.ContentID,
					CONTENT_BLOCK_IDX_b:        origFilter.Idx,
					COMMENT_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},
				},
				bson.M{
					CONTENT_BLOCK_BBOARD_ID_b:  origFilter.BBoardID,
					CONTENT_BLOCK_ARTICLE_ID_b: origFilter.ArticleID,
					CONTENT_BLOCK_CONTENT_ID_b: origFilter.ContentID,
					CONTENT_BLOCK_IDX_b:        origFilter.Idx,
					CONTENT_BLOCK_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},
				},
			},
		}

		each.Filter = filter
		origUpdate := each.Update
		each.Update = bson.M{
			"$set": origUpdate,
			"$unset": bson.M{
				COMMENT_IS_DELETED_b: true,
			},
		}
		updateContentBlocks = append(updateContentBlocks, each)
	}
	_, err = ContentBlock_c.BulkUpdateOneOnlyNoSet(updateContentBlocks)

	return err
}
