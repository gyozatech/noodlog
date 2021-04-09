package noodlog

var traceCaller = false
var traceLevel = 3

func SetTraceCaller(enabled bool) {
   traceCaller = enabled
}

// SetSinglePointTracingON function enables tracing the caller when setting the logger in a single package for the whole project and recalling the logging for the project from that single point
func SetSinglePointTracingON() {
   traceLevel = 4
}

// SetSinglePointTracingOFF function trace function and filename of the directl caller
func SetSinglePointTracingOFF() {
   traceLevel = 3
}

// traceCaller static functions retrieves the filename of the function which wants to log
func traceCaller() map[string]string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(traceLevel, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return map[string]string{
		file:     fmt.Sprintf("%s:%d", frame.File, frame.Line),
		function: frame.Function,
	}
}
