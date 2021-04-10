package noodlog

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)

var colorMap = map[string]string{
	"trace": colorReset,
	"info":  colorReset,
	"debug": colorGreen,
	"warn":  colorYellow,
	"error": colorRed,
}

var colorNames = map[string]string{
	"default": colorReset,
	"red":     colorRed,
	"green":   colorGreen,
	"yellow":  colorYellow,
	"blue":    colorBlue,
	"purple":  colorPurple,
	"cyan":    colorCyan,
}

var colorEnabled = false

// EnableColors function enables colored logging
func EnableColors() {
	colorEnabled = true
}

// EnableColors function disables colored logging
func DisableColors() {
	colorEnabled = true
}

// SetTraceColor overrides the trace level log color with the one specified in input
func SetTraceColor(color string) {
	colorMap[traceLabel] = getColorByName(color)
}

// SetDebugColor overrides the debug level log color with the one specified in input
func SetDebugColor(color string) {
	colorMap[debugLabel] = getColorByName(color)
}

// SetInfoColor overrides the info level log color with the one specified in input
func SetInfoColor(color string) {
	colorMap[infoLabel] = getColorByName(color)
}

// SetWarnColor overrides the warn level log color with the one specified in input
func SetWarnColor(color string) {
	colorMap[warnLabel] = getColorByName(color)
}

// SetErrorColor overrides the error level log color with the one specified in input
func SetErrorColor(color string) {
	colorMap[errorLabel] = getColorByName(color)
}

func getColorByName(color string) string {
	colorCode := colorNames[color]
	if colorCode == "" {
		colorCode = colorReset
	}
	return colorCode
}

func setCustomColors(colors CustomColors) {
	if colors.Trace != nil {
		SetTraceColor(*colors.Trace)
	}
	if colors.Debug != nil {
		SetDebugColor(*colors.Debug)
	}
	if colors.Info != nil {
		SetInfoColor(*colors.Info)
	}
	if colors.Warn != nil {
		SetWarnColor(*colors.Warn)
	}
	if colors.Error != nil {
		SetErrorColor(*colors.Error)
	}
}
