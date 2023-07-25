package goesl

import (
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger

func init() {
	lvl, err := logrus.ParseLevel(os.Getenv("GOESL_LOGLEVEL"))
	if err != nil {
		logrus.SetLevel(logrus.ErrorLevel)
	}
	logrus.SetLevel(lvl)
	logrus.SetFormatter(&nested.Formatter{
		FieldsOrder:   []string{"name"},
		HideKeys:      true,
		TrimMessages:  true,
		ShowFullLevel: true,
	})
	logrus.SetOutput(os.Stdout)
	log = logrus.WithField("name", "GOESL")
}

func Info(msg ...any) {
	log.Info(msg...)
}

func Infof(format string, msg ...any) {
	log.Infof(format, msg...)
}

func Debug(msg ...any) {
	log.Debug(msg...)
}

func Debugf(format string, msg ...any) {
	log.Debugf(format, msg...)
}

func Warn(msg ...any) {
	log.Warn(msg...)
}

func Warnf(format string, msg ...any) {
	log.Warnf(format, msg...)
}

func Error(msg ...any) {
	log.Error(msg...)
}

func Errorf(format string, msg ...any) {
	log.Errorf(format, msg...)
}
