package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
)

func checkClient(clientID string, clientSecret string) (isValid bool, client *schema.Client) {
	client, err := schema.GetClient(clientID)
	if err != nil {
		return false, nil
	}
	if client == nil {
		return false, nil
	}

	if client.ClientSecret != clientSecret {
		return false, nil
	}

	return true, client
}

func getClientInfo(client *schema.Client) string {
	return string(client.ClientType)
}
