package noodlog

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

var colorMap = map[string]string{
   "trace": colorReset,
   "info": colorReset,
   "debug": colorGreen,
   "warn": colorYellow,
   "error": colorRed,
}

SetTraceColor(color string) {
   colorMap[traceLabel] = getColorByName(color)
}

SetDebugColor(color string) {
   colorMap[debugLabel] = getColorByName(color)
}

SetInfoColor(color string) {
   colorMap[infoLabel] = getColorByName(color)
}

SetWarnColor(color string) {
   colorMap[warnLabel] = getColorByName(color)
}

SetErrorColor(color string) {
   colorMap[errorLabel] = getColorByName(color)
}

func getColorByName(color) string {
   // TODO
   return ""
}




