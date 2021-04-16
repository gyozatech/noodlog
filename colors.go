package noodlog

import "strconv"

type Color struct {
	Code *string
}

var colorMap = map[string]string{
	traceLabel: resetColor,
	infoLabel:  resetColor,
	debugLabel: NewColor(Green).ToCode(),
	warnLabel:  NewColor(Yellow).ToCode(),
	errorLabel: NewColor(Red).ToCode(),
}

var colors = map[string]string{
	defaultColor: resetColor,
	redColor:     colorRed,
	greenColor:   colorGreen,
	yellowColor:  colorYellow,
	blueColor:    colorBlue,
	purpleColor:  colorPurple,
	cyanColor:    colorCyan,
}

var backgroundColors = map[string]string{
	defaultColor: resetColor,
	redColor:     backgroundRed,
	greenColor:   backgroundGreen,
	yellowColor:  backgroundYellow,
	blueColor:    backgroundBlue,
	purpleColor:  backgroundPurple,
	cyanColor:    backgroundCyan,
}

var colorEnabled = false

// IsValidTrueColor check if a true color is valid, it has to be included between 0 and 255
func IsValidTrueColor(color int) bool {
	return color >= 0 && color <= 255
}

// NewColor set the color of the text using a string as identifiers
func NewColor(color *string) Color {
	code := resetColor
	colorCode := colors[*color]
	if colorCode != "" {
		code = "\033[" + colorCode + "m"
	}
	return Color{Code: &code}
}

// NewColorRGB set the color of the text using RGB Notations
func NewColorRGB(red int, green int, blue int) Color {
	code := resetColor
	if IsValidTrueColor(red) && IsValidTrueColor(green) && IsValidTrueColor(blue) {
		code = "\033[38;2;" + strconv.Itoa(red) + ";" + strconv.Itoa(green) + ";" + strconv.Itoa(blue) + "m"
	}
	return Color{Code: &code}
}

// Background set the background using a string as identifiers
func Background(color *string) Color {
	code := resetColor
	colorCode := backgroundColors[*color]
	if colorCode != "" {
		code = "\033[" + colorCode + "m"
	}
	return Color{Code: &code}
}

// BackgroundRGB set the background using RGB Notations
func BackgroundRGB(red int, green int, blue int) Color {
	code := resetColor
	if IsValidTrueColor(red) && IsValidTrueColor(green) && IsValidTrueColor(blue) {
		code = "\033[48;2;" + strconv.Itoa(red) + ";" + strconv.Itoa(green) + ";" + strconv.Itoa(blue) + "m"
	}
	return Color{Code: &code}
}

// From a given Color it set the background using a string as identifier
func (c Color) Background(color *string) Color {
	code := *(c.Code) + *Background(color).Code
	return Color{Code: &code}
}

// From a given Color it set the background using RGB Notations
func (c Color) BackgroundRGB(red int, green int, blue int) Color {
	code := *(c.Code) + *BackgroundRGB(red, green, blue).Code
	return Color{Code: &code}
}

// ToCode returns code used into bash for colored log
func (c Color) ToCode() string {
	if c.Code != nil {
		return *c.Code
	}
	return resetColor
}

// EnableColors function enables colored logging
func EnableColors() {
	colorEnabled = true
}

// DisableColors function disables colored logging
func DisableColors() {
	colorEnabled = false
}

// SetTraceColor overrides the trace level log color with the one specified in input
func SetTraceColor(color Color) {
	colorMap[traceLabel] = color.ToCode()
}

// SetDebugColor overrides the debug level log color with the one specified in input
func SetDebugColor(color Color) {
	colorMap[debugLabel] = color.ToCode()
}

// SetInfoColor overrides the info level log color with the one specified in input
func SetInfoColor(color Color) {
	colorMap[infoLabel] = color.ToCode()
}

// SetWarnColor overrides the warn level log color with the one specified in input
func SetWarnColor(color Color) {
	colorMap[warnLabel] = color.ToCode()
}

// SetErrorColor overrides the error level log color with the one specified in input
func SetErrorColor(color Color) {
	colorMap[errorLabel] = color.ToCode()
}

func setCustomColors(colors CustomColors) {
	var empty = Color{}
	if colors.Trace != empty {
		SetTraceColor(colors.Trace)
	}
	if colors.Debug != empty {
		SetDebugColor(colors.Debug)
	}
	if colors.Info != empty {
		SetInfoColor(colors.Info)
	}
	if colors.Warn != empty {
		SetWarnColor(colors.Warn)
	}
	if colors.Error != empty {
		SetErrorColor(colors.Error)
	}
}
