package api

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestJSWrapper(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		args     args
		path     string
		header   map[string]string
		expected string
	}{
		// TODO: Add test cases.
		{
			path:     "/static/js/temp.js",
			expected: "function a() {\n  return \"\";\n}\n",
		},
		{
			path:     "/static/js/temp.js",
			header:   map[string]string{"Origin": "localhost"},
			expected: "function a() {\n  return \"\";\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c, r := testSetRequest(
				tt.path,
				"/static/js/:path",
				nil,
				"",
				"",
				tt.header,
				"POST",
				JSWrapper,
			)

			r.ServeHTTP(w, c.Request)

			if w.Code != http.StatusOK {
				t.Errorf("code: %v", w.Code)
			}

			result := string(w.Body.Bytes())

			re, _ := regexp.Compile("return.*?;")
			result = re.ReplaceAllString(result, "return \"\";")

			if result != tt.expected {
				t.Errorf("JSWrapper: %v expected: %v", result, tt.expected)
				for idx := 0; idx < len(result); idx++ {
					if idx >= len(tt.expected) {
						t.Errorf("(%v/%v) js: %v", idx, len(result), result[idx])
						continue
					}
					if result[idx] != tt.expected[idx] {
						t.Errorf("(%v/%v) js: %v expected: %v", idx, len(result), result[idx], tt.expected[idx])
					}
				}
			}

		})
	}
}
