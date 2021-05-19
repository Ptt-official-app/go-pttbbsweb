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
	case api.LOAD_AUTO_COMPLETE_BOARDS_R:
		return parseResult(LoadAutoCompleteBoards(data.(*api.LoadAutoCompleteBoardsParams)), result)
	case api.LOAD_GENERAL_BOARDS_BY_CLASS_R:
		return parseResult(LoadGeneralBoardsByClass(data.(*api.LoadGeneralBoardsParams)), result)
	case api.LOAD_HOT_BOARDS_R:
		return parseResult(LoadHotBoards(), result)
	case "/board/1_test1/articles":
		return parseResult(LoadGeneralArticles(data.(*api.LoadGeneralArticlesParams)), result)
	case "/board/10_WhoAmI/article/1VtWRel9SYSOP":
		return parseResult(GetArticleDetail(data.(*api.GetArticleParams)), result)
	case "/board/10_WhoAmI/article/1VrooM21SYSOP":
		return parseResult(GetArticleDetail2(data.(*api.GetArticleParams)), result)
	case "/board/10_WhoAmI/article":
		return parseResult(CreateArticle(data.(*api.CreateArticleParams)), result)
	case "/user/SYSOP/information":
		return parseResult(GetUser(), result)
	case "/user/SYSOP/changepasswd":
		return parseResult(ChangePasswd(data.(*api.ChangePasswdParams)), result)
	case "/user/SYSOP/attemptchangeemail":
		return parseResult(AttemptChangeEmail(data.(*api.AttemptChangeEmailParams)), result)
	case "/user/SYSOP/changeemail":
		return parseResult(ChangeEmail(data.(*api.ChangeEmailParams)), result)
	case "/user/SYSOP/attemptsetidemail":
		return parseResult(AttemptSetIDEmail(data.(*api.AttemptSetIDEmailParams)), result)
	case "/user/SYSOP/setidemail":
		return parseResult(SetIDEmail(data.(*api.SetIDEmailParams)), result)
	case "/emailtoken/info":
		return parseResult(GetEmailTokenInfo(data.(*api.GetEmailTokenInfoParams)), result)
	case "/user/SYSOP/favorites":
		return parseResult(GetFavorites(data.(*api.GetFavoritesParams)), result)
	case api.LOAD_BOARDS_BY_BIDS_R:
		return parseResult(LoadBoardsByBids(data.(*api.LoadBoardsByBidsParams)), result)
	case api.CHECK_EXISTS_USER_R:
		return parseResult(CheckExistsUser(data.(*api.CheckExistsUserParams)), result)
	case "/board/10_WhoAmI/isvalid":
		return parseResult(IsBoardValidUser(), result)
	case "/board/1_test1/summary":
		return parseResult(GetBoardSummary(data.(*api.LoadBoardSummaryParams)), result)
	case "/class/2/board":
		return parseResult(CreateBoard(data.(*api.CreateBoardParams)), result)
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
