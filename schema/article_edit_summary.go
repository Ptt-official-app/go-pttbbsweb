package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleEditSummary struct {
	MTime      types.NanoTS `bson:"mtime_nano_ts"`
	FullTitle  string       `bson:"full_title"`
	Title      string       `bson:"title"`
	TitleRegex []string     `bson:"title_regex"`
	Class      string       `bson:"class"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

func NewArticleEditSummary(mTime types.NanoTS, titleRaw []byte, fullTitleRaw []byte, theClassRaw []byte, updateNanoTS types.NanoTS) *ArticleEditSummary {
	title := types.Big5ToUtf8(titleRaw)
	titleRegex := articleTitleToTitleRegex(title)

	return &ArticleEditSummary{
		MTime:      mTime,
		FullTitle:  types.Big5ToUtf8(fullTitleRaw),
		Title:      title,
		TitleRegex: titleRegex,
		Class:      types.Big5ToUtf8(theClassRaw),

		UpdateNanoTS: updateNanoTS,
	}
}

func UpdateArticleEditSummary(boardID bbs.BBoardID, articleID bbs.ArticleID, articleEditSummary *ArticleEditSummary, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:  boardID,
		ARTICLE_ARTICLE_ID_b: articleID,
		ARTICLE_UPDATE_NANO_TS_b: bson.M{
			"$lt": updateNanoTS,
		},
	}

	r, err := Article_c.UpdateOneOnly(query, articleEditSummary)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
