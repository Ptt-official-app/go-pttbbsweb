package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"go.mongodb.org/mongo-driver/mongo"
)

// ArticleDetailSummary
type ArticleDetailSummary struct {
	BBoardID   bbs.BBoardID  `bson:"bid"`
	ArticleID  bbs.ArticleID `bson:"aid"`
	IsDeleted  bool          `bson:"deleted,omitempty"`
	CreateTime types.NanoTS  `bson:"create_time_nano_ts"`
	MTime      types.NanoTS  `bson:"mtime_nano_ts"`

	Recommend int              `bson:"recommend"`
	Owner     bbs.UUserID      `bson:"owner"`
	Title     string           `bson:"title"`
	Money     int              `bson:"money"`
	Class     string           `bson:"class"`
	Filemode  ptttype.FileMode `bson:"mode"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	ContentMTime        types.NanoTS `bson:"content_mtime_nano_ts"`
	ContentMD5          string       `bson:"content_md5"`
	ContentUpdateNanoTS types.NanoTS `bson:"content_update_nano_ts"`

	FirstCommentsMD5          string       `bson:"first_comments_md5"`
	FirstCommentsLastTime     types.NanoTS `bson:"first_comments_last_time_nano_ts"`
	FirstCommentsUpdateNanoTS types.NanoTS `bson:"first_comments_update_nano_ts"`

	NComments            int          `bson:"n_comments"`
	CommentsUpdateNanoTS types.NanoTS `bson:"comments_update_nano_ts"`

	Rank             int          `bson:"rank"` // 評價
	RankUpdateNanoTS types.NanoTS `bson:"rank_update_nano_ts"`
}

var (
	EMPTY_ARTICLE_DETAIL_SUMMARY = &ArticleDetailSummary{}
	articleDetailSummaryFields   = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_DETAIL_SUMMARY)
)

func GetArticleDetailSummary(bboardID bbs.BBoardID, articleID bbs.ArticleID) (result *ArticleDetailSummary, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	result = &ArticleDetailSummary{}
	err = Article_c.FindOne(query, &result, articleDetailSummaryFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
