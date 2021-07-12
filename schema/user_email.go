package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserEmail_c *db.Collection

type UserEmail struct {
	UserID bbs.UUserID `bson:"user_id"`
	Email  string      `bson:"email"`

	IsSet bool `bson:"is_set,omitempty"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_USER_EMAIL = &UserEmail{}

var (
	USER_EMAIL_USER_ID_b        = getBSONName(EMPTY_USER_EMAIL, "UserID")
	USER_EMAIL_EMAIL_b          = getBSONName(EMPTY_USER_EMAIL, "Email")
	USER_EMAIL_IS_SET_b         = getBSONName(EMPTY_USER_EMAIL, "IsSet")
	USER_EMAIL_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_EMAIL, "UpdateNanoTS")
)

func GetUserEmailByUserID(userID bbs.UUserID, updateNanoTS types.NanoTS) (userEmail *UserEmail, err error) {
	query := bson.M{
		USER_EMAIL_USER_ID_b: userID,
	}

	return getUserEmailCore(query, updateNanoTS)
}

func GetUserEmailByEmail(email string, updateNanoTS types.NanoTS) (userEmail *UserEmail, err error) {
	query := bson.M{
		USER_EMAIL_EMAIL_b: email,
	}

	return getUserEmailCore(query, updateNanoTS)
}

func getUserEmailCore(query bson.M, updateNanoTS types.NanoTS) (userEmail *UserEmail, err error) {
	userEmail = &UserEmail{}
	err = UserEmail_c.FindOne(query, userEmail, nil)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
		}
		return nil, err
	}

	if userEmail.IsSet {
		return userEmail, nil
	}

	if updateNanoTS-userEmail.UpdateNanoTS < types.EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS {
		return userEmail, nil
	}

	// to remove query.
	err = UserEmail_c.Remove(query)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func CreateUserEmail(userID bbs.UUserID, email string, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		USER_EMAIL_USER_ID_b: userID,
	}

	userEmail := &UserEmail{
		UserID:       userID,
		Email:        email,
		UpdateNanoTS: updateNanoTS,
	}

	r, createErr := UserEmail_c.CreateOnly(query, userEmail)
	if createErr == nil && r.UpsertedCount > 0 {
		return nil
	}

	// check duplicate
	if createErr != nil {
		return createErr
	}

	gotUserEmail := &UserEmail{}
	err = UserEmail_c.FindOne(query, gotUserEmail, nil)
	if err != nil {
		return err
	}

	if gotUserEmail.IsSet && updateNanoTS-gotUserEmail.UpdateNanoTS < types.EXPIRE_USER_EMAIL_IS_SET_NANO_TS {
		return ErrNoCreate
	}

	if updateNanoTS-gotUserEmail.UpdateNanoTS < types.EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS {
		return ErrNoCreate
	}

	query[USER_EMAIL_UPDATE_NANO_TS_b] = bson.M{
		"$lt": updateNanoTS,
	}

	r, err = UserEmail_c.UpdateOneOnly(query, userEmail)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserEmailIsSet(userID bbs.UUserID, email string, isSet bool, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		USER_EMAIL_USER_ID_b: userID,
		USER_EMAIL_EMAIL_b:   email,
		USER_EMAIL_UPDATE_NANO_TS_b: bson.M{
			"$lt": updateNanoTS,
		},
	}

	toUpdate := bson.M{
		USER_EMAIL_IS_SET_b:         isSet,
		USER_EMAIL_UPDATE_NANO_TS_b: updateNanoTS,
	}

	r, err := UserEmail_c.UpdateOneOnly(query, toUpdate)
	if err != nil {
		return err
	}

	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
