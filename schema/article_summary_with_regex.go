package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ArticleSummaryWithRegex
type ArticleSummaryWithRegex struct {
	// ArticleSummary
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

	TitleRegex []string `bson:"title_regex"`

	Idx string `bson:"pttidx"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	NComments int `bson:"n_comments,omitempty"` // n_comments is read-only in article-summary.

	Rank int `bson:"rank,omitempty"` // 評價, read-only

	SubjectType ptttype.SubjectType `bson:"subject_type"`
}

var (
	EMPTY_ARTICLE_SUMMARY_WITH_REGEX = &ArticleSummaryWithRegex{}
	articleSummaryWithRegexFields    = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY_WITH_REGEX)
)

//NewArticleSummaryWithRegex
//
//no n_comments in bbs.ArticleSummary from backend.
func NewArticleSummaryWithRegex(a_b *bbs.ArticleSummary, updateNanoTS types.NanoTS) *ArticleSummaryWithRegex {
	title := types.Big5ToUtf8(a_b.RealTitle)

	titleRegex := articleTitleToTitleRegex(title)
	return &ArticleSummaryWithRegex{
		BBoardID:   a_b.BBoardID,
		ArticleID:  a_b.ArticleID,
		IsDeleted:  a_b.IsDeleted,
		CreateTime: types.Time4ToNanoTS(a_b.CreateTime),
		MTime:      types.Time4ToNanoTS(a_b.MTime),
		Recommend:  int(a_b.Recommend),
		Owner:      a_b.Owner,
		FullTitle:  types.Big5ToUtf8(a_b.FullTitle),
		Title:      title,
		Money:      int(a_b.Money),
		Class:      types.Big5ToUtf8(a_b.Class),
		Filemode:   a_b.Filemode,

		TitleRegex: titleRegex,

		UpdateNanoTS: updateNanoTS,

		Idx: a_b.Idx,

		SubjectType: a_b.SubjectType,
	}
}

// articleTitleToTitleRegex
//
// params:
//    title:
//    theClass: (with 4-byte length)
func articleTitleToTitleRegex(title string) (titleRegex []string) {
	titleRune := []rune(title)
	nGramTitleRegex := TITLE_REGEX_N_GRAM
	if len(titleRune) < TITLE_REGEX_N_GRAM {
		nGramTitleRegex = len(titleRune)
	}

	lenTitleRegex := nGramTitleRegex * len(titleRune)

	titleRegex = make([]string, 0, lenTitleRegex)
	titleRegex = articleTitleToTitleRegexCore(titleRune, titleRegex, nGramTitleRegex)

	return titleRegex
}

func articleTitleToTitleRegexCore(titleRune []rune, titleRegex []string, nGram int) (newTitleRegex []string) {
	for i := 1; i <= nGram; i++ {
		endPos := len(titleRune) - i
		for idx := 0; idx <= endPos; idx++ {
			titleRegex = append(titleRegex, string(titleRune[idx:(idx+i)]))
		}
	}

	return titleRegex
}

func UpdateArticleSummaryWithRegexes(articleSummaryWithRegexes []*ArticleSummaryWithRegex, updateNanoTS types.NanoTS) (err error) {
	if len(articleSummaryWithRegexes) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(articleSummaryWithRegexes))
	for idx, each := range articleSummaryWithRegexes {
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
	if r.UpsertedCount == int64(len(articleSummaryWithRegexes)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateArticleSummaryWithRegexes := make([]*db.UpdatePair, 0, len(theList))
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
		updateArticleSummaryWithRegexes = append(updateArticleSummaryWithRegexes, each)
	}

	// update items with comparing update-nano-ts
	_, err = Article_c.BulkUpdateOneOnly(updateArticleSummaryWithRegexes)

	return err
}

func GetArticleSummaryWithRegex(bboardID bbs.BBoardID, articleID bbs.ArticleID) (result *ArticleSummaryWithRegex, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	result = &ArticleSummaryWithRegex{}
	err = Article_c.FindOne(query, &result, articleSummaryWithRegexFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
