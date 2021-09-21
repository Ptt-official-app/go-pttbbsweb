package apitypes

import (
	"bytes"

	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type DeleteCommentParams struct {
	CommentID types.CommentID `json:"cid" form:"cid" url:"cid"`
	Reason    string          `json:"reason" form:"reason" url:"reason"`
}

func (d *DeleteCommentParams) ToComment(comment *schema.Comment, userID bbs.UUserID, remoteAddr string, updateNanoTS types.NanoTS) (newComment *schema.Comment) {
	comment.Content = [][]*types.Rune{
		{
			{
				Utf8:   string(userID) + " 刪除 " + string(comment.Owner) + " 的推文: " + d.Reason,
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	theDBCS := dbcs.Utf8ToDBCS(comment.Content)
	// hack to remove '\r'
	lastLine := comment.Content[len(comment.Content)-1]
	lastLineDBCS := lastLine[len(lastLine)-1].DBCS
	if lastLineDBCS[len(lastLineDBCS)-1] == '\r' {
		lastLineDBCS = lastLineDBCS[:len(lastLineDBCS)-1]
	}
	lastLine[len(lastLine)-1].DBCS = lastLineDBCS
	comment.Content[len(comment.Content)-1] = lastLine

	lastDBCS := theDBCS[len(theDBCS)-1]
	if lastDBCS[len(lastDBCS)-1] == '\r' {
		lastDBCS = lastDBCS[:len(lastDBCS)-1]
	}
	theDBCS[len(theDBCS)-1] = lastDBCS

	theDBCS2 := make([][]byte, 0, len(theDBCS)+3)
	theDBCS2 = append(theDBCS2, dbcs.MATCH_COMMENT_DELETED_PREFIX)
	theDBCS2 = append(theDBCS2, theDBCS...)
	theDBCS2 = append(theDBCS2, dbcs.MATCH_COMMENT_DELETED_POSTFIX)
	if types.IS_CARRIAGE_RETURN {
		theDBCS2 = append(theDBCS2, []byte{'\r'})
	}
	comment.DBCS = bytes.Join(theDBCS2, []byte{})
	comment.MD5 = dbcs.Md5sum(comment.DBCS)
	comment.DeleteReason = d.Reason
	comment.UpdateNanoTS = updateNanoTS
	comment.TheType = ptttype.COMMENT_TYPE_DELETED
	comment.TheDate = ""

	return comment
}
