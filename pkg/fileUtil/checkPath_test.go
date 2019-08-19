package fileUtil

import "testing"

func TestPathExists(t *testing.T) {
	exist, err := PathExists("helloTest")
	if err != nil {
		panic(err)
	}
	if exist {
		t.Error(`PathExists("helloTest") == true`)
	}
}
func TestPathMkDir(t *testing.T) {
	var path string = "helloTest"
	exist, err := PathExists(path)
	if err != nil {
		panic(err)
	}
	if !exist {
		PathMkDir(path)
		exist2, err := PathExists(path)
		if err != nil {
			panic(err)
		}
		if !exist2 {
			t.Error(`PathMkDir("helloTest") PathExists != true`)
		}
		PathRemove(path)
	}
}
func TestPathMkdirAll(t *testing.T) {
	var path string = "hello/fang/fang"
	exist, err := PathExists(path)
	if err != nil {
		panic(err)
	}
	if !exist {
		PathMkdirAll(path)
		exist2, err := PathExists(path)
		if err != nil {
			panic(err)
		}
		if !exist2 {
			t.Error(`PathMkdirAll("hello/fang/fang") PathExists != true`)
		}
		PathRemoveAll("hello")
	}
}
