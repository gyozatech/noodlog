package noodlog

import "strconv"

type Color struct {
	Code *string
}

var colorMap = map[string]string{
	traceLabel: colorReset,
	infoLabel:  colorReset,
	debugLabel: NewColor(Green).toCode(),
	warnLabel:  NewColor(Yellow).toCode(),
	errorLabel: NewColor(Red).toCode(),
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

// isValidColor check if a true color is valid, it has to be included between 0 and 255
func isValidColor(color int) bool {
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
	if isValidColor(red) && isValidColor(green) && isValidColor(blue) {
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
func backgroundRGB(red int, green int, blue int) Color {
	code := colorReset
	if isValidColor(red) && isValidColor(green) && isValidColor(blue) {
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
	code := *(c.Code) + *backgroundRGB(red, green, blue).Code
	return Color{Code: &code}
}

// toCode returns code used into bash for colored log
func (c Color) toCode() string {
	if c.Code != nil {
		return *c.Code
	}
	return colorReset
}

// detectColor return Color struct given a generic interface{}. It is an empty Color struct if not a valid type
func detectColor(color interface{}) Color {
	empty := Color{}
	switch color.(type) {
	case Color:
		if color != empty {
			return color.(Color)
		}
		return empty
	case *string:
		return NewColor(color.(*string))

	default:
		return empty
	}
}
