package log

import (
	"testing"
)

func TestInfo2(t *testing.T) {
	pth := "./test.log"
	err := Init(pth)
	println(err)
	Info("TestInfo2 called")
	Info("TestInfo2 called, [%d]", 2)
}
