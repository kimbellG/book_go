package filelog

import (
	"fmt"
	"log"
	"os"
)

var logFile os.File

func init() {

	var err error
	logFile, err = os.Create(".log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Don't open file\n")
		exit(2)
	}

	log.SetOutput(logFile)
}
