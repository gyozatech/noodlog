package noodlog

import (
	"testing"
)

var (
	redCode    = "\033[41m"
	greenCode  = "\033[42m"
	yellowCode = "\033[43m"
	blueCode   = "\033[44m"
	purpleCode = "\033[45m"
	cyanCode   = "\033[46m"
	resetCode  = "\033[0m"
)

var testData = map[string]Color{
	redColor:       {Code: &redCode},
	greenColor:     {Code: &greenCode},
	yellowColor:    {Code: &yellowCode},
	blueColor:      {Code: &blueCode},
	purpleColor:    {Code: &purpleCode},
	cyanColor:      {Code: &cyanCode},
	"invalidColor": {Code: &resetCode},
}

func TestIsValidColor(t *testing.T) {
	errFormat := "TestIsValidColor failed: expected true found %t"

	validColor := 255
	negativeColor := -1
	overflowColor := 256

	if actual := isValidColor(validColor); !actual {
		t.Errorf(errFormat, actual)
	}

	if actual := isValidColor(negativeColor); actual {
		t.Errorf(errFormat, actual)
	}

	if actual := isValidColor(overflowColor); actual {
		t.Errorf(errFormat, actual)
	}
}

func TestNewColor(t *testing.T) {
	var (
		redCode    = "\033[31m"
		greenCode  = "\033[32m"
		yellowCode = "\033[33m"
		blueCode   = "\033[34m"
		purpleCode = "\033[35m"
		cyanCode   = "\033[36m"
		resetCode  = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       {Code: &redCode},
		greenColor:     {Code: &greenCode},
		yellowColor:    {Code: &yellowCode},
		blueColor:      {Code: &blueCode},
		purpleColor:    {Code: &purpleCode},
		cyanColor:      {Code: &cyanCode},
		"invalidColor": {Code: &resetCode},
	}

	for label, expected := range testData {
		if actual := NewColor(&label); *actual.Code != *expected.Code {
			t.Errorf("TestNewColor failed for color %v: got %v, expected %v", label, actual, expected)
		}
	}
}

func TestNewColorRGB(t *testing.T) {
	errFormat := "TestNewColorRGB failed for color %v: got %v, expected %v"

	var (
		redCode     = "\033[38;2;255;0;0m"
		greenCode   = "\033[38;2;0;255;0m"
		blueCode    = "\033[38;2;0;0;255m"
		mixedCode   = "\033[38;2;128;128;128m"
		invalidCode = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       {Code: &redCode},
		greenColor:     {Code: &greenCode},
		blueColor:      {Code: &blueCode},
		"mixedColor":   {Code: &mixedCode},
		"invalidColor": {Code: &invalidCode},
	}

	if actual := NewColorRGB(255, 0, 0); *actual.Code != *testData[redColor].Code {
		t.Errorf(errFormat, redColor, actual, testData[redColor])
	}

	if actual := NewColorRGB(0, 255, 0); *actual.Code != *testData[greenColor].Code {
		t.Errorf(errFormat, greenColor, actual, testData[greenColor])
	}

	if actual := NewColorRGB(0, 0, 255); *actual.Code != *testData[blueColor].Code {
		t.Errorf(errFormat, blueColor, actual, testData[blueColor])
	}

	if actual := NewColorRGB(128, 128, 128); *actual.Code != *testData["mixedColor"].Code {
		t.Errorf(errFormat, "mixedColor", actual, testData["mixedColor"])
	}

	if actual := NewColorRGB(256, 0, 0); *actual.Code != *testData["invalidColor"].Code {
		t.Errorf(errFormat, "invalidColor", actual, testData["invalidColor"])
	}
}

func TestBackground(t *testing.T) {
	var (
		redCode    = "\033[41m"
		greenCode  = "\033[42m"
		yellowCode = "\033[43m"
		blueCode   = "\033[44m"
		purpleCode = "\033[45m"
		cyanCode   = "\033[46m"
		resetCode  = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       {Code: &redCode},
		greenColor:     {Code: &greenCode},
		yellowColor:    {Code: &yellowCode},
		blueColor:      {Code: &blueCode},
		purpleColor:    {Code: &purpleCode},
		cyanColor:      {Code: &cyanCode},
		"invalidColor": {Code: &resetCode},
	}

	for label, expected := range testData {
		if actual := Background(&label); *actual.Code != *expected.Code {
			t.Errorf("TestBackground failed for color %v: got %v, expected %v", label, actual, expected)
		}
	}
}

func TestBackgroundRGB(t *testing.T) {
	errFormat := "TestBackgroundRGB failed for color %v: got %v, expected %v"

	var (
		redCode     = "\033[48;2;255;0;0m"
		greenCode   = "\033[48;2;0;255;0m"
		blueCode    = "\033[48;2;0;0;255m"
		mixedCode   = "\033[48;2;128;128;128m"
		invalidCode = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       {Code: &redCode},
		greenColor:     {Code: &greenCode},
		blueColor:      {Code: &blueCode},
		"mixedColor":   {Code: &mixedCode},
		"invalidColor": {Code: &invalidCode},
	}

	if actual := backgroundRGB(255, 0, 0); *actual.Code != *testData[redColor].Code {
		t.Errorf(errFormat, redColor, actual, testData[redColor])
	}

	if actual := backgroundRGB(0, 255, 0); *actual.Code != *testData[greenColor].Code {
		t.Errorf(errFormat, greenColor, actual, testData[greenColor])
	}

	if actual := backgroundRGB(0, 0, 255); *actual.Code != *testData[blueColor].Code {
		t.Errorf(errFormat, blueColor, actual, testData[blueColor])
	}

	if actual := backgroundRGB(128, 128, 128); *actual.Code != *testData["mixedColor"].Code {
		t.Errorf(errFormat, "mixedColor", actual, testData["mixedColor"])
	}

	if actual := backgroundRGB(256, 0, 0); *actual.Code != *testData["invalidColor"].Code {
		t.Errorf(errFormat, "invalidColor", actual, testData["invalidColor"])
	}
}

func TestToCode(t *testing.T) {
	errFormat := "TestToCode failed for empty Color: got %s, expected %s"

	empty := Color{}
	if actual := empty.toCode(); actual != colorReset {
		t.Errorf(errFormat, actual, colorReset)
	}

	existingCode := "code"
	existing := Color{Code: &existingCode}

	if actual := existing.toCode(); actual != existingCode {
		t.Errorf(errFormat, actual, existingCode)
	}
}

func TestBackgroundOnColor(t *testing.T) {

	var existingCode = "code"
	color := Color{Code: &existingCode}

	for label, expected := range testData {
		if actual := color.Background(&label); *actual.Code != *color.Code+*expected.Code {
			t.Errorf("TestBackgroundOnColor failed for color %v: got %v, expected %v", label, actual, expected)
		}
	}
}

func TestBackgroundRGBOnColor(t *testing.T) {

	expected := "\x1b[38;2;12;12;12m\x1b[48;2;255;255;255m"
	actual := NewColorRGB(12, 12, 12).BackgroundRGB(255, 255, 255).toCode()

	if expected != actual {
		t.Errorf(errorFmt, "TestBackgroundRGBOnColor", expected, actual)
	}
}

func TestDetectColor(t *testing.T) {
	empty := Color{}
	if actual := detectColor(empty); actual != empty {
		t.Errorf("TestDetectColor failed for empty Color: got %v, expected %v", actual, empty)
	}

	pointerOfString := Cyan
	expectedColor := NewColor(Cyan)
	if actual := detectColor(pointerOfString); actual.toCode() != expectedColor.toCode() {
		t.Errorf("TestDetectColor failed for pointer of a string: got %v, expected %v", actual, expectedColor)
	}

	inputColor := NewColor(Cyan)
	expectedColor = NewColor(Cyan)
	if actual := detectColor(inputColor); actual.toCode() != expectedColor.toCode() {
		t.Errorf("TestDetectColor failed for non emptu Color object: got %v, expected %v", actual.toCode(), expectedColor.toCode())
	}

	wrongType := "string"
	if actual := detectColor(wrongType); actual != empty {
		t.Errorf("TestDetectColor failed for invalidType: got %v, expected %v", actual, empty)
	}

	wrongContentPointer := &wrongType
	if actual := detectColor(wrongContentPointer); actual.toCode() != colorReset {
		t.Errorf("TestDetectColor failed for wrong color as type pointer of a string: got %v, expected %v", actual, colorReset)
	}
}
