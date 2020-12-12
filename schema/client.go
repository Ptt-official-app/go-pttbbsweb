package schema

type Client struct {
	ClientID     string `bson:"client_id"`
	ClientSecret string `bson:"client_secret"`
	RemoteAddr   string `bson:"remote_addr"`
	UpdateNanoTS int64  `bson:"update_nano_ts"`
}

type RegisterClientQuery struct {
	ClientID string `bson:"client_id"`
}
