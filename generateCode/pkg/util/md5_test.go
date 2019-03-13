package util

import (
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	if EncodeMD5("hello") != "5d41402abc4b2a76b9719d911017c592" {
		t.Error(`EncodeMD5("hello") != "5d41402abc4b2a76b9719d911017c592"`)
	}
	if EncodeMD5("world") != "7d793037a0760186574b0282f2f435e7" {
		t.Error(`EncodeMD5("world") != "7d793037a0760186574b0282f2f435e7"`)
	}
}
