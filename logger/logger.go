package logger

import (
	"os"
	"log"
)

type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// DefaultLogger returns a default implementation for a Logger interface
func DefaultLogger() *Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}