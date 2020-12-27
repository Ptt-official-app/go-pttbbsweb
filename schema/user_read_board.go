package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	UserReadBoard_c *db.Collection
)

type UserReadBoard struct {
	//已讀板紀錄

	UserID       bbs.UUserID  `bson:"user_id"`
	BBoardID     bbs.BBoardID `bson:"bid"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var (
	EMPTY_USER_READ_BOARD = &UserReadBoard{}
)

var (
	USER_READ_BOARD_USER_ID_b        = getBSONName(EMPTY_USER_READ_BOARD, "UserID")
	USER_READ_BOARD_BBOARD_ID_b      = getBSONName(EMPTY_USER_READ_BOARD, "BBoardID")
	USER_READ_BOARD_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_READ_BOARD, "UpdateNanoTS")
)

func assertUserReadBoardFields() error {
	if err := assertFields(EMPTY_USER_READ_BOARD, EMPTY_USER_READ_BOARD_QUERY); err != nil {
		return err
	}

	return nil
}

type UserReadBoardQuery struct {
	UserID   bbs.UUserID  `bson:"user_id"`
	BBoardID bbs.BBoardID `bson:"bid"`
}

var (
	EMPTY_USER_READ_BOARD_QUERY = &UserReadBoardQuery{}
)

func UpdateUserReadBoard(userReadBoard *UserReadBoard) (err error) {
	query := bson.M{
		USER_READ_BOARD_USER_ID_b:   userReadBoard.UserID,
		USER_READ_BOARD_BBOARD_ID_b: userReadBoard.BBoardID,
	}

	r, err := UserReadBoard_c.CreateOnly(query, userReadBoard)
	if err != nil {
		return err
	}
	if r.UpsertedCount == 1 { //created, no need to update.
		return nil
	}

	query[USER_READ_BOARD_UPDATE_NANO_TS_b] = bson.M{
		"$lt": userReadBoard.UpdateNanoTS,
	}

	_, err = UserReadBoard_c.UpdateOneOnly(query, userReadBoard)

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

	r, err := UserReadBoard_c.BulkCreateOnly(theList)
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

		origFilter, ok := each.Filter.(*UserReadBoardQuery)
		filter := bson.M{
			USER_READ_BOARD_USER_ID_b:   origFilter.UserID,
			USER_READ_BOARD_BBOARD_ID_b: origFilter.BBoardID,
			USER_READ_BOARD_UPDATE_NANO_TS_b: bson.M{
				"$lt": updateNanoTS,
			},
		}
		each.Filter = filter
		updateUserReadBoards = append(updateUserReadBoards, each)
	}

	_, err = UserReadBoard_c.BulkUpdateOneOnly(updateUserReadBoards)

	return err
}

func FindUserReadBoards(userID bbs.UUserID, boardIDs []bbs.BBoardID) ([]*UserReadBoard, error) {

	//query
	query := bson.M{
		USER_READ_BOARD_USER_ID_b: userID,
		USER_READ_BOARD_BBOARD_ID_b: bson.M{
			"$in": boardIDs,
		},
	}

	var dbResults []*UserReadBoard
	err := UserReadBoard_c.Find(query, 0, &dbResults, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}
