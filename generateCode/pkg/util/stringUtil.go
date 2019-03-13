package util

import "strings"

/**
* 首字母大写
 */
func StrFirstToUpper(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}