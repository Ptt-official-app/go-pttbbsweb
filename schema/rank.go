package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

var Rank_c *db.Collection

type Rank struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`
	Owner     bbs.UUserID   `bson:"owner"`

	Rank int `bson:"rank"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_RANK = &Rank{}

var (
	RANK_BBOARD_ID_b      = getBSONName(EMPTY_RANK, "BBoardID")
	RANK_ARTICLE_ID_b     = getBSONName(EMPTY_RANK, "ArticleID")
	RANK_OWNER_b          = getBSONName(EMPTY_RANK, "Owner")
	RANK_RANK_b           = getBSONName(EMPTY_RANK, "Rank")
	RANK_UPDATE_NANO_TS_b = getBSONName(EMPTY_RANK, "UpdateNanoTS")
)

type RankAgged struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`

	Rank int `bson:"rank"`
}

func SumRankByBoardID(boardID bbs.BBoardID, articleIDs []bbs.ArticleID) (ret []*RankAgged, err error) {
	query := bson.M{
		RANK_BBOARD_ID_b: boardID,
	}
	if articleIDs != nil {
		query[RANK_ARTICLE_ID_b] = bson.M{
			"$in": articleIDs,
		}
	}

	group := bson.M{
		"_id": bson.M{
			RANK_ARTICLE_ID_b: "$" + RANK_ARTICLE_ID_b,
		},
		"sum": bson.M{
			"$sum": "$" + RANK_RANK_b,
		},
	}

	rets, err := Rank_c.Aggregate(query, group)
	if err != nil {
		return nil, err
	}

	if len(rets) == 0 {
		return nil, nil
	}

	ret = make([]*RankAgged, len(rets))
	for idx, each := range rets {
		theID := each["_id"].(bson.M)
		eachRet := &RankAgged{
			BBoardID:  boardID,
			ArticleID: bbs.ArticleID(theID[RANK_ARTICLE_ID_b].(string)),
			Rank:      int(each["sum"].(int32)),
		}
		ret[idx] = eachRet
	}

	return ret, nil
}

func SumRank(boardID bbs.BBoardID, articleID bbs.ArticleID) (total int, err error) {
	query := bson.M{
		RANK_BBOARD_ID_b:  boardID,
		RANK_ARTICLE_ID_b: articleID,
	}

	group := bson.M{
		"_id": nil,
		"sum": bson.M{
			"$sum": "$" + RANK_RANK_b,
		},
	}

	rets, err := Rank_c.Aggregate(query, group)
	if err != nil {
		return 0, err
	}

	if len(rets) == 0 {
		return 0, nil
	}

	total = int(rets[0]["sum"].(int32))

	return total, nil
}

func UpdateRank(boardID bbs.BBoardID, articleID bbs.ArticleID, ownerID bbs.UUserID, rank int, updateNanoTS types.NanoTS) (origRank int, err error) {
	query := bson.M{
		RANK_BBOARD_ID_b:  boardID,
		RANK_ARTICLE_ID_b: articleID,
		RANK_OWNER_b:      ownerID,
	}

	theRank := &Rank{
		BBoardID:     boardID,
		ArticleID:    articleID,
		Owner:        ownerID,
		Rank:         rank,
		UpdateNanoTS: updateNanoTS,
	}

	r, err := Rank_c.CreateOnly(query, theRank)
	if err != nil {
		return 0, err
	}
	if r.UpsertedCount > 0 {
		return 0, nil
	}

	query = bson.M{
		"$or": bson.A{
			bson.M{
				RANK_BBOARD_ID_b:  boardID,
				RANK_ARTICLE_ID_b: articleID,
				RANK_OWNER_b:      ownerID,
				RANK_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},
			},
			bson.M{
				RANK_BBOARD_ID_b:  boardID,
				RANK_ARTICLE_ID_b: articleID,
				RANK_OWNER_b:      ownerID,
				RANK_UPDATE_NANO_TS_b: bson.M{
					"$lt": updateNanoTS,
				},
			},
		},
	}

	r2, err := Rank_c.FindOneAndUpdate(query, theRank, false)
	if err != nil {
		return 0, err
	}
	retRank := &Rank{}
	err = r2.Decode(retRank)
	if err != nil {
		return 0, err
	}

	return retRank.Rank, nil
}
