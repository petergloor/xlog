package logger

import (
	// "log"
	"os"
)

type Logger struct {
	file *os.File
}

func (l *Logger) Open(filename string) error {
	var err error
	l.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	// log.SetOutput(l.file)
	return nil
}

func (l *Logger) Log(msg string, data []byte) error {
	var err error
	_, err = l.file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Close() {
	l.file.Close()
}