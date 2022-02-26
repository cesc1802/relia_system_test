package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
)

// LogConfiguration for logging
type LogConfig struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

type Logger struct {
	*zerolog.Logger
}

// LogConfigure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to LogConfiguration set.
func LogConfigure(logConfig LogConfig) *Logger {
	var writers []io.Writer

	if logConfig.ConsoleLoggingEnabled {
		//writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true})
		writers = append(writers, log.Output(os.Stderr))
	}
	if logConfig.FileLoggingEnabled {
		writers = append(writers, newRollingFile(logConfig))
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", logConfig.FileLoggingEnabled).
		Bool("jsonLogOutput", logConfig.EncodeLogsAsJson).
		Str("logDirectory", logConfig.Directory).
		Str("fileName", logConfig.Filename).
		Int("maxSizeMB", logConfig.MaxSize).
		Int("maxBackups", logConfig.MaxBackups).
		Int("maxAgeInDays", logConfig.MaxAge).
		Msg("logging LogConfigured")

	return &Logger{
		Logger: &logger,
	}
}

func newRollingFile(logConfig LogConfig) io.Writer {
	if err := os.MkdirAll(logConfig.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", logConfig.Directory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(logConfig.Directory, logConfig.Filename),
		MaxBackups: logConfig.MaxBackups, // files
		MaxSize:    logConfig.MaxSize,    // megabytes
		MaxAge:     logConfig.MaxAge,     // days
	}
}
