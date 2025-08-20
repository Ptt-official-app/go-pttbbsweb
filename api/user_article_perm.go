package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

// CheckUserArticlePermReadable
//
// Readable
func CheckUserArticlePermReadable(user *UserInfo, boardID bbs.BBoardID, articleID bbs.ArticleID, isCheckBoard bool, c *gin.Context) (err error) {
	if isCheckBoard {
		_, err = CheckUserBoardPermReadable(user, boardID, c)
		if err != nil {
			return err
		}
	}

	articlePermInfo, err := schema.GetArticlePermInfo(boardID, articleID)
	if err != nil {
		return err
	}

	if articlePermInfo == nil {
		return ErrNoArticle
	}

	if articlePermInfo.IsDeleted {
		return ErrNoArticle
	}

	return nil
}

// CheckUserArticlePermEditable
//
// Editable
func CheckUserArticlePermEditable(user *UserInfo, boardID bbs.BBoardID, articleID bbs.ArticleID, isCheckBoard bool, c *gin.Context) (err error) {
	if isCheckBoard {
		_, err = CheckUserBoardPermReadable(user, boardID, c)
		if err != nil {
			return err
		}
	}

	articlePermInfo, err := schema.GetArticlePermInfo(boardID, articleID)
	if err != nil {
		return err
	}
	if articlePermInfo == nil {
		return ErrNoArticle
	}

	return checkUserArticlePermEditableCore(user.UserID, boardID, articleID, articlePermInfo)
}

func checkUserArticlePermEditableCore(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID, articlePermInfo *schema.ArticlePermInfo) (err error) {
	if articlePermInfo.IsDeleted {
		return ErrNoArticle
	}

	if userID == articlePermInfo.Owner {
		return nil
	}

	return ErrInvalidUser
}

//	CheckUserArticlesPermEditable
//
// articles Editable
func CheckUserArticlesPermEditable(user *UserInfo, boardID bbs.BBoardID, articleIDs []bbs.ArticleID, userBoardPerm *UserBoardPermReadable, c *gin.Context) (articlePermMap map[bbs.ArticleID]error, err error) {
	if userBoardPerm == nil {
		userBoardPerm, err = CheckUserBoardPermReadable(user, boardID, c)
		if err != nil {
			return nil, err
		}
	}

	if len(articleIDs) == 0 {
		return make(map[bbs.ArticleID]error), nil
	}

	articlesPermInfo, err := schema.GetArticlesPermInfo(boardID, articleIDs)
	if err != nil {
		return nil, err
	}
	if articlesPermInfo == nil {
		return nil, ErrNoArticle
	}

	return checkUserArticlesPermEditableCore(user.UserID, articlesPermInfo, userBoardPerm), nil
}

func checkUserArticlesPermEditableCore(userID bbs.UUserID, articlesPermInfo []*schema.ArticlePermInfo, userBoardPerm *UserBoardPermReadable) (articlePermMap map[bbs.ArticleID]error) {
	articlePermMap = make(map[bbs.ArticleID]error)
	for _, each := range articlesPermInfo {
		var err error
		if each.IsDeleted {
			err = ErrNoArticle
		} else if userID == each.Owner {
			err = nil
		} else {
			err = ErrInvalidUser
		}

		articlePermMap[each.ArticleID] = err
	}

	return articlePermMap
}

// CheckUserArticlePermDeletable
//
// Deletable
func CheckUserArticlePermDeletable(user *UserInfo, boardID bbs.BBoardID, articleID bbs.ArticleID, c *gin.Context) (err error) {
	userBoardPermReadable, err := CheckUserBoardPermReadable(user, boardID, c)
	if err != nil {
		return err
	}

	articlePermInfo, err := schema.GetArticlePermInfo(boardID, articleID)
	if err != nil {
		return err
	}
	if articlePermInfo == nil {
		return ErrNoArticle
	}

	return checkUserArticlePermDeletableCore(user.UserID, boardID, articleID, articlePermInfo, userBoardPermReadable)
}

func checkUserArticlePermDeletableCore(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID, articlePermInfo *schema.ArticlePermInfo, userBoardPerm *UserBoardPermReadable) (err error) {
	if articlePermInfo.IsDeleted {
		return ErrNoArticle
	}

	if userID == articlePermInfo.Owner {
		return nil
	}

	if userBoardPerm.IsBM {
		return nil
	}

	if userBoardPerm.IsSYSOP {
		return nil
	}

	return ErrInvalidUser
}

// CheckUserArticlePermDeletable
//
// Deletable
func CheckUserArticlesPermDeletable(user *UserInfo, boardID bbs.BBoardID, articleIDs []bbs.ArticleID, userBoardPermReadable *UserBoardPermReadable, c *gin.Context) (articlePermMap map[bbs.ArticleID]error, err error) {
	if userBoardPermReadable == nil {
		userBoardPermReadable, err = CheckUserBoardPermReadable(user, boardID, c)
		if err != nil {
			return nil, err
		}
	}

	if len(articleIDs) == 0 {
		return make(map[bbs.ArticleID]error), nil
	}

	articlesPermInfo, err := schema.GetArticlesPermInfo(boardID, articleIDs)
	if err != nil {
		return nil, err
	}
	if articlesPermInfo == nil {
		return nil, ErrNoArticle
	}

	return checkUserArticlesPermDeletableCore(user.UserID, articlesPermInfo, userBoardPermReadable), nil
}

func checkUserArticlesPermDeletableCore(userID bbs.UUserID, articlesPermInfo []*schema.ArticlePermInfo, userBoardPerm *UserBoardPermReadable) (articlePermMap map[bbs.ArticleID]error) {
	articlePermMap = make(map[bbs.ArticleID]error)
	for _, each := range articlesPermInfo {
		var err error
		if each.IsDeleted {
			err = ErrNoArticle
		} else if userID == each.Owner {
			err = nil
		} else if userBoardPerm.IsBM {
			err = nil
		} else if userBoardPerm.IsSYSOP {
			err = nil
		} else {
			err = ErrInvalidUser
		}
		articlePermMap[each.ArticleID] = err
	}

	return articlePermMap
}

//	CheckUserArticlesPermEditable
//
// articles Editable
func CheckUserArticlesPermEditableDeletable(user *UserInfo, boardID bbs.BBoardID, articleIDs []bbs.ArticleID, userBoardPerm *UserBoardPermReadable, c *gin.Context) (articlePermEditableMap map[bbs.ArticleID]error, articlePermDeletableMap map[bbs.ArticleID]error, err error) {
	if userBoardPerm == nil {
		userBoardPerm, err = CheckUserBoardPermReadable(user, boardID, c)
		if err != nil {
			return nil, nil, err
		}
	}

	if len(articleIDs) == 0 {
		return make(map[bbs.ArticleID]error), make(map[bbs.ArticleID]error), nil
	}

	articlesPermInfo, err := schema.GetArticlesPermInfo(boardID, articleIDs)
	if err != nil {
		return nil, nil, err
	}
	if articlesPermInfo == nil {
		return nil, nil, ErrNoArticle
	}

	return checkUserArticlesPermEditableCore(user.UserID, articlesPermInfo, userBoardPerm), checkUserArticlesPermDeletableCore(user.UserID, articlesPermInfo, userBoardPerm), nil
}
