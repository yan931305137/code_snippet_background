package types

type LogInfo struct {
	Id           int         `json:"id"`
	ClientIP     string      `json:"ClientIP"`
	EdTime       int         `json:"EdTime"`
	Method       string      `json:"Method"`
	RequestURI   string      `json:"RequestURI"`
	ResponseCode int         `json:"ResponseCode"`
	ResponseData interface{} `json:"ResponseData"`
	ResponseMsg  string      `json:"ResponseMsg"`
	RoleName     string      `json:"RoleName"`
	StartTime    string      `json:"StartTime"`
	Level        string      `json:"level"`
	Msg          string      `json:"msg"`
	Time         string      `json:"time"`
}
