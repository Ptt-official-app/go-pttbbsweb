package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
)

const REGISTER_CLIENT_R = "/client/register"

type RegisterClientParams struct {
	ClientID   string           `json:"client_id"`
	ClientType types.ClientType `json:"client_type"`
}

type RegisterClientResult struct {
	ClientSecret string `json:"client_secret"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewRegisterClientParams() *RegisterClientParams {
	return &RegisterClientParams{}
}

func RegisterClientWrapper(c *gin.Context) {
	params := NewRegisterClientParams()
	LoginRequiredJSON(RegisterClient, params, c)
}

func RegisterClient(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	registerClientParams, ok := params.(*RegisterClientParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if !isRegisterClientValidRemoteAddr(remoteAddr) {
		return nil, 401, ErrInvalidRemoteAddr
	}

	client, err := deserializeClientAndUpdateDB(registerClientParams, remoteAddr)
	if err != nil {
		return nil, 500, err
	}

	// result
	result = NewRegisterClientResult(client, userID)
	return result, 200, nil
}

func deserializeClientAndUpdateDB(registerClientParams *RegisterClientParams, remoteAddr string) (client *schema.Client, err error) {
	client = schema.NewClient(registerClientParams.ClientID, registerClientParams.ClientType, remoteAddr)

	err = schema.UpdateClient(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewRegisterClientResult(client *schema.Client, userID bbs.UUserID) *RegisterClientResult {
	return &RegisterClientResult{
		ClientSecret: client.ClientSecret,
		TokenUser:    userID,
	}
}

func isRegisterClientValidRemoteAddr(remoteAddr string) bool {
	if IsTest {
		return true
	}

	if types.SERVICE_MODE == types.DEV {
		return true
	}

	return remoteAddr == "127.0.0.1" || remoteAddr == "localhost"
}
