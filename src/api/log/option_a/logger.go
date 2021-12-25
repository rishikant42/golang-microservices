package log_option_a

import (
	"fmt"
	"os"
	"strings"

	"github.com/rishikant42/golang-microservices/src/api/config"
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)

	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}
	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
		// Log.Formatter = &logrus.JSONFormatter{}
	}
}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s-ERROR-%v", msg, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))

	for _, tag := range tags {
		elments := strings.Split(tag, ":")
		result[strings.TrimSpace(elments[0])] = strings.TrimSpace(elments[1])
	}
	return result
}
