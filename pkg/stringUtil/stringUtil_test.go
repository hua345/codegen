package stringUtil

import (
	"fmt"
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

func TestToCamelCase(t *testing.T){
	t.Log(ToCamelCase("love.hello_fang"))
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
func TestHandleRestfulURL3(t *testing.T) {
	urlStr := "/user"
	urlStrList := HandleRestfulURL(urlStr)
	fmt.Println(urlStrList)
	if len(urlStrList) != 1 {
		t.Error(`len(HandleRestfulURL("/user")) != 2`)
	}
	if urlStrList[0] != "user" {
		t.Error(`HandleRestfulURL("/user")[0] != "user"`)
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
