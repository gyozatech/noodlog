package noodlog

import (
	"testing"
)

func TestIsValidTrueColor(t *testing.T) {
	validColor := 255
	negativeColor := -1
	overflowColor := 256

	if actual := IsValidTrueColor(validColor); actual != true {
		t.Errorf("TestIsValidTrueColor failed: expected true found %t", actual)
	}

	if actual := IsValidTrueColor(negativeColor); actual != false {
		t.Errorf("TestIsValidTrueColor failed: expected false found %t", actual)
	}

	if actual := IsValidTrueColor(overflowColor); actual != false {
		t.Errorf("TestIsValidTrueColor failed: expected false found %t", actual)
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
		redColor:       Color{Code: &redCode},
		greenColor:     Color{Code: &greenCode},
		yellowColor:    Color{Code: &yellowCode},
		blueColor:      Color{Code: &blueCode},
		purpleColor:    Color{Code: &purpleCode},
		cyanColor:      Color{Code: &cyanCode},
		"invalidColor": Color{Code: &resetCode},
	}

	for label, expected := range testData {
		if actual := NewColor(&label); *actual.Code != *expected.Code {
			t.Errorf("TestNewColor failed for color %v: got %v, expected %v", label, actual, expected)
		}
	}
}

func TestNewColorRGB(t *testing.T) {
	var (
		redCode     = "\033[38;2;255;0;0m"
		greenCode   = "\033[38;2;0;255;0m"
		blueCode    = "\033[38;2;0;0;255m"
		mixedCode   = "\033[38;2;128;128;128m"
		invalidCode = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       Color{Code: &redCode},
		greenColor:     Color{Code: &greenCode},
		blueColor:      Color{Code: &blueCode},
		"mixedColor":   Color{Code: &mixedCode},
		"invalidColor": Color{Code: &invalidCode},
	}

	if actual := NewColorRGB(255, 0, 0); *actual.Code != *testData[redColor].Code {
		t.Errorf("TestNewColorRGB failed for color %v: got %v, expected %v", redColor, actual, testData[redColor])
	}

	if actual := NewColorRGB(0, 255, 0); *actual.Code != *testData[greenColor].Code {
		t.Errorf("TestNewColorRGB failed for color %v: got %v, expected %v", greenColor, actual, testData[greenColor])
	}

	if actual := NewColorRGB(0, 0, 255); *actual.Code != *testData[blueColor].Code {
		t.Errorf("TestNewColorRGB failed for color %v: got %v, expected %v", blueColor, actual, testData[blueColor])
	}

	if actual := NewColorRGB(128, 128, 128); *actual.Code != *testData["mixedColor"].Code {
		t.Errorf("TestNewColorRGB failed for color %v: got %v, expected %v", "mixedColor", actual, testData["mixedColor"])
	}

	if actual := NewColorRGB(256, 0, 0); *actual.Code != *testData["invalidColor"].Code {
		t.Errorf("TestNewColorRGB failed for color %v: got %v, expected %v", "invalidColor", actual, testData["invalidColor"])
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
		redColor:       Color{Code: &redCode},
		greenColor:     Color{Code: &greenCode},
		yellowColor:    Color{Code: &yellowCode},
		blueColor:      Color{Code: &blueCode},
		purpleColor:    Color{Code: &purpleCode},
		cyanColor:      Color{Code: &cyanCode},
		"invalidColor": Color{Code: &resetCode},
	}

	for label, expected := range testData {
		if actual := Background(&label); *actual.Code != *expected.Code {
			t.Errorf("TestBackground failed for color %v: got %v, expected %v", label, actual, expected)
		}
	}
}

func TestBackgroundRGB(t *testing.T) {
	var (
		redCode     = "\033[48;2;255;0;0m"
		greenCode   = "\033[48;2;0;255;0m"
		blueCode    = "\033[48;2;0;0;255m"
		mixedCode   = "\033[48;2;128;128;128m"
		invalidCode = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       Color{Code: &redCode},
		greenColor:     Color{Code: &greenCode},
		blueColor:      Color{Code: &blueCode},
		"mixedColor":   Color{Code: &mixedCode},
		"invalidColor": Color{Code: &invalidCode},
	}

	if actual := BackgroundRGB(255, 0, 0); *actual.Code != *testData[redColor].Code {
		t.Errorf("TestBackgroundRGB failed for color %v: got %v, expected %v", redColor, actual, testData[redColor])
	}

	if actual := BackgroundRGB(0, 255, 0); *actual.Code != *testData[greenColor].Code {
		t.Errorf("TestBackgroundRGB failed for color %v: got %v, expected %v", greenColor, actual, testData[greenColor])
	}

	if actual := BackgroundRGB(0, 0, 255); *actual.Code != *testData[blueColor].Code {
		t.Errorf("TestBackgroundRGB failed for color %v: got %v, expected %v", blueColor, actual, testData[blueColor])
	}

	if actual := BackgroundRGB(128, 128, 128); *actual.Code != *testData["mixedColor"].Code {
		t.Errorf("TestBackgroundRGB failed for color %v: got %v, expected %v", "mixedColor", actual, testData["mixedColor"])
	}

	if actual := BackgroundRGB(256, 0, 0); *actual.Code != *testData["invalidColor"].Code {
		t.Errorf("TestBackgroundRGB failed for color %v: got %v, expected %v", "invalidColor", actual, testData["invalidColor"])
	}
}

func TestBackgroundOnColor(t *testing.T) {
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
		redColor:       Color{Code: &redCode},
		greenColor:     Color{Code: &greenCode},
		yellowColor:    Color{Code: &yellowCode},
		blueColor:      Color{Code: &blueCode},
		purpleColor:    Color{Code: &purpleCode},
		cyanColor:      Color{Code: &cyanCode},
		"invalidColor": Color{Code: &resetCode},
	}

	var existingCode = "code"
	color := Color{Code: &existingCode}

	for label, expected := range testData {
		if actual := color.Background(&label); *actual.Code != *color.Code+*expected.Code {
			t.Errorf("TestBackgroundOnColor failed for color %v: got %v, expected %v", label, actual, expected)
		}
	}
}

func TestBackgroundRGBOnColor(t *testing.T) {
	var (
		redCode     = "\033[48;2;255;0;0m"
		greenCode   = "\033[48;2;0;255;0m"
		blueCode    = "\033[48;2;0;0;255m"
		mixedCode   = "\033[48;2;128;128;128m"
		invalidCode = "\033[0m"
	)

	var testData = map[string]Color{
		redColor:       Color{Code: &redCode},
		greenColor:     Color{Code: &greenCode},
		blueColor:      Color{Code: &blueCode},
		"mixedColor":   Color{Code: &mixedCode},
		"invalidColor": Color{Code: &invalidCode},
	}

	var existingCode = "code"

	if actual := (Color{Code: &existingCode}).BackgroundRGB(255, 0, 0); *actual.Code != existingCode+*testData[redColor].Code {
		t.Errorf("TestBackgroundRGBOnColor failed for color %v: got %v, expected %v", redColor, actual, testData[redColor])
	}

	if actual := (Color{Code: &existingCode}).BackgroundRGB(0, 255, 0); *actual.Code != existingCode+*testData[greenColor].Code {
		t.Errorf("TestBackgroundRGBOnColor failed for color %v: got %v, expected %v", greenColor, actual, testData[greenColor])
	}

	if actual := (Color{Code: &existingCode}).BackgroundRGB(0, 0, 255); *actual.Code != existingCode+*testData[blueColor].Code {
		t.Errorf("TestBackgroundRGBOnColor failed for color %v: got %v, expected %v", blueColor, actual, testData[blueColor])
	}

	if actual := (Color{Code: &existingCode}).BackgroundRGB(128, 128, 128); *actual.Code != existingCode+*testData["mixedColor"].Code {
		t.Errorf("TestBackgroundRGBOnColor failed for color %v: got %v, expected %v", "mixedColor", actual, testData["mixedColor"])
	}

	if actual := (Color{Code: &existingCode}).BackgroundRGB(256, 0, 0); *actual.Code != existingCode+*testData["invalidColor"].Code {
		t.Errorf("TestBackgroundRGBOnColor failed for color %v: got %v, expected %v", "invalidColor", actual, testData["invalidColor"])
	}
}

func TestToCode(t *testing.T) {
	empty := Color{}
	if actual := empty.ToCode(); actual != colorReset {
		t.Errorf("TestToCode failed for empty Color: got %s, expected %s", actual, colorReset)
	}

	existingCode := "code"
	existing := Color{Code: &existingCode}

	if actual := existing.ToCode(); actual != existingCode {
		t.Errorf("TestToCode failed for existing Color: got %s, expected %s", actual, existingCode)
	}
}

func TestEnableDisableColors(t *testing.T) {
	//TODO
}

func TestSetTraceColor(t *testing.T) {
	//TODO
}

func TestSetDebugColor(t *testing.T) {
	//TODO
}

func TestSetInfoColor(t *testing.T) {
	//TODO
}

func TestSetWarnColor(t *testing.T) {
	//TODO
}

func TestSetErrorColor(t *testing.T) {
	//TODO
}

func TestDetectColor(t *testing.T) {
	empty := Color{}
	if actual := DetectColor(empty); actual != empty {
		t.Errorf("TestDetectColor failed for empty Color: got %v, expected %v", actual, empty)
	}

	pointerOfString := Cyan
	expectedColor := NewColor(Cyan)
	if actual := DetectColor(pointerOfString); actual.ToCode() != expectedColor.ToCode() {
		t.Errorf("TestDetectColor failed for pointer of a string: got %v, expected %v", actual, expectedColor)
	}

	wrongType := "string"
	if actual := DetectColor(wrongType); actual != empty {
		t.Errorf("TestDetectColor failed for invalidType: got %v, expected %v", actual, empty)
	}

	wrongContentPointer := &wrongType
	if actual := DetectColor(wrongContentPointer); actual.ToCode() != colorReset {
		t.Errorf("TestDetectColor failed for wrong color as type pointer of a string: got %v, expected %v", actual, colorReset)
	}
}
