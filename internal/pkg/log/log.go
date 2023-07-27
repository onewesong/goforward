package log

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

// Log is the under log object
var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(200)
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(Log.GetLogFuncCallDepth() + 1)
}

func InitLog(logWay, logFile, logLevel string, maxdays int64, disableLogColor bool) {
	SetLogFile(logWay, logFile, maxdays, disableLogColor)
	SetLogLevel(logLevel)
}

// SetLogFile to configure log params
// logWay: file or console
func SetLogFile(logWay, logFile string, maxdays int64, disableLogColor bool) {
	if logWay == "console" {
		params := ""
		if disableLogColor {
			params = fmt.Sprintf(`{"color": false}`)
		}
		Log.SetLogger("console", params)
	} else {
		params := fmt.Sprintf(`{"filename": "%q", "maxdays": %d}`, logFile, maxdays)
		Log.SetLogger("file", params)
	}
}

// SetLogLevel set log level, default is warning
// value: error, warning, info, debug, trace
func SetLogLevel(logLevel string) {
	level := 4 // warning
	switch logLevel {
	case "error":
		level = 3
	case "warn":
		level = 4
	case "info":
		level = 6
	case "debug":
		level = 7
	case "trace":
		level = 8
	default:
		level = 4
	}
	Log.SetLevel(level)
}

// wrap log

func Fatal(format string, v ...any) {
	Log.Critical(format, v...)
	os.Exit(1)
}

func Alert(format string, v ...any) {
	Log.Alert(format, v...)
}

// Critical logs a message at critical level.
func Critical(format string, v ...any) {
	Log.Critical(format, v...)
}

func Error(format string, v ...any) {
	Log.Error(format, v...)
}

func Warn(format string, v ...any) {
	Log.Warn(format, v...)
}

func Info(format string, v ...any) {
	Log.Info(format, v...)
}

func Debug(format string, v ...any) {
	Log.Debug(format, v...)
}

func Trace(format string, v ...any) {
	Log.Trace(format, v...)
}
