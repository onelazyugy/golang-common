package constants

import "os"

type GCPStruct struct {
	ProjectID string
}

type ENVStruct struct {
	Profile string //Env the code is running in(PROD, STAGING, DEV)
}

// App constants
const (
	Token           = "token"
	CorrelationID   = "nep-correlation-id"
	TimestampFormat = "01-02-2006T15:04:05" // will produce this format 02-03-2021 T16:36:42

	// Profiles
	DEV     = "DEV"
	STAGING = "STAGING"
	PROD    = "PROD"
)

//ENV and GCP Constants
var ENV ENVStruct
var GCP GCPStruct

func init() {
	ENV.Profile = os.Getenv("profile")
	GCP.ProjectID = os.Getenv("GCP_Project")
}
