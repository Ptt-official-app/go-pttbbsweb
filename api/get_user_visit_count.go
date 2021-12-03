package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/gin-gonic/gin"
)

const GET_USER_VISIT_COUNT_R = "/uservisitcount"

type GetUserVisitCountResult struct {
	Total int64 `json:"total"`
}

func GetUserVisitCountWrapper(c *gin.Context) {
	Query(GetUserVisitCount, nil, c)
}

func GetUserVisitCount(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	// get backend(go-pttbbs || c-pttbbs) user count
	url := pttbbsapi.GET_USER_VISIT_COUNT_R
	var result_b *pttbbsapi.GetUserVisitCountResult
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, 500, err
	}
	// get openbbsmiddleware user count
	currentUserVisitCount := schema.GetUserVisitCount()
	// total user
	result = GetUserVisitCountResult{int64(result_b.Total) + currentUserVisitCount}
	return result, 200, nil
}
