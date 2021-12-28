package micrologger

import (
	"context"
	"errors"
	"io"
	"log"

	logstashclientmicro "github.com/alexvelfr/logstash-client-micro"
)

var ErrLogNotInit = errors.New("logger not init, please run InitLogger first")
var logClient logstashclientmicro.Client
var consoleMode = false

// InitLogger init logstash logger client
func InitLogger(servceName, uri string, useInsecureSSL bool) {
	logClient = logstashclientmicro.NewClient(servceName, uri, useInsecureSSL)
}

func LogError(reqID, action, file, data string, err error) error {
	return logCommon(reqID, action, file, data, logstashclientmicro.Error, err)
}

func LogInfo(reqID, action, file, data string) error {
	return logCommon(reqID, action, file, data, logstashclientmicro.Info, nil)
}

func EnabldeConsoleMode() {
	consoleMode = true
}

func DisableConsoleMode() {
	consoleMode = false
}

func LogDebug(reqID, action, file, data string) error {
	return logCommon(reqID, action, file, data, logstashclientmicro.Debug, nil)
}

func LogWarning(reqID, action, file, data string) error {
	return logCommon(reqID, action, file, data, logstashclientmicro.Warning, nil)
}

func LogErrorStrict(err error) {
	logCommon("", "Stric error", "", "", logstashclientmicro.Error, err)
}

func logCommon(reqID, action, file, data string, t logstashclientmicro.LogType, err error) error {
	if logClient == nil {
		return ErrLogNotInit
	}
	if consoleMode {
		log.Printf("XReqID=%s;\nData=%s;\nFile=%s;\nAction=%s;\nErr=%s;\nType=%s\n",
			reqID,
			data,
			file,
			action,
			err,
			t,
		)
		return nil
	}
	return logClient.LogError(context.Background(), logstashclientmicro.Message{
		XReqID: reqID,
		Data:   data,
		File:   file,
		Action: action,
		Error:  err,
		Type:   t,
	})
}

func GetWriter() io.Writer {
	return logWriter{}
}

type logWriter struct{}

func (l logWriter) Write(p []byte) (int, error) {
	LogError("", "PANIC", "", string(p), errors.New("PANIC"))
	return len(p), nil
}
