package utils

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestXX(t *testing.T) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter) // default
	log.SetOutput(os.Stdout)
	log.WithFields(logrus.Fields{
		"a": "a",
		"b": "b",
	}).Trace("hello world")
}
