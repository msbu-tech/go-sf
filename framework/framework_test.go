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

	if app.Date != "2016-09-19" {
		t.Error("Test init date failed", app.Date)
	}
}