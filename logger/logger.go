package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*Init creates a instance of logrus logger
JSON formatter is set is set for GCP cloud environments(PROD, DEV, STAGING)
JSON messages are indexed by GCP making them searchable on Cloud Logging
Text formatter is set for local environments

Levels:
logrus has seven log levels: Trace, Debug, Info, Warn, Error, Fatal, and Panic.
The severity of each level increases as you go down the list.
By default, logrus will log anything that is Info or above (Warn, Error, Fatal, or Panic).
*/
func Init() *logrus.Logger {
	log := logrus.New()
	env := os.Getenv("SHASTA_ENV")
	if true || env == "PROD" || env == "DEV" || env == "STAGING" {
		jsonFormatter := new(logrus.JSONFormatter)
		log.SetFormatter(jsonFormatter)
		switch env {
		case "PROD":
			log.SetLevel(logrus.InfoLevel)
		default:
			log.SetLevel(logrus.TraceLevel)
		}
	} else {
		textFormatter := new(logrus.TextFormatter)
		textFormatter.ForceColors = true
		textFormatter.FullTimestamp = true
		textFormatter.DisableLevelTruncation = true
		textFormatter.QuoteEmptyFields = true
		textFormatter.PadLevelText = true
		log.SetFormatter(textFormatter)
		log.SetLevel(logrus.TraceLevel)
	}
	return log
}

// //TODO return a request logger
// func RequestLogger(ctx context.Context) {

// }
