package noodlog

var logLevels = map[string]int{
	traceLabel: traceLevel,
	debugLabel: debugLevel,
	infoLabel:  infoLevel,
	warnLabel:  warnLevel,
	errorLabel: errorLevel,
	panicLabel: panicLevel,
	fatalLabel: fatalLevel,
}

func getLogLevel(level string) int {
	logLevel := logLevels[level]
	if logLevel == 0 {
		logLevel = infoLevel
	}
	return logLevel
}
