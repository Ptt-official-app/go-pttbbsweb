package types

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpPost(c *gin.Context, url string, data interface{}, headers map[string]string, result interface{}) (statusCode int, err error) {

	remoteAddr := c.ClientIP()

	if headers == nil {
		headers = make(map[string]string)
	}

	headers["Host"] = HTTP_HOST
	headers["X-Forwarded-For"] = remoteAddr

	authorization := c.GetHeader("Authorization")
	if authorization != "" {
		headers["Authorization"] = authorization
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return 500, err
	}

	buf := bytes.NewBuffer(jsonBytes)

	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		return 500, err
	}

	statusCode = resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 501, err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return 501, err
	}

	return statusCode, nil
}
