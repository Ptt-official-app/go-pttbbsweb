package fav

import (
	"unsafe"

	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type FavLine struct {
	Lid pttbbsfav.Lid `bson:"lid"`
}

const SIZE_OF_FAV_LINE = unsafe.Sizeof(FavLine{})

type FavFolder struct {
	Fid        pttbbsfav.Fid `bson:"fid"`
	Title      string        `bson:"title"`
	ThisFolder *Fav          `bson:"folder"`
}

const SIZE_OF_FAV_FOLDER = unsafe.Sizeof(FavFolder{})

type FavBoard struct {
	Bid       ptttype.Bid    `bson:"bid"`
	LastVisit int32          `bson:"-"` /* UNUSED */
	Attr      pttbbsfav.Favh `bson:"-"`
}

const SIZE_OF_FAV_BOARD = unsafe.Sizeof(FavBoard{})
