package apitypes

import (
	"bytes"

	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type ReplyCommentParams struct {
	CommentID types.CommentID `json:"cid" form:"cid" url:"cid"`
	Content   [][]*types.Rune `json:"content" form:"content" url:"content"`
}

// ToComment
//
// referring to:
//
//	dbcs.parseReply
//	dbcs.EDBlock.ForwardInferTS
func (r *ReplyCommentParams) ToComment(userID bbs.UUserID, remoteAddr string, boardID bbs.BBoardID, articleID bbs.ArticleID, commentSortTime types.NanoTS, updateNanoTS types.NanoTS) (replyComment *schema.Comment) {
	color := &types.Color{}
	*color = types.DefaultColor

	replyDBCSs := dbcs.Utf8ToDBCS(r.Content)
	replyDBCS := bytes.Join(replyDBCSs, []byte{'\n'})

	replySortTime := commentSortTime + dbcs.REPLY_STEP_NANO_TS

	md5sum := dbcs.Md5sum(replyDBCS)

	replyComment = &schema.Comment{
		BBoardID:   boardID,
		ArticleID:  articleID,
		CommentID:  types.ToReplyID(r.CommentID),
		TheType:    ptttype.COMMENT_TYPE_REPLY,
		RefIDs:     []types.CommentID{r.CommentID},
		CreateTime: updateNanoTS,
		Owner:      userID,
		Content:    r.Content,
		IP:         remoteAddr,
		MD5:        md5sum,

		NewCreateTime: updateNanoTS,
		SortTime:      replySortTime,
		DBCS:          replyDBCS,

		EditNanoTS:   updateNanoTS,
		UpdateNanoTS: updateNanoTS,
	}
	replyComment.SetSortTime(replySortTime)

	return replyComment
}
