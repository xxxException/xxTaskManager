package controller

type ApiJsonResponse struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func ApiResource(status bool, msg string, data interface{}) *ApiJsonResponse {
	apiResource := &ApiJsonResponse{Status: status, Msg: msg, Data: data}
	return apiResource
}
