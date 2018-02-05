package logger

import (
	"io"
	"log"

	"github.com/sirupsen/logrus"
	"source.golabs.io/core/shipper/pkg/config"
)

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
}

type logger struct {
	l *logrus.Logger
}

// NewLogger - create a new logrus logger
func NewLogger(config *config.Config, w io.Writer) Logger {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		log.Fatalf(err.Error())
	}

	l := &logrus.Logger{
		// Out:       os.Stdout,
		Out:       w,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
		Formatter: &logrus.TextFormatter{},
	}

	return &logger{
		l: l,
	}
}

func (l *logger) Errorln(args ...interface{}) {
	l.l.Errorln(args...)
}

func (l *logger) Infoln(args ...interface{}) {
	l.l.Infoln(args...)
}
