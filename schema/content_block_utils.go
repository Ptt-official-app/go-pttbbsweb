package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

func getContentBlock(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID, idx int, coll *db.Collection) (contentBlock *ContentBlock, err error) {
	query := bson.M{
		CONTENT_BLOCK_BBOARD_ID_b:  boardID,
		CONTENT_BLOCK_ARTICLE_ID_b: articleID,
		CONTENT_BLOCK_CONTENT_ID_b: contentID,
		CONTENT_BLOCK_IDX_b:        idx,
	}
	contentBlock = &ContentBlock{}
	err = coll.FindOne(query, &contentBlock, nil)
	if err != nil {
		return nil, err
	}

	return contentBlock, nil
}

func getContentBlocks(boardID bbs.BBoardID, articleID bbs.ArticleID, contentID types.ContentID, startIdx int, n int64, coll *db.Collection) (contentBlocks []*ContentBlock, err error) {
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

	err = coll.Find(query, n, &contentBlocks, nil, sort)
	if err != nil {
		return nil, err
	}

	return contentBlocks, nil
}

func updateContentBlocks(contentBlocks []*ContentBlock, updateNanoTS types.NanoTS, coll *db.Collection) (err error) {
	if len(contentBlocks) == 0 {
		return nil
	}

	p_contentBlocks := contentBlocks

	var first []*ContentBlock

	for block := getBlock(len(p_contentBlocks), MAX_CONTENT_BLOCK); len(p_contentBlocks) > 0; block = getBlock(len(p_contentBlocks), MAX_CONTENT_BLOCK) {
		first, p_contentBlocks = p_contentBlocks[:block], p_contentBlocks[block:]

		err = updateContentBlocksCore(first, updateNanoTS, coll)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateContentBlocksCore(contentBlocks []*ContentBlock, updateNanoTS types.NanoTS, coll *db.Collection) (err error) {
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
	r, err := coll.BulkCreateOnly(theList)
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
	_, err = coll.BulkUpdateOneOnlyNoSet(updateContentBlocks)

	return err
}
