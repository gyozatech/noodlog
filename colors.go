package noodlog

var colorMap = map[string]string{
	traceLabel: colorReset,
	infoLabel:  colorReset,
	debugLabel: colorGreen,
	warnLabel:  colorYellow,
	errorLabel: colorRed,
}

var colors = map[string]string{
	defaultColor: colorReset,
	redColor:     colorRed,
	greenColor:   colorGreen,
	yellowColor:  colorYellow,
	blueColor:    colorBlue,
	purpleColor:  colorPurple,
	cyanColor:    colorCyan,
}

var colorEnabled = false

// EnableColors function enables colored logging
func EnableColors() {
	colorEnabled = true
}

// DisableColors function disables colored logging
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
	colorCode := colors[color]
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
