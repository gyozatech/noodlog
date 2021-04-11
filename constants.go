package noodlog

const (
	traceLabel = "trace"
	debugLabel = "debug"
	infoLabel  = "info"
	warnLabel  = "warn"
	errorLabel = "error"
	panicLabel = "panic"
	fatalLabel = "fatal"

	traceLevel = 1
	debugLevel = 2
	infoLevel  = 3
	warnLevel  = 4
	errorLevel = 5
	panicLevel = 6
	fatalLevel = 7

	defaultColor = "default"
	redColor     = "red"
	yellowColor  = "yellow"
	greenColor   = "green"
	blueColor    = "blue"
	purpleColor  = "purple"
	cyanColor    = "cyan"

	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)
