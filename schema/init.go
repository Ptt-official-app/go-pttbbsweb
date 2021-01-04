package schema

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	redis "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
)

func Init() (err error) {
	if client != nil {
		return nil
	}
	client, err = db.NewClient(MONGO_PROTOCOL, MONGO_HOST, MONGO_PORT, MONGO_DBNAME)
	if err != nil {
		return err
	}

	//access-token
	AccessToken_c = client.Collection("access_token")

	keys := &bson.D{{Key: ACCESS_TOKEN_USER_ID_b, Value: 1}}
	err = AccessToken_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	//article
	Article_c = client.Collection("article")
	keys = &bson.D{
		{Key: ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: ARTICLE_ARTICLE_ID_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	//board
	Board_c = client.Collection("board")
	keys = &bson.D{
		{Key: BOARD_BBOARD_ID_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	//BoardBanuser
	BoardBanuser_c = client.Collection("board_banuser")

	//BoardChildren
	BoardChildren_c = client.Collection("board_children")

	//BoardFriend
	BoardFriend_c = client.Collection("board_friend")

	//Client
	Client_c = client.Collection("client")

	//comment
	Comment_c = client.Collection("comment")
	keys = &bson.D{
		{Key: COMMENT_BBOARD_ID_b, Value: 1},
		{Key: COMMENT_ARTICLE_ID_b, Value: 1},
		{Key: COMMENT_COMMENT_ID_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}
	keys = &bson.D{
		{Key: COMMENT_BBOARD_ID_b, Value: 1},
		{Key: COMMENT_ARTICLE_ID_b, Value: 1},
		{Key: COMMENT_UPDATE_NANO_TS_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	//User
	User_c = client.Collection("user")

	//UserAloha
	UserAloha_c = client.Collection("user_aloha")

	//UserFriend
	UserFriend_c = client.Collection("user_friend")

	//UserReadArticle
	UserReadArticle_c = client.Collection("user_read_article")

	//UserReadBoard
	UserReadBoard_c = client.Collection("user_read_board")

	//UserReject
	UserReject_c = client.Collection("user_reject")

	//assert-all-fields
	err = assertAllFields()
	if err != nil {
		return err
	}

	//redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}

//Close
//
//XXX do not really close to avoid db connection-error in tests.
func Close() (err error) {
	return nil
	/*
		err = client.Close()
		if err != nil {
			log.Errorf("schema.Close: unable to close mongo: e: %v", err)
		}

		client = nil
		Client_c = nil
		User_c = nil
		AccessToken_c = nil
		UserReadArticle_c = nil
		UserReadBoard_c = nil

		return nil
	*/
}

func getBSONName(empty interface{}, fieldName string) string {
	val := reflect.ValueOf(empty).Elem().Interface()
	theType := reflect.TypeOf(val)
	field, ok := theType.FieldByName(fieldName)
	if !ok {
		return ""
	}
	return strings.Split(field.Tag.Get("bson"), ",")[0]
}

func getFields(empty interface{}, fields_i interface{}) (fields map[string]bool) {
	fieldsType := reflect.ValueOf(fields_i).Elem().Type()
	nFieldNames := fieldsType.NumField()
	fieldNames := make([]string, nFieldNames)
	for idx := 0; idx < nFieldNames; idx++ {
		fieldNames[idx] = fieldsType.Field(idx).Name
	}

	fields = make(map[string]bool)
	for _, each := range fieldNames {
		bsonName := getBSONName(empty, each)
		fields[bsonName] = true
	}

	return fields
}

func assertFields(empty interface{}, fields_i interface{}) error {
	emptyType := reflect.ValueOf(empty).Elem().Type()
	fieldsType := reflect.ValueOf(fields_i).Elem().Type()
	nFieldNames := fieldsType.NumField()
	for idx := 0; idx < nFieldNames; idx++ {
		fieldName := fieldsType.Field(idx).Name
		emptyBSONName := getBSONName(empty, fieldName)
		if emptyBSONName == "" {
			return errors.New(fmt.Sprintf("invalid fieldName: (%v/%v): field: %v", fieldsType.Name(), emptyType.Name(), fieldName))
		}
		bsonName := getBSONName(fields_i, fieldName)
		if emptyBSONName != bsonName {
			return errors.New(fmt.Sprintf("invalid bsonName: (%v/%v): field: %v bson: %v expected: %v ", fieldsType.Name(), emptyType.Name(), fieldName, bsonName, emptyBSONName))
		}
	}

	return nil
}

func assertAllFields() error {
	if err := assertArticleFields(); err != nil {
		return err
	}

	if err := assertBoardFields(); err != nil {
		return err
	}

	if err := assertCommentFields(); err != nil {
		return err
	}

	if err := assertUserReadArticleFields(); err != nil {
		return err
	}

	if err := assertUserReadBoardFields(); err != nil {
		return err
	}

	if err := assertUserFields(); err != nil {
		return err
	}

	return nil
}
