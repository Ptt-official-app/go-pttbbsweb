package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type ArticleSummary struct {
	FBoardID   FBoardID         `json:"bid"`         //
	ArticleID  FArticleID       `json:"aid"`         //
	IsDeleted  bool             `json:"deleted"`     //
	CreateTime types.Time8      `json:"create_time"` //
	MTime      types.Time8      `json:"modified"`    //
	Recommend  int              `json:"recommend"`   //
	NComments  int              `json:"n_comments"`  //
	Owner      bbs.UUserID      `json:"owner"`       //
	Title      string           `json:"title"`       //
	Money      int              `json:"money"`       //
	Class      string           `json:"class"`       // can be: R: 轉, [class]
	Filemode   ptttype.FileMode `json:"mode"`        //

	URL  string `json:"url"`  //
	Read bool   `json:"read"` //

	Idx string `json:"idx"`

	Rank int `json:"rank"`

	SubjectType ptttype.SubjectType `json:"subject_type"`
}

func ToFTitle(title string) string {
	switch types.SERVICE_MODE {
	case types.DEV:
		return "[測試]" + title
	default:
		return title
	}
}

func NewArticleSummary(a_db *schema.ArticleSummary) *ArticleSummary {
	fboardID := ToFBoardID(a_db.BBoardID)
	// if article is deleted, hide articleId
	// https://github.com/Ptt-official-app/go-openbbsmiddleware/issues/253#issuecomment-971526173
	var fTitle string
	var farticleID FArticleID
	if a_db.IsDeleted {
		farticleID = ""
		fTitle = "本文已被刪除"
	} else {
		fTitle =  ToFTitle(a_db.Title)
		farticleID = ToFArticleID(a_db.ArticleID)
	}


	url := ToURL(fboardID, farticleID)

	return &ArticleSummary{
		FBoardID:   fboardID,                  //
		ArticleID:  farticleID,                //
		IsDeleted:  a_db.IsDeleted,            //
		CreateTime: a_db.CreateTime.ToTime8(), //
		MTime:      a_db.MTime.ToTime8(),      //
		Recommend:  a_db.Recommend,            //
		NComments:  a_db.NComments,            //
		Owner:      a_db.Owner,                //
		Title:      fTitle,      //
		Money:      a_db.Money,                //
		Filemode:   a_db.Filemode,             //
		Class:      a_db.Class,                //

		URL: url, //

		Idx: a_db.Idx,

		Rank: a_db.Rank,

		SubjectType: a_db.SubjectType,
	}
}

func NewArticleSummaryFromWithRegex(a_db *schema.ArticleSummaryWithRegex) *ArticleSummary {
	fboardID := ToFBoardID(a_db.BBoardID)
	// if article is deleted, hide articleId
	// https://github.com/Ptt-official-app/go-openbbsmiddleware/issues/253#issuecomment-971526173
	var fTitle string
	var farticleID FArticleID
	if a_db.IsDeleted {
		farticleID = ""
		fTitle = "本文已被刪除"
	} else {
		fTitle =  ToFTitle(a_db.Title)
		farticleID = ToFArticleID(a_db.ArticleID)
	}

	url := ToURL(fboardID, farticleID)

	return &ArticleSummary{
		FBoardID:   fboardID,                  //
		ArticleID:  farticleID,                //
		IsDeleted:  a_db.IsDeleted,            //
		CreateTime: a_db.CreateTime.ToTime8(), //
		MTime:      a_db.MTime.ToTime8(),      //
		Recommend:  a_db.Recommend,            //
		NComments:  a_db.NComments,            //
		Owner:      a_db.Owner,                //
		Title:      fTitle,      //
		Money:      a_db.Money,                //
		Filemode:   a_db.Filemode,             //
		Class:      a_db.Class,                //

		URL: url, //

		Idx: a_db.Idx,

		Rank: a_db.Rank,

		SubjectType: a_db.SubjectType,
	}
}
