package logger

import (
	// "log"
	"os"
)

const (
	Out = 0
	In = 1
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

func (l *Logger) Log(direction int, data []byte) error {
	var err error
	// create an output buffer with a header 
	output := []byte("[!]<")
	if direction == Out {
		output =  []byte("[!]>")
	}

	// Append the data to the header and write it to the file
	output = append(output, data...)
	
	_, err = l.file.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Close() {
	l.file.Close()
}