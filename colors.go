package noodlog

import "strconv"

type Color struct {
	Code *string
}

var colorMap = map[string]string{
	traceLabel: colorReset,
	infoLabel:  colorReset,
	debugLabel: NewColor(Green).ToCode(),
	warnLabel:  NewColor(Yellow).ToCode(),
	errorLabel: NewColor(Red).ToCode(),
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

var backgroundColors = map[string]string{
	defaultColor: colorReset,
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
	code := colorReset
	colorCode := colors[*color]
	if colorCode != "" {
		code = "\033[" + colorCode + "m"
	}
	return Color{Code: &code}
}

// NewColorRGB set the color of the text using RGB Notations
func NewColorRGB(red int, green int, blue int) Color {
	code := colorReset
	if IsValidTrueColor(red) && IsValidTrueColor(green) && IsValidTrueColor(blue) {
		code = "\033[38;2;" + strconv.Itoa(red) + ";" + strconv.Itoa(green) + ";" + strconv.Itoa(blue) + "m"
	}
	return Color{Code: &code}
}

// Background set the background using a string as identifiers
func Background(color *string) Color {
	code := colorReset
	colorCode := backgroundColors[*color]
	if colorCode != "" {
		code = "\033[" + colorCode + "m"
	}
	return Color{Code: &code}
}

// BackgroundRGB set the background using RGB Notations
func BackgroundRGB(red int, green int, blue int) Color {
	code := colorReset
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
	return colorReset
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

// DetectColor return Color struct given a generic interface{}. It is an empty Color struct if not a valid type
func DetectColor(color interface{}) Color {
	empty := Color{}
	switch color.(type) {
	case *string:
		return NewColor(color.(*string))
	case Color:
		if color != empty {
			return color.(Color)
		}
		return empty
	default:
		return empty
	}
}

// setCustomColors overrides defaultColor when custom color is passed into CustomColor configs
func setCustomColors(colors CustomColors) {
	empty := Color{}
	if traceColor := DetectColor(colors.Trace); traceColor != empty {
		SetTraceColor(traceColor)
	}
	if debugColor := DetectColor(colors.Debug); debugColor != empty {
		SetDebugColor(debugColor)
	}
	if infoColor := DetectColor(colors.Info); infoColor != empty {
		SetInfoColor(infoColor)
	}
	if warnColor := DetectColor(colors.Warn); warnColor != empty {
		SetWarnColor(warnColor)
	}
	if errorColor := DetectColor(colors.Error); errorColor != empty {
		SetErrorColor(errorColor)
	}
}
