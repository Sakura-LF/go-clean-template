package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
}

// Logger -.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMilli},
	).With().Timestamp().Caller().Logger()

	return &Logger{
		logger: &logger,
	}
}

// Debug -.
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

// Info -.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

// Warn -.
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

// Error -.
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

// Fatal -.
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}
