package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mand"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type ManArticleSummary struct {
	BBoardID bbs.BBoardID `bson:"bid"` //

	ArticleID types.ManArticleID `bson:"aid"`       //
	LevelIdx  types.ManArticleID `bson:"level_idx"` //

	CreateTime types.NanoTS `bson:"create_time_nano_ts"` //
	MTime      types.NanoTS `bson:"mtime_nano_ts"`

	Title string `bson:"title"` //
	IsDir bool   `bson:"is_dir"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"` // used by article-summary

	Idx int `bson:"pttidx"`
}

var (
	EMPTY_MAN_ARTICLE_SUMMARY = &ManArticleSummary{}
	manArticleSummaryFields   = getFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_SUMMARY)
)

func NewManArticleSummaryFromPB(entry *mand.Entry, boardID bbs.BBoardID, levelIdx types.ManArticleID, updateNanoTS types.NanoTS, idx int) (articleSummary *ManArticleSummary) {
	articleID := types.ManArticleID(entry.Path)
	createTime := articleID.ToCreateTime()

	return &ManArticleSummary{
		BBoardID:  boardID,
		LevelIdx:  levelIdx,
		ArticleID: articleID,

		CreateTime: createTime,
		MTime:      createTime,

		Title: entry.Title,
		IsDir: entry.IsDir,

		UpdateNanoTS: updateNanoTS,

		Idx: idx,
	}
}

func UpdateManArticleSummaries(articleSummaries []*ManArticleSummary, updateNanoTS types.NanoTS) (err error) {
	if len(articleSummaries) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(articleSummaries))
	for idx, each := range articleSummaries {
		query := &ManArticleQuery{
			BBoardID: each.BBoardID,
			LevelIdx: each.LevelIdx,
			Idx:      each.Idx,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := ManArticle_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(articleSummaries)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateArticleSummaries := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*ManArticleQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					MAN_ARTICLE_BBOARD_ID_b: origFilter.BBoardID,
					MAN_ARTICLE_LEVEL_IDX_b: origFilter.LevelIdx,
					MAN_ARTICLE_IDX_b:       origFilter.Idx,
					MAN_ARTICLE_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},

					MAN_ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
				},
				bson.M{
					MAN_ARTICLE_BBOARD_ID_b: origFilter.BBoardID,
					MAN_ARTICLE_IDX_b:       origFilter.Idx,
					MAN_ARTICLE_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},

					MAN_ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
				},
			},
		}
		each.Filter = filter
		updateArticleSummaries = append(updateArticleSummaries, each)
	}

	// update items with comparing update-nano-ts
	_, err = ManArticle_c.BulkUpdateOneOnly(updateArticleSummaries)

	return err
}

func RemoveManArticleSummaries(boardID bbs.BBoardID, levelIdx types.ManArticleID, idx int) (err error) {
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b: boardID,
		MAN_ARTICLE_LEVEL_IDX_b: levelIdx,
		MAN_ARTICLE_IDX_b: bson.M{
			"$gte": idx,
		},
	}

	return ManArticle_c.Remove(query)
}

func GetManArticleSummaries(boardID bbs.BBoardID, LevelIdx types.ManArticleID) (result []*ManArticleSummary, err error) {
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b: boardID,
		MAN_ARTICLE_LEVEL_IDX_b: LevelIdx,
	}

	// sort opts
	sortOpts := bson.D{
		{Key: MAN_ARTICLE_IDX_b, Value: 1},
	}

	// find
	err = ManArticle_c.Find(query, int64(0), &result, manArticleSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}
