package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ContentInfo
type ArticleContentInfo struct {
	ContentMD5 string `bson:"content_md5"`

	Content             [][]*types.Rune `bson:"content"`
	IP                  string          `bson:"ip"`
	Host                string          `bson:"host"` // ip 的中文呈現, 外國則為國家.
	BBS                 string          `bson:"bbs"`
	ContentUpdateNanoTS types.NanoTS    `bson:"content_update_nano_ts"`

	SignatureDBCS []byte `bson:"signature_dbcs"`
	SignatureMD5  string `bson:"signature_md5"`

	IsDeleted bool `bson:"deleted,omitempty"` //
}

var (
	EMPTY_ARTICLE_CONTENT_INFO = &ArticleContentInfo{}
	articleContentInfoFields   = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_INFO)
)

func GetArticleContentInfo(bboardID bbs.BBoardID, articleID bbs.ArticleID) (contentInfo *ArticleContentInfo, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	contentInfo = &ArticleContentInfo{}
	err = Article_c.FindOne(query, &contentInfo, articleContentInfoFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return contentInfo, nil
}

func UpdateArticleContentInfo(bboardID bbs.BBoardID, articleID bbs.ArticleID, contentInfo *ArticleContentInfo) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:  bboardID,
		ARTICLE_ARTICLE_ID_b: articleID,
	}

	r, err := Article_c.CreateOnly(query, contentInfo)
	if err != nil {
		return err
	}
	if r.UpsertedCount > 0 {
		return nil
	}

	query = bson.M{
		"$or": bson.A{
			bson.M{
				ARTICLE_BBOARD_ID_b:  bboardID,
				ARTICLE_ARTICLE_ID_b: articleID,
				ARTICLE_CONTENT_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				ARTICLE_BBOARD_ID_b:  bboardID,
				ARTICLE_ARTICLE_ID_b: articleID,
				ARTICLE_CONTENT_UPDATE_NANO_TS_b: bson.M{
					"$lt": contentInfo.ContentUpdateNanoTS,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	r, err = Article_c.UpdateOneOnly(query, contentInfo)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}
	return nil
}
