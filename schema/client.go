package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client_c *db.Collection

type Client struct {
	// 可信任的 app-client

	ClientID     string           `bson:"client_id"`
	ClientSecret string           `bson:"client_secret"`
	ClientType   types.ClientType `bson:"client_type"`
	RemoteAddr   string           `bson:"ip"`
	UpdateNanoTS types.NanoTS     `bson:"update_nano_ts"`
}

var EMPTY_CLIENT = &Client{}

var (
	CLIENT_CLIENT_ID_b      = getBSONName(EMPTY_CLIENT, "ClientID")
	CLIENT_CLIENT_SECRET_b  = getBSONName(EMPTY_CLIENT, "ClientSecret")
	CLIENT_REMOTE_ADDR_b    = getBSONName(EMPTY_CLIENT, "RemoteAddr")
	CLIENT_UPDATE_NANO_TS_b = getBSONName(EMPTY_CLIENT, "UpdateNanoTS")
)

func NewClient(clientID string, clientType types.ClientType, remoteAddr string) *Client {
	clientSecret := genClientSecret()
	nowNanoTS := types.NowNanoTS()

	return &Client{
		ClientID:     clientID,
		ClientType:   clientType,
		ClientSecret: clientSecret,
		RemoteAddr:   remoteAddr,
		UpdateNanoTS: nowNanoTS,
	}
}

func UpdateClient(c *Client) (err error) {
	query := bson.M{
		CLIENT_CLIENT_ID_b: c.ClientID,
	}

	r, err := Client_c.CreateOnly(query, c)
	if err != nil {
		return err
	}
	if r.UpsertedCount > 0 {
		return nil
	}

	query[CLIENT_UPDATE_NANO_TS_b] = bson.M{
		"$lt": c.UpdateNanoTS,
	}
	r, err = Client_c.UpdateOneOnly(query, c)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}

func GetClient(clientID string) (ret *Client, err error) {
	query := bson.M{
		CLIENT_CLIENT_ID_b: clientID,
	}

	ret = &Client{}
	err = Client_c.FindOne(query, ret, nil)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func genClientSecret() string {
	if types.SERVICE_MODE == types.DEV {
		return "test_client_secret"
	}

	return utils.GenRandomString()
}
