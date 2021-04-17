package noodlog

import (
	"testing"
)

func TestGetColorByName(t *testing.T) {
	testData := map[string]string{
		defaultColor:   colorReset,
		redColor:       colorRed,
		greenColor:     colorGreen,
		yellowColor:    colorYellow,
		blueColor:      colorBlue,
		purpleColor:    colorPurple,
		cyanColor:      colorCyan,
		"invalidcolor": colorReset,
	}

	for label, expected := range testData {
		if actual := getColorByName(label); actual != expected {
			t.Errorf("TestGetColorByName failed for color %s: got %s, expected %s", label, actual, expected)
		}
	}
}

func TestSetCustomColors(t *testing.T) {
	errFormat := "TestSetCustomColors failed: expected %s found %s"

	setCustomColors(CustomColors{
		Trace: Blue,
		Debug: Purple,
		Info:  Yellow,
		Warn:  Green,
		Error: Default,
	})

	if colorMap[traceLabel] != colorBlue {
		t.Errorf(errFormat, blueColor, colorMap[traceLabel])
	}
	if colorMap[debugLabel] != colorPurple {
		t.Errorf(errFormat, purpleColor, colorMap[debugLabel])
	}
	if colorMap[infoLabel] != colorYellow {
		t.Errorf(errFormat, yellowColor, colorMap[infoLabel])
	}
	if colorMap[warnLabel] != colorGreen {
		t.Errorf(errFormat, greenColor, colorMap[warnLabel])
	}
	if colorMap[errorLabel] != colorReset {
		t.Errorf(errFormat, defaultColor, colorMap[errorLabel])
	}
}

func TestEnableDisableColors(t *testing.T) {
	errFormat := "TestEnableDisableColors failed: expected %v colorEnabled, got %v"
	if colorEnabled {
		t.Errorf(errFormat, false, true)
	}
	EnableColors()
	if !colorEnabled {
		t.Errorf(errFormat, true, false)
	}
}

func TestSetTraceColor(t *testing.T) {
	setColorAssertions("TestSetTraceColor", traceLabel, SetTraceColor, t)
}

func TestSetDebugColor(t *testing.T) {
	setColorAssertions("TestSetDebugColor", traceLabel, SetTraceColor, t)
}

func TestSetInfoColor(t *testing.T) {
	setColorAssertions("TestSetInfoColor", traceLabel, SetTraceColor, t)
}

func TestSetWarnColor(t *testing.T) {
	setColorAssertions("TestSetWarnColor", traceLabel, SetTraceColor, t)
}

func TestSetErrorColor(t *testing.T) {
	setColorAssertions("TestSetErrorColor", traceLabel, SetTraceColor, t)
}

func setColorAssertions(testName string, label string, setFunction func(color string), t *testing.T) {
	errFormat := "%s failed: expected %s, found %s"
	testMap := map[string]string{
		blueColor:   colorBlue,
		purpleColor: colorPurple,
	}
	for color, colorCode := range testMap {
		setFunction(color)
		if actualCode := colorMap[label]; actualCode != colorCode {
			t.Errorf(errFormat, testName, colorCode, actualCode)
		}
	}
}
