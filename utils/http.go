package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mock_http"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

//HttpPost
//
//Params
//  postData: http-post data
//  result: resp-data, requires pointer of pointer to malloc.
//
//Ex:
//    url := backend.LOGIN_R
//    postData := &backend.LoginParams{}
//    result := &backend.LoginResult{}
//    HttpPost(c, url, postData, nil, &result)
func HttpPost(c *gin.Context, url string, postData interface{}, headers map[string]string, result interface{}) (statusCode int, err error) {

	if isTest {
		return mock_http.HttpPost(url, postData, result)
	}

	remoteAddr := c.ClientIP()

	if headers == nil {
		headers = make(map[string]string)
	}

	headers["Content-Type"] = "application/json"
	headers["Host"] = types.HTTP_HOST
	headers["X-Forwarded-For"] = remoteAddr

	authorization := c.GetHeader("Authorization")
	if authorization != "" {
		headers["Authorization"] = authorization
	}

	jsonBytes, err := json.Marshal(postData)
	if err != nil {
		return 500, err
	}

	buf := bytes.NewBuffer(jsonBytes)

	// req
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return 500, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// send http
	resp, err := httpClient.Do(req)
	if err != nil {
		return 500, err
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 501, err
	}

	if statusCode != 200 {
		errResult := &httpErrResult{}
		err = json.Unmarshal(body, errResult)
		if err != nil {
			return statusCode, errors.New(string(body))
		}
		return statusCode, errors.New(errResult.Msg)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return 501, err
	}

	return 200, nil
}
