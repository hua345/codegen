package util

import (
	"strings"
)

/**
* 首字母大写
 */
func StrFirstToUpper(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

/**
* 首字母小写
 */
func StrFirstToLower(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}
func AppendURL(baseUrl, url string) string {
	resultUrlStrList := HandleRestfulURL(baseUrl)
	urlStrList := HandleRestfulURL(url)
	for _, value := range urlStrList {
		resultUrlStrList = append(resultUrlStrList, value)
	}
	var resultURL string
	for _, value := range resultUrlStrList {
		resultURL = resultURL + "/" + value
	}
	return resultURL
}

/**
 * 获取有效URL路径列表
 */
func HandleRestfulURL(restfulURL string) []string {
	urlStrList := strings.Split(restfulURL, "/")
	var resultURL []string
	for _, value := range urlStrList {
		if len(value) >= 1 {
			resultURL = append(resultURL, value)
		}
	}
	return resultURL
}

/**
 * 测试Append函数
 */
func DemoAppend(strList []string) []string {
	strList = append(strList, "demo")
	return strList
}
