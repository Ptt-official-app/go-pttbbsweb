package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
)

// CheckUserArticlePermReadable
//
// Readable
func CheckUserArticlePermReadable(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID, isCheckBoard bool) (err error) {
	if isCheckBoard {
		_, err = CheckUserBoardPermReadable(userID, boardID)
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
func CheckUserArticlePermEditable(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID, isCheckBoard bool) (err error) {
	if isCheckBoard {
		_, err = CheckUserBoardPermReadable(userID, boardID)
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

	return checkUserArticlePermEditableCore(userID, boardID, articleID, articlePermInfo)
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
func CheckUserArticlesPermEditable(userID bbs.UUserID, boardID bbs.BBoardID, articleIDs []bbs.ArticleID, userBoardPerm *UserBoardPermReadable) (articlePermMap map[bbs.ArticleID]error, err error) {
	if userBoardPerm == nil {
		userBoardPerm, err = CheckUserBoardPermReadable(userID, boardID)
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

	return checkUserArticlesPermEditableCore(userID, articlesPermInfo, userBoardPerm), nil
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
func CheckUserArticlePermDeletable(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID) (err error) {
	userBoardPermReadable, err := CheckUserBoardPermReadable(userID, boardID)
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

	return checkUserArticlePermDeletableCore(userID, boardID, articleID, articlePermInfo, userBoardPermReadable)
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
func CheckUserArticlesPermDeletable(userID bbs.UUserID, boardID bbs.BBoardID, articleIDs []bbs.ArticleID, userBoardPermReadable *UserBoardPermReadable) (articlePermMap map[bbs.ArticleID]error, err error) {
	if userBoardPermReadable == nil {
		userBoardPermReadable, err = CheckUserBoardPermReadable(userID, boardID)
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

	return checkUserArticlesPermDeletableCore(userID, articlesPermInfo, userBoardPermReadable), nil
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
func CheckUserArticlesPermEditableDeletable(userID bbs.UUserID, boardID bbs.BBoardID, articleIDs []bbs.ArticleID, userBoardPerm *UserBoardPermReadable) (articlePermEditableMap map[bbs.ArticleID]error, articlePermDeletableMap map[bbs.ArticleID]error, err error) {
	if userBoardPerm == nil {
		userBoardPerm, err = CheckUserBoardPermReadable(userID, boardID)
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

	return checkUserArticlesPermEditableCore(userID, articlesPermInfo, userBoardPerm), checkUserArticlesPermDeletableCore(userID, articlesPermInfo, userBoardPerm), nil
}
