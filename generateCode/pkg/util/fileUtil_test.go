package util

import (
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	var path string = "README.md"
	var dstPath string = "Fang.md"
	var content string = "FangFang"
	WriteFileWithIoUtil(path, content)
	_, err := CopyFile(dstPath, path)
	if err != nil {
		panic(err)
	}
	srcContent := ReadFileWithIoUtil(path)
	dstContent := ReadFileWithIoUtil(dstPath)
	if srcContent != content {
		t.Error(`srcContent != "FangFang"` )
	}
	if dstContent != content {
		t.Error(`dstContent != "FangFang"` )
	}
	os.Remove(path)
	os.Remove(dstPath)
}