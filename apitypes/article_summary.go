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
}

func NewArticleSummary(a_db *schema.ArticleSummary) *ArticleSummary {
	fboardID := ToFBoardID(a_db.BBoardID)
	farticleID := ToFArticleID(a_db.ArticleID)

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
		Title:      a_db.Title,                //
		Money:      a_db.Money,                //
		Filemode:   a_db.Filemode,             //
		Class:      a_db.Class,                //

		URL: url, //

		Idx: a_db.Idx,

		Rank: a_db.Rank,
	}
}
