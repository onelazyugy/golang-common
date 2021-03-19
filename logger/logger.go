package logger

import (
	"github.com/onelazyugy/golang-common/constants"
	"github.com/sirupsen/logrus"
)

// New returns a logrus logger based off a profile.
// For DEV, STAGING, and PROD profile environments, a json formatter is used.
// For any other profile, a text formatter is used to pretty print logs.
//
// Levels:
// logrus has seven log levels: Trace, Debug, Info, Warn, Error, Fatal, and Panic.
// The severity of each level increases as you go down the list.
// By default, logrus will log anything that is Info or above (Warn, Error, Fatal, or Panic).
func New(profile string) *logrus.Logger {
	log := logrus.New()
	f, l := newLoggerComponents(profile)
	log.SetFormatter(f)
	log.SetLevel(l)

	return log
}

func newLoggerComponents(profile string) (logrus.Formatter, logrus.Level) {
	var f logrus.Formatter

	l := logrus.TraceLevel

	switch profile {
	case constants.PROD:
		l = logrus.InfoLevel
		fallthrough
	case constants.DEV, constants.STAGING, "": // empty string case to account for unset env vars, TODO: remove once profile is properly set
		f = &logrus.JSONFormatter{}
	default:
		f = &logrus.TextFormatter{
			ForceColors:            true,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
			PadLevelText:           true,
		}
	}

	return f, l
}
