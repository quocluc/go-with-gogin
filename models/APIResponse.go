package models

type APIResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}

func APIResponseOk(data interface{}) APIResponse {
	return APIResponse{Success: true, Data: data, StatusCode: 200}
}
func APIResponseFail(data interface{}) APIResponse {
	return APIResponse{Success: false, Data: data, StatusCode: 500}
}
