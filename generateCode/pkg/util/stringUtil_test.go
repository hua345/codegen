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
