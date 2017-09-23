package simplelog

import (
	"fmt"
	"log"
	"os"
)

var infoLog = log.New(os.Stdout, "[info] ", log.LstdFlags)
var warnLog = log.New(os.Stdout, "[warn] ", log.LstdFlags)
var errorLog = log.New(os.Stderr, "[error] ", log.LstdFlags)
var verboseLog = log.New(os.Stdout, "[verbose] ", log.LstdFlags)

func Info(format string, args ...interface{}) {
	infoLog.Println(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	warnLog.Println(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	errorLog.Println(fmt.Sprintf(format, args...))
}

func Verbose(format string, args ...interface{}) {
	verboseLog.Println(fmt.Sprintf(format, args...))
}
