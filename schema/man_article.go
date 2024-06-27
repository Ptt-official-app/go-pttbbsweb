package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var ManArticle_c *db.Collection

type ManArticle struct {
	Version    int                `bson:"version"`
	BBoardID   bbs.BBoardID       `bson:"bid"`                 //
	ArticleID  types.ManArticleID `bson:"aid"`                 //
	LevelIdx   types.ManArticleID `bson:"level_idx"`           //
	IsDeleted  bool               `bson:"deleted,omitempty"`   //
	CreateTime types.NanoTS       `bson:"create_time_nano_ts"` //
	MTime      types.NanoTS       `bson:"mtime_nano_ts"`

	Title    string           `bson:"title"` //
	Filemode ptttype.FileMode `bson:"mode"`  //

	IsDir bool `bson:"is_dir"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"` // used by article-summary

	ContentMTime types.NanoTS    `bson:"content_mtime_nano_ts"` //
	ContentMD5   string          `bson:"content_md5"`
	ContentID    types.ContentID `bson:"content_id"`
	Content      [][]*types.Rune `bson:"content"`

	ContentUpdateNanoTS types.NanoTS `bson:"content_update_nano_ts"`

	Idx int `bson:"pttidx"`
}

var EMPTY_MAN_ARTICLE = &ManArticle{}

var ( // bson-name
	MAN_ARTICLE_BBOARD_ID_b   = getBSONName(EMPTY_MAN_ARTICLE, "BBoardID")
	MAN_ARTICLE_ARTICLE_ID_b  = getBSONName(EMPTY_MAN_ARTICLE, "ArticleID")
	MAN_ARTICLE_LEVEL_IDX_b   = getBSONName(EMPTY_MAN_ARTICLE, "LevelIdx")
	MAN_ARTICLE_IS_DELETED_b  = getBSONName(EMPTY_MAN_ARTICLE, "IsDeleted")
	MAN_ARTICLE_CREATE_TIME_b = getBSONName(EMPTY_MAN_ARTICLE, "CreateTime")
	MAN_ARTICLE_TITLE_b       = getBSONName(EMPTY_MAN_ARTICLE, "Title")
	MAN_ARTICLE_FILEMODE_b    = getBSONName(EMPTY_MAN_ARTICLE, "Filemode")
	MAN_ARTICLE_IS_DIR_b      = getBSONName(EMPTY_MAN_ARTICLE, "IsDir")

	MAN_ARTICLE_UPDATE_NANO_TS_b = getBSONName(EMPTY_MAN_ARTICLE, "UpdateNanoTS")
	MAN_ARTICLE_CONTENT_MTIME_b  = getBSONName(EMPTY_MAN_ARTICLE, "ContentMTime")
	MAN_ARTICLE_CONTENT_MD5_b    = getBSONName(EMPTY_MAN_ARTICLE, "ContentMD5")
	MAN_ARTICLE_CONTENT_b        = getBSONName(EMPTY_MAN_ARTICLE, "Content")
	MAN_ARTICLE_CONTENT_ID_b     = getBSONName(EMPTY_MAN_ARTICLE, "ContentID")

	MAN_ARTICLE_IP_b                     = getBSONName(EMPTY_MAN_ARTICLE, "IP")
	MAN_ARTICLE_HOST_b                   = getBSONName(EMPTY_MAN_ARTICLE, "Host")
	MAN_ARTICLE_BBS_b                    = getBSONName(EMPTY_MAN_ARTICLE, "BBS")
	MAN_ARTICLE_CONTENT_UPDATE_NANO_TS_b = getBSONName(EMPTY_MAN_ARTICLE, "ContentUpdateNanoTS")

	MAN_ARTICLE_IDX_b = getBSONName(EMPTY_MAN_ARTICLE, "Idx")
)

type ManArticleQuery struct {
	BBoardID bbs.BBoardID       `bson:"bid"`
	LevelIdx types.ManArticleID `bson:"level_idx"` //
	Idx      int                `bson:"pttidx"`
}

var EMPTY_MAN_ARTICLE_QUERY = &ManArticleQuery{}

func assertManArticleFields() error {
	if err := assertFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_SUMMARY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_DETAIL_SUMMARY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_CONTENT_INFO); err != nil {
		return err
	}

	return nil
}
