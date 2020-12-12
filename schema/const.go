package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/db"

var (
	client            *db.Client
	Client_c          *db.Collection
	User_c            *db.Collection
	AccessToken_c     *db.Collection
	UserReadArticle_c *db.Collection
	UserReadBoard_c   *db.Collection
)
