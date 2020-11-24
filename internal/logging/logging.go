package logging

import "github.com/go-logr/logr"
import "go.uber.org/zap"
import "github.com/go-logr/zapr"

var (
	log logr.Logger
)

func init() {
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	log = zapr.NewLogger(zapLog)
}

// New create a new logger with a certain name
func New(name string) logr.Logger {
	return log.WithName(name)
}
