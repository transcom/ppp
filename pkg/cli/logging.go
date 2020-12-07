package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// LoggingEnvFlag is the logging environment flag
	LoggingEnvFlag string = "logging-env"
	// LogTaskMetadataFlag is the Log Task Metadata Flag
	LogTaskMetadataFlag string = "log-task-metadata"
	// LoggingLevelFlag is the flag that defines the logging level
	// Possible values are: fatal, error, warn, info, debug
	// The env var value is not case-sensitive. This works:
	// export LOGGING_LEVEL=INFO
	LoggingLevelFlag string = "logging-level"

	// LoggingEnvProduction is the production logging environment
	LoggingEnvProduction string = "production"
	// LoggingEnvDevelopment is the development logging environment
	LoggingEnvDevelopment string = "development"

	// LoggingLevelFatal is the fatal logging level
	LoggingLevelFatal string = "fatal"
	// LoggingLevelError is the error logging level
	LoggingLevelError string = "error"
	// LoggingLevelWarn is the warn logging level
	LoggingLevelWarn string = "warn"
	// LoggingLevelInfo is the info logging level
	LoggingLevelInfo string = "info"
	// LoggingLevelDebug is the debug logging level
	LoggingLevelDebug string = "debug"
)

var (
	allLoggingEnvs = []string{
		LoggingEnvProduction,
		LoggingEnvDevelopment,
	}
)

var (
	allLoggingLevels = []string{
		LoggingLevelFatal,
		LoggingLevelError,
		LoggingLevelWarn,
		LoggingLevelInfo,
		LoggingLevelDebug,
	}
)

type errInvalidLoggingEnv struct {
	Value       string
	LoggingEnvs []string
}

type errInvalidLoggingLevel struct {
	Value         string
	LoggingLevels []string
}

func (e *errInvalidLoggingEnv) Error() string {
	return fmt.Sprintf("invalid logging env %s, must be one of: ", e.Value) + strings.Join(e.LoggingEnvs, ", ")
}

func (e *errInvalidLoggingLevel) Error() string {
	return fmt.Sprintf("invalid logging level %s, must be one of: ", e.Value) + strings.Join(e.LoggingLevels, ", ")
}

// InitLoggingFlags initializes the logging command line flags
func InitLoggingFlags(flag *pflag.FlagSet) {
	flag.String(LoggingEnvFlag, LoggingEnvDevelopment, "logging environment: "+strings.Join(allLoggingEnvs, ", "))
	flag.Bool(LogTaskMetadataFlag, false, "Fetch AWS Task Metadata and add to log lines.")
	flag.String(LoggingLevelFlag, LoggingLevelInfo, "logging level: "+strings.Join(allLoggingLevels, ", "))
}

// CheckLogging validates logging command line flags
func CheckLogging(v *viper.Viper) error {
	if str := v.GetString(LoggingEnvFlag); !stringSliceContains(allLoggingEnvs, str) {
		return &errInvalidLoggingEnv{Value: str, LoggingEnvs: allLoggingEnvs}
	}
	if str := strings.ToLower(v.GetString(LoggingLevelFlag)); !stringSliceContains(allLoggingLevels, str) {
		return &errInvalidLoggingLevel{Value: str, LoggingLevels: allLoggingLevels}
	}
	return nil
}

// LogLevelIsDebug is a helper for functions that require a boolean to determine
// log verbosity
func LogLevelIsDebug(v *viper.Viper) bool {
	logLevel := strings.ToLower(v.GetString(LoggingLevelFlag))
	return logLevel == LoggingLevelDebug
}
