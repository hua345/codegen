package stringUtil

import (
	"fmt"
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

/**
* 驼峰格式
 */
func ToCamelCase(str string) string {
	separators := []string{".", "-", "_"}
	in := []string{str}
	out := make([]string, 0)
	for _, sep := range separators {
		for _, inStr := range in {
			parts := strings.Split(inStr, sep)
			out = append(out, parts...)
		}
		in = out
		out = make([]string, 0)
	}
	words := in
	for i := range words {
		words[i] = StrFirstToUpper(words[i])
	}
	return strings.Join(words, "")
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
	// 处理windows Git命令行,URL输入/开头的情况
	const WindowsGit = "Git"
	if strings.Contains(restfulURL, WindowsGit) {
		urlStrList := strings.Split(restfulURL, WindowsGit)
		if len(urlStrList) != 2 || len(urlStrList[1]) <= 1 {
			fmt.Println("Windows Git命令行下，URL最好不要用`/`开头")
		}
		restfulURL = urlStrList[1]
	}
	urlStrList := strings.Split(restfulURL, "/")
	var resultURL []string
	for _, value := range urlStrList {
		if len(value) >= 1 {
			resultURL = append(resultURL, value)
		}
	}
	return resultURL
}
