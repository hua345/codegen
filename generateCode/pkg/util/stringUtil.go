package util

import "strings"

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
func DemoAppend(strList []string) []string {
	strList = append(strList, "demo")
	return strList
}
