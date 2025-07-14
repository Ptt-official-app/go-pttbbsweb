package schema

import "go.mongodb.org/mongo-driver/bson"

type BoardIsPopular struct {
	IsPopular bool `bson:"is_popular"`
}

var EMPTY_BOARD_IS_POPULAR = &BoardIsPopular{}

func ResetBoardIsPopular() (err error) {
	filter := bson.M{
		BOARD_IS_POPULAR_b: true,
	}
	update := bson.M{
		BOARD_IS_POPULAR_b: false,
	}

	_, err = Board_c.UpdateManyOnly(filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GetPopularBoardSummaries() (boardSummaries []*BoardSummary, err error) {
	query := bson.M{
		BOARD_IS_POPULAR_b: true,
	}

	sortOpts := bson.D{
		{Key: BOARD_NUSER_b, Value: -1},
	}

	err = Board_c.Find(query, 0, &boardSummaries, boardSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return boardSummaries, nil
}
