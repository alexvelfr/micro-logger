package micrologger

import (
	"context"
	"errors"
	"io"

	logstashclientmicro "github.com/alexvelfr/logstash-client-micro"
)

type logWriter struct{}

var logClient logstashclientmicro.Client

// InitLogger init logstash logger client
func InitLogger(servceName, uri string, useInsecureSSL bool) {
	logClient = logstashclientmicro.NewClient(servceName, uri, useInsecureSSL)
}

// LogError log it
func LogError(reqID, action, file, data string, err error) error {
	if logClient == nil {
		return errors.New("logger not init, please run InitLogger first")
	}
	return logClient.LogError(context.Background(), logstashclientmicro.Message{
		XReqID: reqID,
		Data:   data,
		File:   file,
		Action: action,
		Error:  err,
	})
}

func GetWriter() io.Writer {
	return logWriter{}
}

func (l logWriter) Write(p []byte) (int, error) {
	LogError("", "PANIC", "", string(p), errors.New("PANIC"))
	return len(p), nil
}
