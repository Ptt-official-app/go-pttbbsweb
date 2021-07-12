package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/gin-gonic/gin"
)

const GET_VERSION_R = "/version"

type GetVersionResult struct {
	Version    string `json:"version"`
	GitVersion string `json:"commit"`

	PttVersion    string `json:"pttversion"`
	PttGitVersion string `json:"pttcommit"`
}

func GetVersionWrapper(c *gin.Context) {
	Query(GetVersion, nil, c)
}

func GetVersion(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	url := pttbbsapi.GET_VERSION_R

	var result_b *pttbbsapi.GetVersionResult

	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	return &GetVersionResult{
		Version:       types.VERSION,
		GitVersion:    types.GIT_VERSION,
		PttVersion:    result_b.Version,
		PttGitVersion: result_b.GitVersion,
	}, 200, nil
}
