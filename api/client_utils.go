package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
)

func isValidClient(clientID string, clientSecret string) bool {
	client, err := schema.GetClient(clientID)
	if err != nil {
		return false
	}
	if client == nil {
		return false
	}

	return client.ClientSecret == clientSecret
}
