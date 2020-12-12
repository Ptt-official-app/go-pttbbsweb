package mock_http

import (
	"encoding/json"

	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
)

func HttpPost(url string, data interface{}, result interface{}) (statusCode int, err error) {
	url = url[len(backend.HTTP_PREFIX):]
	switch url {
	case backend.LOGIN_R:
		return parseResult(Login(data.(*backend.LoginParams)), result)
	case backend.REGISTER_R:
		return parseResult(Register(data.(*backend.RegisterParams)), result)
	default:
		return 500, ErrURL
	}
}

func parseResult(backendResult interface{}, httpResult interface{}) (statusCode int, err error) {
	jsonBytes, err := json.Marshal(backendResult)
	if err != nil {
		return 500, err
	}
	err = json.Unmarshal(jsonBytes, httpResult)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
