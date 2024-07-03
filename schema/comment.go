package schema

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

var Comment_c *db.Collection

type Comment struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`

	//XXX currently it's very hard to maintain the comment-id.
	//if we do comment-id only based on MD5:
	//  got duplicated md5-id if the owner posts the same comments
	//  within 1 min.
	//
	//if we add the inferred CreateTime into the comment-id:
	//  the CreateTime may be changed if the author deletes
	//  some other comments within same minute.
	CommentID    types.CommentID     `bson:"cid"`
	TheType      ptttype.CommentType `bson:"type"`
	RefIDs       []types.CommentID   `bson:"refids"`
	IsDeleted    bool                `bson:"deleted,omitempty"`
	DeleteReason string              `bson:"delete_reason,omitempty"`
	CreateTime   types.NanoTS        `bson:"create_time_nano_ts"`
	Owner        bbs.UUserID         `bson:"owner"`
	Content      [][]*types.Rune     `bson:"content"` // content in comment is colorless.
	IP           string              `bson:"ip"`
	Host         string              `bson:"host"` // ip 的中文呈現, 外國則為國家.
	MD5          string              `bson:"md5"`

	FirstCreateTime    types.NanoTS `bson:"first_create_time_nano_ts,omitempty"`    // create-time from first-comments.
	InferredCreateTime types.NanoTS `bson:"inferred_create_time_nano_ts,omitempty"` // create-time from inferred.
	NewCreateTime      types.NanoTS `bson:"new_create_time_nano_ts,omitempty"`      // create-time from new comment.

	SortTime types.NanoTS `bson:"sort_time_nano_ts"`

	TheDate string `bson:"the_date"`
	DBCS    []byte `json:"-" bson:"dbcs"`
	DBCSStr string `json:"-" bson:"dbcsstr"`

	EditNanoTS types.NanoTS `bson:"edit_nano_ts"` // for reply.

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_COMMENT = &Comment{}

var (
	COMMENT_BBOARD_ID_b     = getBSONName(EMPTY_COMMENT, "BBoardID")
	COMMENT_ARTICLE_ID_b    = getBSONName(EMPTY_COMMENT, "ArticleID")
	COMMENT_COMMENT_ID_b    = getBSONName(EMPTY_COMMENT, "CommentID")
	COMMENT_THE_TYPE_b      = getBSONName(EMPTY_COMMENT, "TheType")
	COMMENT_REF_IDS_b       = getBSONName(EMPTY_COMMENT, "RefIDs")
	COMMENT_IS_DELETED_b    = getBSONName(EMPTY_COMMENT, "IsDeleted")
	COMMENT_DELETE_REASON_b = getBSONName(EMPTY_COMMENT, "DeleteReason")
	COMMENT_CREATE_TIME_b   = getBSONName(EMPTY_COMMENT, "CreateTime")
	COMMENT_OWNER_b         = getBSONName(EMPTY_COMMENT, "Owner")
	COMMENT_CONTENT_b       = getBSONName(EMPTY_COMMENT, "Content")
	COMMENT_IP_b            = getBSONName(EMPTY_COMMENT, "IP")
	COMMENT_HOST_b          = getBSONName(EMPTY_COMMENT, "Host")
	COMMENT_MD5_b           = getBSONName(EMPTY_COMMENT, "MD5")

	COMMENT_FIRST_CREATE_TIME_b    = getBSONName(EMPTY_COMMENT, "FirstCreateTime")
	COMMENT_INFERRED_CREATE_TIME_b = getBSONName(EMPTY_COMMENT, "InferredCreateTime")
	COMMENT_NEW_CREATE_TIME_b      = getBSONName(EMPTY_COMMENT, "NewCreateTime")

	COMMENT_SORT_TIME_b = getBSONName(EMPTY_COMMENT, "SortTime")

	COMMENT_THE_DATE_b = getBSONName(EMPTY_COMMENT, "TheDate")
	COMMENT_DBCS_b     = getBSONName(EMPTY_COMMENT, "DBCS")

	COMMENT_EDIT_NANO_TS_b = getBSONName(EMPTY_COMMENT, "EditNanoTS")

	COMMENT_UPDATE_NANO_TS_b = getBSONName(EMPTY_COMMENT, "UpdateNanoTS")
)

func assertCommentFields() error {
	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_ARTICLE_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_SUMMARY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_IS_DELETED); err != nil {
		return err
	}

	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_MD5); err != nil {
		return err
	}

	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_DBCS); err != nil {
		return err
	}

	if err := assertFields(EMPTY_COMMENT, EMPTY_COMMENT_SORT_TIME); err != nil {
		return err
	}

	return nil
}

type CommentQuery struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	ArticleID bbs.ArticleID   `bson:"aid"`
	CommentID types.CommentID `bson:"cid"`
}

var EMPTY_COMMENT_QUERY = &CommentQuery{}

type CommentArticleQuery struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`
}

var EMPTY_COMMENT_ARTICLE_QUERY = &CommentArticleQuery{}

func GetCommentMapByCommentIDs(boardID bbs.BBoardID, articleID bbs.ArticleID, commentIDs []types.CommentID) (commentMap map[types.CommentID]*Comment, err error) {
	query := bson.M{
		COMMENT_BBOARD_ID_b:  boardID,
		COMMENT_ARTICLE_ID_b: articleID,
		COMMENT_COMMENT_ID_b: bson.M{
			"$in": commentIDs,
		},
	}

	// find
	var comments []*Comment
	err = Comment_c.Find(query, 0, &comments, nil, nil)
	if err != nil {
		return nil, err
	}

	commentMap = make(map[types.CommentID]*Comment)
	for _, each := range comments {
		commentMap[each.CommentID] = each
	}

	return commentMap, nil
}

// GetComments
func GetComments(boardID bbs.BBoardID, articleID bbs.ArticleID, sortNanoTS types.NanoTS, commentID types.CommentID, descending bool, limit int) (comments []*Comment, err error) {
	// setup query
	var query bson.M
	if sortNanoTS == 0 {
		query = bson.M{
			COMMENT_BBOARD_ID_b:  boardID,
			COMMENT_ARTICLE_ID_b: articleID,
			COMMENT_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	} else {
		theDirCommentID := "$gte"
		if descending {
			theDirCommentID = "$lte"
		}

		query = bson.M{
			COMMENT_BBOARD_ID_b:  boardID,
			COMMENT_ARTICLE_ID_b: articleID,
			COMMENT_SORT_TIME_b: bson.M{
				theDirCommentID: sortNanoTS,
			},
			COMMENT_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	}
	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: COMMENT_SORT_TIME_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: COMMENT_SORT_TIME_b, Value: 1},
		}
	}

	// find
	err = Comment_c.Find(query, int64(limit), &comments, nil, sortOpts)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func CountComments(boardID bbs.BBoardID, articleID bbs.ArticleID) (nComments int, err error) {
	query := bson.M{
		COMMENT_BBOARD_ID_b:  boardID,
		COMMENT_ARTICLE_ID_b: articleID,
		COMMENT_IS_DELETED_b: bson.M{
			"$exists": false,
		},
	}

	count, err := Comment_c.Count(query, 0)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// UpdateComments
//
// XXX hack in updateCommentsCore:
//
//	treat all the comments as non-deleted and unset IsDeleted.
func UpdateComments(comments []*Comment, updateNanoTS types.NanoTS) (err error) {
	if len(comments) == 0 {
		return nil
	}

	p_comments := comments

	var first []*Comment

	for block := getBlock(len(p_comments), MAX_COMMENT_BLOCK); len(p_comments) > 0; block = getBlock(len(p_comments), MAX_COMMENT_BLOCK) {
		first, p_comments = p_comments[:block], p_comments[block:]

		err = updateCommentsCore(first, updateNanoTS)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateCommentsCore(comments []*Comment, updateNanoTS types.NanoTS) (err error) {
	theList := make([]*db.UpdatePair, len(comments))
	for idx, each := range comments {
		filter := &CommentQuery{
			BBoardID:  each.BBoardID,
			ArticleID: each.ArticleID,
			CommentID: each.CommentID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: filter,
			Update: each,
		}
	}
	r, err := Comment_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(comments)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateComments := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*CommentQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					COMMENT_BBOARD_ID_b:  origFilter.BBoardID,
					COMMENT_ARTICLE_ID_b: origFilter.ArticleID,
					COMMENT_COMMENT_ID_b: origFilter.CommentID,
					COMMENT_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},
				},
				bson.M{
					COMMENT_BBOARD_ID_b:  origFilter.BBoardID,
					COMMENT_ARTICLE_ID_b: origFilter.ArticleID,
					COMMENT_COMMENT_ID_b: origFilter.CommentID,
					COMMENT_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},
				},
			},
		}

		each.Filter = filter
		origUpdate := each.Update
		each.Update = bson.M{
			"$set": origUpdate,
			"$unset": bson.M{
				COMMENT_IS_DELETED_b: true,
			},
		}

		updateComments = append(updateComments, each)
	}
	_, err = Comment_c.BulkUpdateOneOnlyNoSet(updateComments)

	return err
}

func (c *Comment) CleanComment() {
	trimStr := " \t\r"
	if c.TheType == ptttype.COMMENT_TYPE_REPLY {
		trimStr = "\t\r"
	}

	for _, each := range c.Content {
		lenEach := len(each)
		if lenEach == 0 {
			continue
		}
		lastRune := each[lenEach-1]
		lastRune.Utf8 = strings.TrimRight(lastRune.Utf8, trimStr)
	}
}

func (c *Comment) CleanReply() {
	lenContent := len(c.Content)

	cleanedReply := make([][]*types.Rune, lenContent)
	for idx, each := range c.Content {
		cleanedReply[idx] = cleanReplyPerLine(each)
	}

	idxFirstGoodContent := 0
	for idx, each := range cleanedReply {
		if len(each) != 0 {
			idxFirstGoodContent = idx
			break
		}
	}

	idxLastGoodContent := 0
	for idx := lenContent - 1; idx >= 0; idx-- {
		each := cleanedReply[idx]
		if len(each) != 0 {
			idxLastGoodContent = idx + 1
			break
		}
	}
	if idxFirstGoodContent >= idxLastGoodContent {
		c.Content = nil
		return
	}

	c.Content = c.Content[idxFirstGoodContent:idxLastGoodContent]
}

func cleanReplyPerLine(origLine []*types.Rune) (newLine []*types.Rune) {
	count := 0
	for _, each := range origLine {
		count += len(each.Utf8)
	}
	if count == 0 {
		return nil
	}

	return origLine
}

func (c *Comment) SetSortTime(sortTime types.NanoTS) {
	c.SortTime = sortTime
	if c.TheType == ptttype.COMMENT_TYPE_REPLY && types.IsReplyID(c.CommentID) { // special treat for reply
		return
	}
	c.CommentID = types.ToCommentID(sortTime, c.MD5)
}
