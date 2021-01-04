package mock_http

import (
	"reflect"

	"github.com/Ptt-official-app/go-pttbbs/api"
	log "github.com/sirupsen/logrus"
)

func HttpPost(url string, data interface{}, result interface{}) (statusCode int, err error) {
	log.Infof("HttpPost: url: %v", url)
	switch url {
	case api.LOGIN_R:
		return parseResult(Login(data.(*api.LoginParams)), result)
	case api.REGISTER_R:
		return parseResult(Register(data.(*api.RegisterParams)), result)
	case api.LOAD_GENERAL_BOARDS_R:
		return parseResult(LoadGeneralBoards(data.(*api.LoadGeneralBoardsParams)), result)
	case "/boards/1_test1/articles":
		return parseResult(LoadGeneralArticles(data.(*api.LoadGeneralArticlesParams)), result)
	case "/boards/10_WhoAmI/articles/19bWBI4ZSYSOP":
		return parseResult(GetArticleDetail(data.(*api.GetArticleParams)), result)
	case "/boards/10_WhoAmI/articles/1VrooM21SYSOP":
		return parseResult(GetArticleDetail2(data.(*api.GetArticleParams)), result)
	case "/users/SYSOP/information":
		return parseResult(GetUser(), result)
	default:
		return 500, ErrURL
	}
}

func parseResult(backendResult interface{}, httpResult interface{}) (statusCode int, err error) {

	convert(backendResult, httpResult)

	return 200, nil
}

func convert(dataBackend interface{}, dataResult interface{}) {
	valueBackend := reflect.ValueOf(dataResult)
	valuePttbbs := reflect.ValueOf(dataBackend)
	valueBackend.Elem().Set(valuePttbbs)
}
