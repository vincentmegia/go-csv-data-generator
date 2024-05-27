package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var format = "[%s] - %s"

type Logger struct {
	LogFile *os.File
}

func (l *Logger) Initialize() {
	log.SetFlags(0)
	log.SetOutput(l)
}

func (l *Logger) Log(v ...any) {
	log.Print(v...)
	bytes := fmt.Append([]byte{}, v...)
	l.LogFile.WriteString(fmt.Sprintf(format, time.Now().UTC().Format(time.RFC3339), string(bytes)))
}

func (l *Logger) Write(b []byte) (int, error) {
	timeString := time.Now().UTC().Format(time.RFC3339)
	output := fmt.Sprintf(format, timeString, string(b))
	return fmt.Print(output)
}
