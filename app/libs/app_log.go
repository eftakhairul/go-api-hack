package libs

import "github.com/sirupsen/logrus"

// LoadAppLog return new logrus.Logger
func LoadAppLog() *logrus.Logger {
	logger := logrus.New()
	return logger
}
