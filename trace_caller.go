package noodlog

import (
	"fmt"
	"runtime"
)

var traceCallerEnabled = false
var traceCallerLevel = 3

// EnableTraceCaller enables the tracing of the caller
func EnableTraceCaller() {
	traceCallerEnabled = true
}

// DisableTraceCaller disables the tracing of the caller
func DisableTraceCaller() {
	traceCallerEnabled = false
}

// EnableSinglePointTracing function enables tracing the caller when setting the logger in a single package for the whole project and recalling the logging for the project from that single point
func EnableSinglePointTracing() {
	traceCallerLevel = 4
}

// DisableSinglePointTracing function trace function and filename of the directl caller
func DisableSinglePointTracing() {
	traceCallerLevel = 3
}

// traceCaller static functions retrieves the filename of the function which wants to log
func traceCaller() map[string]string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(traceCallerLevel, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return map[string]string{
		file:     fmt.Sprintf("%s:%d", frame.File, frame.Line),
		function: frame.Function,
	}
}
