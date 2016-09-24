package framework

import (
	"testing"
)

func TestFrameworkInit(t *testing.T) {
	err, app := New("a", "b", "c", "d")
	if err != nil {
		t.Error("Test init framework failed")
	}

	if app.AppName != "a" {
		t.Error("Test init framework failed: values not match")
	}
}

func TestGenerate(t *testing.T) {
	_, app := New("a", "msbu-test", "../template/service", "../test")

	app.Generate()

	// file_list, _ := walkFileList("../test")
}

func TestWalkFileListWithDot(t *testing.T) {
	file_list, err := walkFileList("./")

	if err != nil {
		t.Error("Cannot walk .")
	}
	if len(file_list) != 2 || file_list[0] != "framework.go" || file_list[1] != "framework_test.go" {
		t.Error(file_list)
	}
}

func TestWalkFileListWithDir(t *testing.T) {
	file_list, err := walkFileList("../cmd")

	if err != nil {
		t.Error("Cannot walk .")
	}
	if len(file_list) != 2 || file_list[0] != "../cmd/go-sf/main.go" || file_list[1] != "../cmd/version/version.go" {
		t.Error(file_list)
	}
}
