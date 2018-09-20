package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

var Logger *log.Logger = log.New(os.Stdout, "[logging] ", log.LstdFlags|log.Lshortfile)

func GetLogger(w io.Writer, codename string) *log.Logger {
	if w == nil {
		w = os.Stderr
	}
	if len(codename) == 0 {
		codename = "logging"
	}
	prefix := fmt.Sprintf("[%s] ", codename)
	return log.New(w, prefix, log.LstdFlags|log.Lshortfile)
}

type Interface interface {
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

type logger struct {
	prefix string
}

func New(codename string) Interface {
	if len(codename) == 0 {
		codename = "logging"
	}
	prefix := fmt.Sprintf("[%s] ", codename)
	return &logger{prefix}
}

func (m *logger) Infof(format string, a ...interface{}) {
	l := GetLogger(os.Stdout, m.prefix)
	l.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, a...))
}

func (m *logger) Warnf(format string, a ...interface{}) {
	l := GetLogger(os.Stdout, m.prefix)
	l.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, a...))
}

func (m *logger) Errorf(format string, a ...interface{}) {
	l := GetLogger(os.Stderr, m.prefix)
	l.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf(format, a...))
}
