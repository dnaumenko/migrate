package logger

import (
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	l := DefaultLogger()
	l.Printf("hello")
}