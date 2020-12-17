package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
)

func Init() (err error) {
	if client != nil {
		return nil
	}
	client, err = db.NewClient(MONGO_PROTOCOL, MONGO_HOST, MONGO_PORT, MONGO_DBNAME)
	if err != nil {
		return err
	}

	Client_c = client.Collection("client")
	User_c = client.Collection("user")
	AccessToken_c = client.Collection("access_token")
	UserReadArticle_c = client.Collection("user_read_article")
	UserReadBoard_c = client.Collection("user_read_board")

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
