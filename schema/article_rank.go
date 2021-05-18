package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleRank struct {
	BBoardID  bbs.BBoardID  `bson:"bid"` //
	ArticleID bbs.ArticleID `bson:"aid"` //

	Rank               int          `bson:"rank"` //評價
	RankToUpdateNanoTS types.NanoTS `bson:"rank_to_update_nano_ts"`
	RankUpdateNanoTS   types.NanoTS `bson:"rank_update_nano_ts"`
}

func UpdateArticleRank(boardID bbs.BBoardID, articleID bbs.ArticleID, diffRank int, updateNanoTS types.NanoTS) (newRank int, err error) {

	query := bson.M{
		ARTICLE_BBOARD_ID_b:  boardID,
		ARTICLE_ARTICLE_ID_b: articleID,
	}

	update := bson.M{
		"$inc": bson.M{
			ARTICLE_RANK_b: diffRank,
		},
		"$set": bson.M{
			ARTICLE_RANK_TO_UPDATE_NANO_TS_b: updateNanoTS,
		},
	}

	r, err := Article_c.FindOneAndUpdateNoSet(query, update, true)
	if err != nil {
		return 0, err
	}
	ret := &ArticleRank{}
	err = r.Decode(ret)
	if err != nil {
		return 0, err
	}

	return ret.Rank, nil
}
