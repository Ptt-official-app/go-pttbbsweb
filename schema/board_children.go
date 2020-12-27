package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	BoardChildren_c *db.Collection
)

type BoardChildren struct {
	//板的子板名單

	ParentID bbs.BBoardID `bson:"parent_id"`
	ChildID  bbs.BBoardID `bson:"child_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
