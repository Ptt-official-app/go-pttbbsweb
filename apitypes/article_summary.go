package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type ArticleSummary struct {
	BBoardID   bbs.BBoardID     `json:"bid"`         //0
	ArticleID  bbs.ArticleID    `json:"aid"`         //1
	IsDeleted  bool             `json:"deleted"`     //2
	CreateTime types.Time8      `json:"create_time"` //3
	MTime      types.Time8      `json:"modified"`    //4
	Recommend  int              `json:"recommend"`   //5
	Owner      bbs.UUserID      `json:"owner"`       //6
	Title      string           `json:"title"`       //7
	Money      int              `json:"money"`       //8
	Class      string           `json:"class"`       //can be: R: 轉, [class]
	Filemode   ptttype.FileMode `json:"mode"`        //10

	URL  string `json:"url"`  //11
	Read bool   `json:"read"` //12
}

func NewArticleSummary(a_db *schema.ArticleSummary) *ArticleSummary {
	url := ToURL(a_db.BBoardID, a_db.ArticleID)

	return &ArticleSummary{
		BBoardID:   a_db.BBoardID,             //0
		ArticleID:  a_db.ArticleID,            //1
		IsDeleted:  a_db.IsDeleted,            //2
		CreateTime: a_db.CreateTime.ToTime8(), //3
		MTime:      a_db.MTime.ToTime8(),      //4
		Recommend:  a_db.Recommend,            //5
		Owner:      a_db.Owner,                //6
		Title:      a_db.Title,                //7
		Money:      a_db.Money,                //8
		Filemode:   a_db.Filemode,             //9
		Class:      a_db.Class,                //10

		URL: url, //11
	}
}

func (a *ArticleSummary) Serialize() *schema.ArticleSummary {
	return &schema.ArticleSummary{
		BBoardID:   a.BBoardID,              //0
		ArticleID:  a.ArticleID,             //1
		IsDeleted:  a.IsDeleted,             //2
		CreateTime: a.CreateTime.ToNanoTS(), //3
		MTime:      a.MTime.ToNanoTS(),      //4
		Recommend:  a.Recommend,             //5
		Owner:      a.Owner,                 //6
		Title:      a.Title,                 //7
		Money:      a.Money,                 //8
		Filemode:   a.Filemode,              //9
		Class:      a.Class,                 //10
	}
}
