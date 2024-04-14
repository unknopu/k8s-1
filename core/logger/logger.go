package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type OutputSplitter struct{}

func (splitter *OutputSplitter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("level=error")) {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}

func LoggerFile(serviceName, path string) (*os.File, error) {
	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s.log", path, serviceName),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666,
	)
	if err != nil {
		return nil, err
	}

	_ = io.MultiWriter(os.Stdout, file)
	return file, nil
}
