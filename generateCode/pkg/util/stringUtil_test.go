package util

import "testing"

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
