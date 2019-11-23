package libs

import "github.com/sirupsen/logrus"

func LoadAppLog() *Logger {
	logger := logrus.New()
	return logger
}
