package schema

type AccessToken struct {
	AccessToken  string `bson:"access_token"`
	UserID       string `bson:"user_id"`
	UpdateNanoTS int64  `bson:"update_nano_ts"`
}
