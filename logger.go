package noodlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Logger represent the logger object
type Logger struct {
	level                int
	logWriter            io.Writer
	prettyPrint          bool
	traceCaller          bool
	traceCallerLevel     int
	obscureSensitiveData bool
	sensitiveParams      []string
	colors               bool
	colorMap             map[string]string
}

// NewLogger func is the default constructor of a Logger
func NewLogger() *Logger {
	return &Logger{
		level:                infoLevel,
		logWriter:            os.Stdout,
		prettyPrint:          false,
		traceCaller:          false,
		traceCallerLevel:     5,
		obscureSensitiveData: false,
		sensitiveParams:      nil,
		colors:               false,
		colorMap:             colorMap,
	}
}

// SetConfigs function allows you to rewrite all the configs at once
func (l *Logger) SetConfigs(configs Configs) *Logger {
	if configs.LogLevel != nil {
		l.Level(*configs.LogLevel)
	}
	if configs.LogWriter != nil {
		l.LogWriter(configs.LogWriter)
	}
	if configs.JSONPrettyPrint != nil {
		if *configs.JSONPrettyPrint {
			l.EnableJSONPrettyPrint()
		} else {
			l.DisableJSONPrettyPrint()
		}
	}
	if configs.TraceCaller != nil {
		if *configs.TraceCaller {
			l.EnableTraceCaller()
		} else {
			l.DisableTraceCaller()
		}
	}
	if configs.SinglePointTracing != nil {
		if *configs.SinglePointTracing {
			l.EnableSinglePointTracing()
		} else {
			l.DisableSinglePointTracing()
		}
	}
	if configs.Colors != nil {
		if *configs.Colors {
			l.EnableColors()
		} else {
			l.DisableColors()
		}
	}
	if configs.CustomColors != nil {
		l.SetCustomColors(*configs.CustomColors)
	}
	if configs.ObscureSensitiveData != nil {
		if *configs.ObscureSensitiveData {
			l.EnableObscureSensitiveData(nil)
		} else {
			l.DisableObscureSensitiveData()
		}
	}
	if configs.SensitiveParams != nil {
		l.SetSensitiveParams(configs.SensitiveParams)
	}
	return l
}

// Level func let you establish the log level for a specified logger instance
func (l *Logger) Level(level string) *Logger {
	l.level = getLogLevel(level)
	return l
}

// LogWriter function let you define a logWriter (os.Stdout, a file, a buffer etc.)
func (l *Logger) LogWriter(w io.Writer) *Logger {
	l.logWriter = w
	return l
}

// EnableJSONPrettyPrint func let you enable JSON pretty printing for the specified logger instance
func (l *Logger) EnableJSONPrettyPrint() *Logger {
	l.prettyPrint = true
	return l
}

// DisableJSONPrettyPrint func let you disable JSON pretty printing for the specified logger instance
func (l *Logger) DisableJSONPrettyPrint() *Logger {
	l.prettyPrint = false
	return l
}

// EnableTraceCaller enables the tracing of the caller for the specified logger instance
func (l *Logger) EnableTraceCaller() *Logger {
	l.traceCaller = true
	return l
}

// DisableTraceCaller disables the tracing of the caller for the specified logger instance
func (l *Logger) DisableTraceCaller() *Logger {
	l.traceCaller = false
	return l
}

// EnableSinglePointTracing function enables tracing the caller when setting the logger in a single package for the whole project and recalling the logging for the project from that single point for the specified logger instance
func (l *Logger) EnableSinglePointTracing() *Logger {
	l.traceCaller = true
	l.traceCallerLevel = 6
	return l
}

// DisableSinglePointTracing function trace function and filename of the directl caller
func (l *Logger) DisableSinglePointTracing() *Logger {
	l.traceCallerLevel = 5
	return l
}

// EnableColors function let you enable colored logs for a specified logger
func (l *Logger) EnableColors() *Logger {
	l.colors = true
	return l
}

// DisableColors function let you disable colored logs for a specified logger
func (l *Logger) DisableColors() *Logger {
	l.colors = false
	return l
}

// SetCustomColors overrides defaultColor when custom color is passed into CustomColor configs
func (l *Logger) SetCustomColors(colors CustomColors) *Logger {

	empty := Color{}
	if traceColor := detectColor(colors.Trace); traceColor != empty {
		l.SetTraceColor(traceColor)
	}
	if debugColor := detectColor(colors.Debug); debugColor != empty {
		l.SetDebugColor(debugColor)
	}
	if infoColor := detectColor(colors.Info); infoColor != empty {
		l.SetInfoColor(infoColor)
	}
	if warnColor := detectColor(colors.Warn); warnColor != empty {
		l.SetWarnColor(warnColor)
	}
	if errorColor := detectColor(colors.Error); errorColor != empty {
		l.SetErrorColor(errorColor)
	}

	return l
}

// SetTraceColor overrides the trace level log color with the one specified in input
func (l *Logger) SetTraceColor(color Color) {
	l.colorMap[traceLabel] = color.toCode()
}

// SetDebugColor overrides the debug level log color with the one specified in input
func (l *Logger) SetDebugColor(color Color) {
	l.colorMap[debugLabel] = color.toCode()
}

// SetInfoColor overrides the info level log color with the one specified in input
func (l *Logger) SetInfoColor(color Color) {
	l.colorMap[infoLabel] = color.toCode()
}

// SetWarnColor overrides the warn level log color with the one specified in input
func (l *Logger) SetWarnColor(color Color) {
	l.colorMap[warnLabel] = color.toCode()
}

// SetErrorColor overrides the error level log color with the one specified in input
func (l *Logger) SetErrorColor(color Color) {
	l.colorMap[errorLabel] = color.toCode()
}

// EnableObscureSensitiveData enables sensitive data obscuration from json logs for a given logger instance
func (l *Logger) EnableObscureSensitiveData(params []string) *Logger {
	l.obscureSensitiveData = true
	l.sensitiveParams = params
	return l
}

// DisableObscureSensitiveData disables sensitive data obscuration from json logs for a given logger instance
func (l *Logger) DisableObscureSensitiveData() *Logger {
	l.obscureSensitiveData = false
	return l
}

// SetSensitiveParams sets sensitive data obscuration from json logs
func (l *Logger) SetSensitiveParams(params []string) *Logger {
	l.sensitiveParams = params
	return l
}

// Trace function prints a log with trace log level
func (l *Logger) Trace(message ...interface{}) {
	l.printLog(traceLabel, message)
}

// Debug function prints a log with debug log level
func (l *Logger) Debug(message ...interface{}) {
	l.printLog(debugLabel, message)
}

// Info function prints a log with info log level
func (l *Logger) Info(message ...interface{}) {
	l.printLog(infoLabel, message)
}

// Warn function prints a log with warn log level
func (l *Logger) Warn(message ...interface{}) {
	l.printLog(warnLabel, message)
}

// Error function prints a log with error log level
func (l *Logger) Error(message ...interface{}) {
	l.printLog(errorLabel, message)
}

// Panic function prints a log with panic log level
func (l *Logger) Panic(message ...interface{}) {
	panic(l.composeLog(panicLabel, message))
}

// Fatal function prints a log with fatal log level
func (l *Logger) Fatal(message ...interface{}) {
	l.printLog(fatalLabel, message)
	if os.Getenv("EXIT_ON_FATAL_DISABLED") != "true" {
		os.Exit(1)
	}
}

func (l *Logger) printLog(label string, message []interface{}) {
	if logLevels[label] >= l.level {
		fmt.Fprintln(l.logWriter, l.composeLog(label, message))
	}
}

func (l *Logger) composeLog(level string, message []interface{}) string {

	logMsg := record{
		Level:   level,
		Message: l.composeMessage(message),
		Time:    strings.Split(time.Now().String(), "m")[0],
	}

	if l.traceCaller {
		f, fx := traceCaller(l.traceCallerLevel)
		logMsg.File = &f
		logMsg.Function = &fx
	}

	var jsn []byte
	if l.prettyPrint {
		jsn, _ = json.MarshalIndent(logMsg, "", "   ")
	} else {
		jsn, _ = json.Marshal(logMsg)
	}

	logRecord := string(jsn)
	if l.colors {
		logRecord = fmt.Sprintf("%s%s%s", l.colorMap[level], logRecord, colorReset)
	}

	return logRecord
}

func (l *Logger) composeMessage(message []interface{}) interface{} {
	switch len(message) {
	case 0:
		return ""
	case 1:
		return l.adaptMessage(message[0])
	default:
		switch message[0].(type) {
		case string:
			msg0 := message[0].(string)
			if strings.Contains(msg0, "%") {
				return fmt.Sprintf(msg0, message[1:]...)
			}
		}
		return stringify(message)
	}
}

func (l *Logger) adaptMessage(message interface{}) interface{} {
	switch message.(type) {
	case string:
		strMsg := message.(string)
		if l.obscureSensitiveData && len(l.sensitiveParams) != 0 {
			return strToObj(obscureParams(strMsg, l.sensitiveParams))
		}
		return strToObj(strMsg)
	default:
		if l.obscureSensitiveData && len(l.sensitiveParams) != 0 {
			jsn, _ := json.Marshal(message)
			strMsg := obscureParams(string(jsn), l.sensitiveParams)
			return strToObj(strMsg)
		}
	}
	return message
}
