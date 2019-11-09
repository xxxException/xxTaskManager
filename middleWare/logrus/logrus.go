package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func NewLogger() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
}
