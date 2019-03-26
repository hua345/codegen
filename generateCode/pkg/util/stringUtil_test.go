package util

import (
	"fmt"
	"generateCode/config"
	"regexp"
	"strings"
	"testing"
)

func TestStrFirstToUpper(t *testing.T) {
	if StrFirstToUpper("hello") != "Hello" {
		t.Error(`StrFirstToUpper("hello") != "Hello"`)
	}
	if StrFirstToUpper("hello") == "hello" {
		t.Error(`StrFirstToUpper("hello") == "hello"`)
	}
}
func TestStrFirstToLower(t *testing.T) {
	if StrFirstToLower("Hello") != "hello" {
		t.Error(`StrFirstToUpper("Hello") != "hello"`)
	}
	if StrFirstToLower("Hello") == "Hello" {
		t.Error(`StrFirstToLower("Hello") == "Hello"`)
	}
}
func TestHandleRestfulURL(t *testing.T) {
	urlStr := "///user/add"
	urlStrList := HandleRestfulURL(urlStr)
	if len(urlStrList) != 2 {
		t.Error(`len(HandleRestfulURL("///user/add")) != 2`)
	}
	if urlStrList[0] != "user" {
		t.Error(`HandleRestfulURL("///user/add")[0] != "user"`)
	}
}
func TestHandleRestfulURL2(t *testing.T) {
	urlStr := "///user/{name}"
	urlStrList := HandleRestfulURL(urlStr)
	if len(urlStrList) != 2 {
		t.Error(`len(HandleRestfulURL("///user/add")) != 2`)
	}
	if urlStrList[0] != "user" {
		t.Error(`HandleRestfulURL("///user/add")[0] != "user"`)
	}
	if urlStrList[1] != "{name}" {
		t.Error(`HandleRestfulURL("///user/add")[1] != "{name}"`)
	}
}
func TestAppendURL(t *testing.T) {
	urlStr := AppendURL("api/v1/sdc", "user")
	if urlStr != "/api/v1/sdc/user" {
		t.Error(`AppendURL("api/v1/sdc", "user") != "/api/v1/sdc/user"`)
	}
}
func TestAppendURL2(t *testing.T) {
	urlStr2 := AppendURL("//api/v1/sdc//", "/user/")
	if urlStr2 != "/api/v1/sdc/user" {
		t.Error(`AppendURL("//api/v1/sdc//", "/user/") != "/api/v1/sdc/user"`)
	}
}
func TestAppendURL3(t *testing.T) {
	urlStr2 := AppendURL("//api//v1//sdc//", "///////user/////{name}////")
	if urlStr2 != "/api/v1/sdc/user/{name}" {
		t.Error(`AppendURL("//api//v1//sdc//", "///////user/////{name}////") != "/api/v1/sdc/user/{name}"`)
	}
}
func TestDemoAppend(t *testing.T) {
	nameList := []string{"fang", "fangfang"}
	nameList = DemoAppend(nameList)
	if len(nameList) != 3 {
		t.Error(`len(nameList) != 3`)
	}
	if nameList[2] != "demo" {
		t.Error(`nameList[2] != "demo"`)
	}
}
func TestDemoAppend2(t *testing.T) {
	nameList := []string{"liu", "fang", "hua"}
	aa := nameList[0 : len(nameList)-1]
	bb := nameList[len(nameList)-1]
	fmt.Printf("%q %s", aa, bb)
}
func TestReg(t *testing.T) {
	var testStr string = `public interface UserService {

    /**
     * addUser
     *
     * @param param 入参
     * @return 出参
     */
    AddUserOutputDTO addUser(AddUserInputDTO param);

    /**
     * deleteUser
     *
     * @param param 入参
     * @return 出参
     */
    DeleteUserOutputDTO deleteUser(DeleteUserInputDTO param);
}
`
	contentReg := regexp.MustCompile(`.*}[\r\n|\n]`)
	srcContentSlice := contentReg.Split(testStr, 10)
	resultControllerContent := srcContentSlice[0 : len(srcContentSlice)-1]
	resultControllerContent = append(resultControllerContent, " dd{} ")

	var resultContent string
	for _, value := range resultControllerContent {
		if len(strings.TrimSpace(value)) >= 3 {
			resultContent = resultContent + value
		}
	}
	resultContent = resultContent + "}" + config.RowLimiter
	fmt.Printf("%q", resultContent)
}
