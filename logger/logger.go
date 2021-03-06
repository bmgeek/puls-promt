package logger

import (
	"fmt"
	"time"
	"log"
	"os"
)

var (
	outfile, _    = os.OpenFile(fmt.Sprintf(os.Getenv("XPIE_LOGS_FILE")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile       = log.New(outfile, "", 0)
)

func ForError(text string, err error) {
	if err != nil {
		LogFile.Println(fmt.Sprintf("%s ERROR %s\n%s", time.Now().Format("2006.01.02 15:04:05"), err, text))
	}
}

func ForInfo(text string) {
	LogFile.Println(fmt.Sprintf("%s INFO %s", time.Now().Format("2006.01.02 15:04:05"), text))
}
