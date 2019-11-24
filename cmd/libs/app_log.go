package libs

import "github.com/sirupsen/logrus"

func LoadAppLog() *logrus.Logger {
	logger := logrus.New()
	return logger
}
