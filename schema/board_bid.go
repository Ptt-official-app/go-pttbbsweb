package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"go.mongodb.org/mongo-driver/bson"
)

type BoardBid struct {
	Bid ptttype.Bid `bson:"pttbid"`

	IdxByName  string `bson:"pttidxname"`
	IdxByClass string `bson:"pttidxclass"`
}

var (
	EMPTY_BOARD_BID = &BoardBid{}
	boardBidFields  = getFields(EMPTY_BOARD, EMPTY_BOARD_BID)
)

func GetBoardBidsByClsID(clsID ptttype.Bid, startIdx string, isAsc bool, limit int, sortBy ptttype.BSortBy) (boardBids []*BoardBid, err error) {
	idx := ""
	switch sortBy {
	case ptttype.BSORT_BY_NAME:
		idx = BOARD_IDX_BY_NAME_b
	case ptttype.BSORT_BY_CLASS:
		idx = BOARD_IDX_BY_CLASS_b
	}

	var query bson.M
	if startIdx == "" {
		query = bson.M{
			BOARD_GID_b: clsID,
			BOARD_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	} else {
		theDIR := "$lte"
		if isAsc {
			theDIR = "$gte"
		}

		query = bson.M{
			BOARD_GID_b: clsID,
			idx: bson.M{
				theDIR: startIdx,
			},
			BOARD_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	}

	// sort opts
	var sortOpts bson.D
	if isAsc {
		sortOpts = bson.D{
			{Key: idx, Value: 1},
		}
	} else {
		sortOpts = bson.D{
			{Key: idx, Value: -1},
		}
	}

	err = Board_c.Find(query, int64(limit), &boardBids, boardBidFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return boardBids, nil
}
