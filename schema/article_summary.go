package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	ptttypes "github.com/Ptt-official-app/go-pttbbs/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//ArticleSummary
type ArticleSummary struct {
	//ArticleSummary
	BBoardID   bbs.BBoardID  `bson:"bid"`
	ArticleID  bbs.ArticleID `bson:"aid"`
	IsDeleted  bool          `bson:"deleted,omitempty"`
	CreateTime types.NanoTS  `bson:"create_time_nano_ts"`
	MTime      types.NanoTS  `bson:"mtime_nano_ts"`

	Recommend int              `bson:"recommend"`
	Owner     bbs.UUserID      `bson:"owner"`
	FullTitle string           `bson:"full_title"`
	Title     string           `bson:"title"`
	Money     int              `bson:"money"`
	Class     string           `bson:"class"`
	Filemode  ptttype.FileMode `bson:"mode"`

	Idx string `bson:"pttidx"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	NComments int `bson:"n_comments,omitempty"` //n_comments is read-only in article-summary.

	Rank int `bson:"rank,omitempty"` //評價, read-only
}

var (
	EMPTY_ARTICLE_SUMMARY = &ArticleSummary{}
	articleSummaryFields  = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY)
)

func GetArticleSummary(bboardID bbs.BBoardID, articleID bbs.ArticleID) (result *ArticleSummary, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	result = &ArticleSummary{}
	err = Article_c.FindOne(query, &result, articleSummaryFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

//NewARticleSummary
//
//no n_comments in bbs.ArticleSummary from backend.
func NewArticleSummary(a_b *bbs.ArticleSummary, updateNanoTS types.NanoTS) *ArticleSummary {

	//XXX we use Title[6:] for now.
	//TODO: rely on go-pttbbs Title.ToRealTitle()
	fullTitle := ptttypes.CstrToBytes(a_b.Title)
	title := fullTitle
	if len(a_b.Class) > 0 {
		title = title[6:]
	}
	return &ArticleSummary{
		BBoardID:   a_b.BBoardID,
		ArticleID:  a_b.ArticleID,
		IsDeleted:  a_b.IsDeleted,
		CreateTime: types.Time4ToNanoTS(a_b.CreateTime),
		MTime:      types.Time4ToNanoTS(a_b.MTime),
		Recommend:  int(a_b.Recommend),
		Owner:      a_b.Owner,
		FullTitle:  types.Big5ToUtf8(fullTitle),
		Title:      types.Big5ToUtf8(title),
		Money:      int(a_b.Money),
		Class:      types.Big5ToUtf8(a_b.Class),
		Filemode:   a_b.Filemode,

		UpdateNanoTS: updateNanoTS,

		Idx: a_b.Idx,
	}
}

func UpdateArticleSummaries(articleSummaries []*ArticleSummary, updateNanoTS types.NanoTS) (err error) {
	if len(articleSummaries) == 0 {
		return nil
	}

	//create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(articleSummaries))
	for idx, each := range articleSummaries {
		query := &ArticleQuery{
			BBoardID:  each.BBoardID,
			ArticleID: each.ArticleID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := Article_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(articleSummaries)) { //all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateArticleSummaries := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*ArticleQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					ARTICLE_BBOARD_ID_b:  origFilter.BBoardID,
					ARTICLE_ARTICLE_ID_b: origFilter.ArticleID,
					ARTICLE_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},

					ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
				},
				bson.M{
					ARTICLE_BBOARD_ID_b:  origFilter.BBoardID,
					ARTICLE_ARTICLE_ID_b: origFilter.ArticleID,
					ARTICLE_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},

					ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
				},
			},
		}
		each.Filter = filter
		updateArticleSummaries = append(updateArticleSummaries, each)
	}

	//update items with comparing update-nano-ts
	_, err = Article_c.BulkUpdateOneOnly(updateArticleSummaries)

	return err
}
