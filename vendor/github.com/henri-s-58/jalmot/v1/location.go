package jalmot

import (
	"fmt"
	"runtime"
)

type Location interface {
	FuncName() string
	FileName() string
	FileLine() int
	String() string
}

func TakeLocation(err error) Location {
	if jalmot, ok := err.(Jalmot); ok {
		return jalmot.From()
	}
	return nil
}

func newLocation(dep int) Location {
	pc, file, line, _ := runtime.Caller(dep)
	fn := runtime.FuncForPC(pc)
	return &locaiton{
		funcName: fn.Name(),
		fileName: file,
		fileLine: line,
	}
}

type locaiton struct {
	funcName string
	fileName string
	fileLine int
}

func (l *locaiton) FuncName() string {
	if l == nil {
		return ""
	}
	return l.funcName
}

func (l *locaiton) FileName() string {
	if l == nil {
		return ""
	}
	return l.fileName
}

func (l *locaiton) FileLine() int {
	if l == nil {
		return -1
	}
	return l.fileLine
}

func (l *locaiton) String() string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf("%s:%d#%s()", l.fileName, l.fileLine, l.funcName)
}
