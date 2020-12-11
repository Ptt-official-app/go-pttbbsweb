package api

type IndexParams struct {
	In int `form:"in,omitempty"`
}

type IndexResult struct {
	Data interface{}
}

func Index(remoteAddr string, params interface{}) (interface{}, error) {
	result := &IndexResult{Data: params}
	return result, nil
}
