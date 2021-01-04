package api

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIndexHtmlWrapper(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		args     args
		header   map[string]string
		expected string
	}{
		// TODO: Add test cases.
		{
			expected: "<document>\n  <body>\n    Hello World\n  </body>\n</document>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c, r := testSetRequest(
				"/",
				"/",
				nil,
				"",
				"",
				tt.header,
				"GET",
				IndexHtmlWrapper,
			)

			r.ServeHTTP(w, c.Request)

			if w.Code != http.StatusOK {
				t.Errorf("code: %v", w.Code)
			}

			result := string(w.Body.Bytes())

			re, _ := regexp.Compile("return.*?;")
			result = re.ReplaceAllString(result, "return \"\";")

			if result != tt.expected {
				t.Errorf("IndexHtmlWrapper: %v expected: %v", result, tt.expected)
				for idx := 0; idx < len(result); idx++ {
					if idx >= len(tt.expected) {
						t.Errorf("(%v/%v) html: %v", idx, len(result), result[idx])
						continue
					}
					if result[idx] != tt.expected[idx] {
						t.Errorf("(%v/%v) html: %v expected: %v", idx, len(result), result[idx], tt.expected[idx])
					}
				}
			}

		})
	}
}
