package mock_http

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/sirupsen/logrus"
)

func HttpPost(url string, data interface{}, result interface{}) (statusCode int, err error) {
	url = url[len(backend.HTTP_PREFIX):]
	logrus.Infof("HttpPost: url: %v", url)
	switch url {
	case backend.LOGIN_R:
		return parseResult(Login(data.(*backend.LoginParams)), result)
	case backend.REGISTER_R:
		return parseResult(Register(data.(*backend.RegisterParams)), result)
	case backend.LOAD_GENERAL_BOARDS_R:
		return parseResult(LoadGeneralBoards(data.(*backend.LoadGeneralBoardsParams)), result)
	case "/boards/1_test1/articles":
		return parseResult(LoadGeneralArticles(data.(*backend.LoadGeneralArticlesParams)), result)
	default:
		return 500, ErrURL
	}
}

func parseResult(backendResult interface{}, httpResult interface{}) (statusCode int, err error) {

	backend.Convert(backendResult, httpResult)

	return 200, nil
}
