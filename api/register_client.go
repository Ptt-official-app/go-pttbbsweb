package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/gin-gonic/gin"
)

const REGISTER_CLIENT_R = "/client/register"

type RegisterClientParams struct {
	ClientID string `json:"client_id"`
}

type RegisterClientResult struct {
	ClientSecret string `json:"client_secret"`
	Success      bool
}

func NewRegisterClientParams() *RegisterClientParams {
	return &RegisterClientParams{}
}

func RegisterClient(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	RegisterClientParams, ok := params.(*RegisterClientParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	clientSecret := genClientSecret()

	//update db
	nowNanoTS := utils.GetNowNanoTS()
	query := &schema.RegisterClientQuery{ClientID: RegisterClientParams.ClientID}
	update := &schema.Client{
		ClientID:     RegisterClientParams.ClientID,
		ClientSecret: clientSecret,
		RemoteAddr:   remoteAddr,
		UpdateNanoTS: nowNanoTS,
	}

	_, err = schema.Client_c.Update(query, update)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = &RegisterClientResult{
		ClientSecret: clientSecret,
		Success:      true,
	}
	return result, 200, nil
}

func genClientSecret() string {
	if types.SERVICE_MODE == types.DEV {
		return "test_client_secret"
	}

	return utils.GenRandomString()
}
