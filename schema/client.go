package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type Client struct {
	//可信任的 app-client

	ClientID     string       `bson:"client_id"`
	ClientSecret string       `bson:"client_secret"`
	RemoteAddr   string       `bson:"remote_addr"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

type RegisterClientQuery struct {
	ClientID string `bson:"client_id"`
}
