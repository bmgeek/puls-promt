package logger

import (
	"fmt"
	"time"
)

var (
	LogFile       = log.New(outfile, "", 0)
)

func forError(text string, er error) {
	if er != nil {
		LogFile.Println(text, er)
		LogFile.Println(fmt.Sprintf("%s: %s", time.Now().Format("2006.01.02 15:04:05"), er))
	}
}
