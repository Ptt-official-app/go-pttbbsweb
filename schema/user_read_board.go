package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserReadBoard struct {
	UserID           bbs.UUserID  `bson:"user_id"`
	BBoardID         bbs.BBoardID `bson:"bid"`
	ReadUpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_USER_READ_BOARD = &UserReadBoard{}

type UserReadBoardQuery struct {
	UserID   bbs.UUserID  `bson:"user_id"`
	BBoardID bbs.BBoardID `bson:"bid"`
}

var (
	EMPTY_USER_READ_BOARD_QUERY = &UserReadBoardQuery{}

	userReadBoardFields = getFields(EMPTY_USER_BOARD, EMPTY_USER_READ_BOARD)
)

func UpdateUserReadBoard(userReadBoard *UserReadBoard) (err error) {
	query := bson.M{
		USER_BOARD_USER_ID_b:   userReadBoard.UserID,
		USER_BOARD_BBOARD_ID_b: userReadBoard.BBoardID,
	}

	r, err := UserBoard_c.CreateOnly(query, userReadBoard)
	if err != nil {
		return err
	}
	if r.UpsertedCount == 1 { // created, no need to update.
		return nil
	}

	query[USER_BOARD_READ_UPDATE_NANO_TS_b] = bson.M{
		"$lt": userReadBoard.ReadUpdateNanoTS,
	}

	_, err = UserBoard_c.UpdateOneOnly(query, userReadBoard)

	return err
}

func UpdateUserReadBoards(userReadBoards []*UserReadBoard, updateNanoTS types.NanoTS) (err error) {
	if len(userReadBoards) == 0 {
		return nil
	}

	theList := make([]*db.UpdatePair, len(userReadBoards))
	for idx, each := range userReadBoards {
		query := &UserReadBoardQuery{UserID: each.UserID, BBoardID: each.BBoardID}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := UserBoard_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(userReadBoards)) {
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateUserReadBoards := make([]*db.UpdatePair, 0, len(userReadBoards))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*UserReadBoardQuery)
		filter := bson.M{
			USER_BOARD_USER_ID_b:   origFilter.UserID,
			USER_BOARD_BBOARD_ID_b: origFilter.BBoardID,
			USER_BOARD_READ_UPDATE_NANO_TS_b: bson.M{
				"$lt": updateNanoTS,
			},
		}
		each.Filter = filter
		updateUserReadBoards = append(updateUserReadBoards, each)
	}

	_, err = UserBoard_c.BulkUpdateOneOnly(updateUserReadBoards)

	return err
}

func FindUserReadBoards(userID bbs.UUserID, boardIDs []bbs.BBoardID) ([]*UserReadBoard, error) {
	// query
	query := bson.M{
		USER_BOARD_USER_ID_b: userID,
		USER_BOARD_BBOARD_ID_b: bson.M{
			"$in": boardIDs,
		},
	}

	var dbResults []*UserReadBoard
	err := UserBoard_c.Find(query, 0, &dbResults, userReadBoardFields, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}
