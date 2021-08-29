package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/cron"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func withPrefix(path string) string {
	return types.API_PREFIX + path
}

func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
			f()
		}
	}()

	return ctx
}

func initGin() (*gin.Engine, error) {
	router := gin.Default()

	// options
	router.OPTIONS("/*path", api.OptionsWrapper)

	// index
	router.GET(withPrefix(api.INDEX_R), api.IndexWrapper)

	router.GET(withPrefix(api.GET_VERSION_R), api.GetVersionWrapper)

	// register/login
	router.POST(withPrefix(api.REGISTER_CLIENT_R), api.RegisterClientWrapper)
	router.POST(withPrefix(api.REGISTER_USER_R), api.RegisterUserWrapper)
	router.POST(withPrefix(api.LOGIN_R), api.LoginWrapper)
	router.POST(withPrefix(api.ATTEMPT_REGISTER_USER_R), api.AttemptRegisterUserWrapper)
	router.POST(withPrefix(api.CHECK_EXISTS_USER_R), api.CheckExistsUserWrapper)

	// board
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_R), api.LoadGeneralBoardsWrapper)
	router.GET(withPrefix(api.GET_BOARD_DETAIL_R), api.GetBoardDetailWrapper)
	router.GET(withPrefix(api.GET_BOARD_SUMMARY_R), api.GetBoardSummaryWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_BOARDS_R), api.LoadPopularBoardsWrapper)
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_BY_CLASS_R), api.LoadGeneralBoardsByClassWrapper)
	router.GET(withPrefix(api.LOAD_AUTO_COMPLETE_BOARDS_R), api.LoadAutoCompleteBoardsWrapper)
	router.POST(withPrefix(api.CREATE_BOARD_R), api.CreateBoardWrapper)
	router.GET(withPrefix(api.LOAD_CLASS_BOARDS_R), api.LoadClassBoardsWrapper)

	// article
	router.GET(withPrefix(api.LOAD_GENERAL_ARTICLES_R), api.LoadGeneralArticlesWrapper)
	router.GET(withPrefix(api.LOAD_BOTTOM_ARTICLES_R), api.LoadBottomArticlesWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_R), api.GetArticleDetailWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_ARTICLES_R), api.LoadPopularArticlesWrapper)
	router.POST(withPrefix(api.CREATE_ARTICLE_R), api.CreateArticleWrapper)
	router.GET(withPrefix(api.CROSS_POST_R), api.CrossPostWrapper)

	// user
	router.GET(withPrefix(api.GET_USER_INFO_R), api.GetUserInfoWrapper)
	router.GET(withPrefix(api.LOAD_FAVORITE_BOARDS_R), api.LoadFavoriteBoardsWrapper)
	router.GET(withPrefix(api.LOAD_USER_ARTICLES_R), api.LoadUserArticlesWrapper)
	router.POST(withPrefix(api.CHANGE_PASSWD_R), api.ChangePasswdWrapper)
	router.POST(withPrefix(api.ATTEMPT_CHANGE_EMAIL_R), api.AttemptChangeEmailWrapper)
	router.POST(withPrefix(api.CHANGE_EMAIL_R), api.ChangeEmailWrapper)
	router.POST(withPrefix(api.ATTEMPT_SET_ID_EMAIL_R), api.AttemptSetIDEmailWrapper)
	router.POST(withPrefix(api.SET_ID_EMAIL_R), api.SetIDEmailWrapper)
	router.GET(withPrefix(api.GET_USER_ID_R), api.GetUserIDWrapper)

	// comments
	router.GET(withPrefix(api.LOAD_ARTICLE_COMMENTS_R), api.LoadArticleCommentsWrapper)
	router.GET(withPrefix(api.LOAD_USER_COMMENTS_R), api.LoadUserCommentsWrapper)
	router.POST(withPrefix(api.CREATE_COMMENT_R), api.CreateCommentWrapper)

	// ranks
	router.POST(withPrefix(api.CREATE_RANK_R), api.CreateRankWrapper)

	// html
	router.GET(api.ROOT_HTML_R, api.IndexHTMLWrapper)
	router.GET(api.INDEX_HTML_R, api.IndexHTMLWrapper)

	router.GET(api.REGISTER_HTML_R, api.RegisterHTMLWrapper)
	router.GET(api.LOGIN_HTML_R, api.LoginHTMLWrapper)

	router.GET(api.USER_HTML_R, api.UserHTMLWrapper)
	router.GET(api.USER_CHANGE_PASSWD_HTML_R, api.UserChangePasswdHTMLWrapper)
	router.GET(api.USER_ATTEMPT_CHANGE_EMAIL_HTML_R, api.UserAttemptChangeEmailHTMLWrapper)
	router.GET(api.USER_CHANGE_EMAIL_HTML_R, api.UserChangeEmailHTMLWrapper)
	router.GET(api.USER_ATTEMPT_SET_ID_EMAIL_HTML_R, api.UserAttemptSetIDEmailHTMLWrapper)
	router.GET(api.USER_SET_ID_EMAIL_HTML_R, api.UserSetIDEmailHTMLWrapper)

	router.GET(api.BOARDS_FAVORITES_HTML_R, api.AllHTMLWrapper)
	router.GET(api.BOARDS_POPULAR_HTML_R, api.AllHTMLWrapper)
	router.GET(api.BOARDS_HTML_R, api.AllHTMLWrapper)

	router.GET(api.ARTICLES_HTML_R, api.AllHTMLWrapper)
	router.GET(api.ARTICLE_HTML_R, api.AllHTMLWrapper)
	router.GET(api.CREATE_ARTICLE_HTML_R, api.AllHTMLWrapper)

	router.Static("/static", filepath.Join(types.STATIC_DIR, "static"))

	staticFiles := []string{
		"asset-manifest.json",
		"favicon.ico",
		"logo192.png",
		"logo512.png",
		"manifest.json",
		"robots.txt",
	}

	for _, each := range staticFiles {
		router.StaticFile("/"+each, filepath.Join(types.STATIC_DIR, each))
	}

	return router, nil
}

func main() {
	finished := make(chan struct{})

	if err := queue.Start(); err != nil {
		log.Fatal(err)
	}

	// ctrl + c
	ctx := withContextFunc(
		context.Background(),
		func() {
			// handle graceful shutdown
			log.Info("interrupt received, terminating process")
			queue.Close()
			close(finished)
		},
	)

	err := initMain()
	if err != nil {
		log.Fatalf("unable to initMain: e: %v", err)
	}

	router, err := initGin()
	if err != nil {
		log.Fatal(err)
	}

	// retry load general boards
	go cron.RetryLoadGeneralBoards()

	// retry load full class boards
	go cron.RetryLoadFullClassBoards()

	s := &http.Server{
		Addr:    types.HTTP_HOST,
		Handler: router,
	}

	var g errgroup.Group

	// graceful shutdown
	g.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.Shutdown(ctx)
	})

	// start the server
	g.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// wait all job are complete.
	g.Go(func() error {
		<-finished
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
