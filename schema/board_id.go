package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type BoardID struct {
	BBoardID bbs.BBoardID `bson:"bid"`
	Brdname  string       `bson:"brdname"`
}

var (
	EMPTY_BOARD_ID = &BoardID{}
	boardIDFields  = getFields(EMPTY_BOARD, EMPTY_BOARD_ID)
)

func GetBoardID(brdname string) (boardID bbs.BBoardID, err error) {
	query := bson.M{
		BOARD_BRDNAME_b: brdname,
	}

	result := &BoardID{}
	err = Board_c.FindOne(query, result, boardIDFields)
	if err != nil {
		logrus.Warnf("schema.GetBoardID: unable to FindOne: query: %v e: %v", query, err)
		return "", err
	}

	return result.BBoardID, nil
}

func GetBoardIDByBid(bid ptttype.Bid) (boardID bbs.BBoardID, err error) {
	query := bson.M{
		BOARD_BID_b: bid,
	}

	result := &BoardID{}
	err = Board_c.FindOne(query, result, boardIDFields)
	if err != nil {
		logrus.Warnf("schema.GetBoardID: unable to FindOne: query: %v e: %v", query, err)
		return "", err
	}

	return result.BBoardID, nil
}
