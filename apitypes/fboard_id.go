package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

// FBoardID as board-id for frontend.
type FBoardID string

func ToFBoardID(boardID bbs.BBoardID) FBoardID {
	return FBoardID(boardID.ToBrdname())
}

func (f FBoardID) ToBBoardID() (boardID bbs.BBoardID, err error) {
	return schema.GetBoardID(string(f))
}
