package test

import (
	"testing"

	"github.com/ZEL-30/gin-web-app/util"
)

func TestEncodeMD5(t *testing.T) {
	t.Log("TestEncodeMD5()")

	// TestEncodeMD5()
	if util.EncodeMD5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("EncodeMD5() failed")
	}

}
