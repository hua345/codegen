package models

type MethodMappingInfo struct {
	MethodName string `json:"methodName"`
	UrlPath string `json:"urlPath"`
	HttpMethod string `json:"httpMethod"`
}