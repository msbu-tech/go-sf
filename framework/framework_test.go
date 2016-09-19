package framework

import (
	// "fmt"
	"testing"
)

func TestFrameworkdInit(t *testing.T) {
	err, app := New("a", "b", "c", "d")
	if err != nil {
		t.Error("Test init framework failed")
	}

	if app.Appname != "a" {
		t.Error("Test init framework failed: values not match")
	}
}

func TestGenerate(t *testing.T) {
	_, app := New("a", "msbu-test", "template/service", "./test/")

	app.Generate()
}
