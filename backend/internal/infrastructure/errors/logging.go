package errors

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// LogLevel defines the level of logging
type LogLevel int

const (
	// LogLevelDebug logs everything
	LogLevelDebug LogLevel = iota
	// LogLevelInfo logs info, warnings and errors
	LogLevelInfo
	// LogLevelWarn logs warnings and errors
	LogLevelWarn
	// LogLevelError logs only errors
	LogLevelError
)

// LogConfig contains configuration for the logger
type LogConfig struct {
	// Level sets the logging level
	Level LogLevel
	// IncludeTimestamp includes timestamp in logs
	IncludeTimestamp bool
	// IncludeStackTrace includes stack trace for errors
	IncludeStackTrace bool
	// StackTraceSize sets the size of the stack trace (number of frames)
	StackTraceSize int
}

// DefaultLogConfig is the default logger configuration
var DefaultLogConfig = LogConfig{
	Level:            LogLevelInfo,
	IncludeTimestamp: true,
	IncludeStackTrace: true,
	StackTraceSize:   10,
}

// Logger is a custom logger for the application
type Logger struct {
	config LogConfig
	echo   *echo.Echo
}

// NewLogger creates a new logger with the given configuration
func NewLogger(config LogConfig) *Logger {
	return &Logger{
		config: config,
	}
}

// SetEcho sets the Echo instance for the logger
func (l *Logger) SetEcho(e *echo.Echo) {
	l.echo = e
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.config.Level <= LogLevelDebug {
		l.log(log.DEBUG, format, args...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	if l.config.Level <= LogLevelInfo {
		l.log(log.INFO, format, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.config.Level <= LogLevelWarn {
		l.log(log.WARN, format, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	if l.config.Level <= LogLevelError {
		l.log(log.ERROR, format, args...)
	}
}

// ErrorWithStack logs an error message with stack trace
func (l *Logger) ErrorWithStack(err error) {
	if l.config.Level <= LogLevelError {
		msg := err.Error()
		if l.config.IncludeStackTrace {
			msg = fmt.Sprintf("%s\n%s", msg, l.getStackTrace())
		}
		l.log(log.ERROR, msg)
	}
}

// log logs a message with the given level
func (l *Logger) log(level log.Lvl, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	
	if l.config.IncludeTimestamp {
		msg = fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), msg)
	}
	
	if l.echo != nil {
		// Echo logger doesn't have a Lvl method, so we just use the appropriate log level method
		switch level {
		case log.DEBUG:
			l.echo.Logger.Debug(msg)
		case log.INFO:
			l.echo.Logger.Info(msg)
		case log.WARN:
			l.echo.Logger.Warn(msg)
		case log.ERROR:
			l.echo.Logger.Error(msg)
		}
	} else {
		// Fallback to standard logging if Echo is not set
		fmt.Println(msg)
	}
}

// getStackTrace returns a formatted stack trace
func (l *Logger) getStackTrace() string {
	stackSize := l.config.StackTraceSize
	if stackSize <= 0 {
		stackSize = 10 // Default stack size
	}
	
	stack := make([]uintptr, stackSize)
	length := runtime.Callers(3, stack)
	stack = stack[:length]
	
	frames := runtime.CallersFrames(stack)
	var trace strings.Builder
	trace.WriteString("Stack trace:\n")
	
	for {
		frame, more := frames.Next()
		// Skip runtime and standard library frames
		if !strings.Contains(frame.File, "runtime/") {
			trace.WriteString(fmt.Sprintf("  %s:%d %s\n", frame.File, frame.Line, frame.Function))
		}
		if !more {
			break
		}
	}
	
	return trace.String()
}

// LogErrorWithContext logs an error with request context information
func LogErrorWithContext(c echo.Context, err error, statusCode int) {
	req := c.Request()
	logger := c.Logger()
	
	// Format the log message with request details
	logMsg := fmt.Sprintf("[%s] %s %s - Error: %v (Status: %d)",
		req.Method,
		req.Host,
		req.URL.Path,
		err,
		statusCode,
	)
	
	// Log with appropriate level based on status code
	if statusCode >= 500 {
		logger.Error(logMsg)
	} else if statusCode >= 400 {
		logger.Warn(logMsg)
	} else {
		logger.Info(logMsg)
	}
}