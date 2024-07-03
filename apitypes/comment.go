package apitypes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

type Comment struct {
	FBoardID   FBoardID            `json:"bid"`
	FArticleID FArticleID          `json:"aid"`
	CommentID  types.CommentID     `json:"cid"`
	TheType    ptttype.CommentType `json:"type"`
	RefID      types.CommentID     `json:"refid"`
	IsDeleted  bool                `json:"deleted"`
	CreateTime types.Time8         `json:"create_time"`
	SortTime   types.Time8         `json:"sort_time"`
	Owner      bbs.UUserID         `json:"owner"`
	Content    [][]*types.Rune     `json:"content"`
	IP         string              `json:"ip"`
	Host       string              `json:"host"` // ip 的中文呈現, 外國則為國家.
	Idx        string              `json:"idx"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func NewComment(comment_db *schema.Comment, userID bbs.UUserID) (comment *Comment) {
	var refID types.CommentID
	if len(comment_db.RefIDs) > 0 {
		refID = comment_db.RefIDs[0]
	}
	comment = &Comment{
		FBoardID:   ToFBoardID(comment_db.BBoardID),
		FArticleID: ToFArticleID(comment_db.ArticleID),
		CommentID:  comment_db.CommentID,
		TheType:    comment_db.TheType,
		RefID:      refID,
		IsDeleted:  comment_db.IsDeleted,
		CreateTime: comment_db.CreateTime.ToTime8(),
		SortTime:   comment_db.SortTime.ToTime8(),
		Owner:      comment_db.Owner,
		Content:    comment_db.Content,
		IP:         comment_db.IP,
		Host:       comment_db.Host,
		Idx:        SerializeCommentIdx(comment_db.SortTime, comment_db.CommentID),

		TokenUser: userID,
	}

	return comment
}

func DeserializeCommentIdx(theIdx string) (sortNanoTS types.NanoTS, commentID types.CommentID) {
	theList := strings.Split(theIdx, "@")
	if len(theList) != 2 {
		return 0, ""
	}

	nanoTSInt, err := strconv.Atoi(theList[0])
	if err != nil {
		return 0, ""
	}

	return types.NanoTS(nanoTSInt), types.CommentID(theList[1])
}

func SerializeCommentIdx(sortNanoTS types.NanoTS, commentID types.CommentID) string {
	return fmt.Sprintf("%v@%v", sortNanoTS, commentID)
}
