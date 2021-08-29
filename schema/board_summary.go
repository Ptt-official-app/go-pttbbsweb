package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BoardSummary
type BoardSummary struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	Brdname   string          `bson:"brdname"`
	Title     string          `bson:"title"`
	BrdAttr   ptttype.BrdAttr `bson:"flag"`
	BoardType string          `bson:"the_type"`
	Category  string          `bson:"class"`
	NUser     int             `bson:"nuser"`
	BMs       []bbs.UUserID   `bson:"bms"`
	Total     int             `bson:"total"` /* total articles, 需要即時知道. 因為 read 頻率高. 並且跟 last-post-time-ts 一樣 write 頻率 << read 頻率 */

	LastPostTime types.NanoTS `bson:"last_post_time_nano_ts"` /* 需要即時知道來做板的已讀 */

	IsDeleted bool `bson:"deleted,omitempty"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	Gid        ptttype.Bid `bson:"pttgid"`
	Bid        ptttype.Bid `bson:"pttbid"`
	IdxByName  string      `bson:"pttidxname"`
	IdxByClass string      `bson:"pttidxclass"`
}

var (
	EMPTY_BOARD_SUMMARY = &BoardSummary{}
	boardSummaryFields  = getFields(EMPTY_BOARD, EMPTY_BOARD_SUMMARY)
)

func NewBoardSummary(b_b *bbs.BoardSummary, updateNanoTS types.NanoTS) *BoardSummary {
	return &BoardSummary{
		BBoardID:  b_b.BBoardID,
		Brdname:   b_b.Brdname,
		Title:     types.Big5ToUtf8(b_b.RealTitle),
		BrdAttr:   b_b.BrdAttr,
		BoardType: types.Big5ToUtf8(b_b.BoardType),
		Category:  types.Big5ToUtf8(b_b.BoardClass),
		BMs:       b_b.BM,
		Total:     int(b_b.Total),
		NUser:     int(b_b.NUser),

		LastPostTime: types.Time4ToNanoTS(b_b.LastPostTime),

		UpdateNanoTS: updateNanoTS,

		Gid:        b_b.Gid,
		Bid:        b_b.Bid,
		IdxByName:  b_b.IdxByName,
		IdxByClass: b_b.IdxByClass,
	}
}

func UpdateBoardSummaries(boardSummaries []*BoardSummary, updateNanoTS types.NanoTS) (err error) {
	if len(boardSummaries) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(boardSummaries))
	for idx, each := range boardSummaries {
		query := &BoardQuery{
			BBoardID: each.BBoardID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := Board_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(boardSummaries)) { // all are created
		return nil
	}

	// update items with comparing update-nano-ts
	upsertedIDs := r.UpsertedIDs
	updateBoardSummaries := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*BoardQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					BOARD_BBOARD_ID_b: origFilter.BBoardID,
					BOARD_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},

					BOARD_IS_DELETED_b: bson.M{"$exists": false},
				},
				bson.M{
					BOARD_BBOARD_ID_b: origFilter.BBoardID,
					BOARD_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},

					BOARD_IS_DELETED_b: bson.M{"$exists": false},
				},
			},
		}
		each.Filter = filter
		updateBoardSummaries = append(updateBoardSummaries, each)
	}

	_, err = Board_c.BulkUpdateOneOnly(updateBoardSummaries)

	return err
}

func GetBoardSummary(bboardID bbs.BBoardID) (result *BoardSummary, err error) {
	query := &BoardQuery{
		BBoardID: bboardID,
	}

	result = &BoardSummary{}
	err = Board_c.FindOne(query, result, boardSummaryFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetBoardSummariesByClsID(clsID ptttype.Bid, startIdx string, isAsc bool, limit int, sortBy ptttype.BSortBy) (boardSummaries []*BoardSummary, err error) {
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

	err = Board_c.Find(query, int64(limit), &boardSummaries, boardSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return boardSummaries, nil
}
