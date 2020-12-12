package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
)

func isValidClient(clientID string, clientSecret string) bool {

	query := &schema.RegisterClientQuery{
		ClientID: clientID,
	}

	ret := &schema.Client{}
	err := schema.Client_c.FindOne(query, ret, nil)
	if err != nil {
		return false
	}

	return ret.ClientSecret == clientSecret
}
