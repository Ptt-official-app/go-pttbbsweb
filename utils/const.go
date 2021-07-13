package utils

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 3 * time.Second}
