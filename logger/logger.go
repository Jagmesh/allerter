package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() *logrus.Logger {
	log = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		},
		Level: logrus.DebugLevel,
	}

	return log
}

func GetLogger() *logrus.Logger {
	return log
}
