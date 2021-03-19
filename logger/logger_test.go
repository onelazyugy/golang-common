package logger

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/onelazyugy/golang-common/constants"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		testName    string
		profile     string
		expectedOut string
	}{
		{
			testName:    "logs for DEV",
			profile:     constants.DEV,
			expectedOut: `^{\"level\":\"info\",\"msg\":\"test-info\",\"time\":\".*\"}\n{\"level\":\"trace\",\"msg\":\"test-trace\",\"time\":\".*\"}\n$`,
		},
		{
			testName:    "logs for STAGING",
			profile:     constants.STAGING,
			expectedOut: `^{\"level\":\"info\",\"msg\":\"test-info\",\"time\":\".*\"}\n{\"level\":\"trace\",\"msg\":\"test-trace\",\"time\":\".*\"}\n$`,
		},
		{
			testName:    "logs for PROD",
			profile:     constants.PROD,
			expectedOut: `^{\"level\":\"info\",\"msg\":\"test-info\",\"time\":\".*\"}\n$`,
		},
		{
			testName:    "pretty logs",
			profile:     "UNKNOWN",
			expectedOut: `^\x1b[36mINFO.*\x1b[0m\\[.*\\].*test-info.*\n\x1b[37mTRACE.*\x1b[0m\\[.*\\].*test-trace.*\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			log := New(tt.profile)
			buf := &bytes.Buffer{}
			log.SetOutput(buf)

			log.Info("test-info")
			log.Trace("test-trace")

			require.Regexp(t, regexp.MustCompile(tt.expectedOut), buf.String())
		})
	}
}
