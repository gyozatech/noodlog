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
	setCustomColors(CustomColors{
		Trace: Blue,
		Debug: Purple,
		Info:  Yellow,
		Warn:  Green,
		Error: Default,
	})

	if colorMap[traceLabel] != colorBlue {
		t.Errorf("TestSetCustomColors failed: expected blue found %s", colorMap[traceLabel])
	}
	if colorMap[debugLabel] != colorPurple {
		t.Errorf("TestSetCustomColors failed: expected purple found %s", colorMap[debugLabel])
	}
	if colorMap[infoLabel] != colorYellow {
		t.Errorf("TestSetCustomColors failed: expected yellow found %s", colorMap[infoLabel])
	}
	if colorMap[warnLabel] != colorGreen {
		t.Errorf("TestSetCustomColors failed: expected green found %s", colorMap[warnLabel])
	}
	if colorMap[errorLabel] != colorReset {
		t.Errorf("TestSetCustomColors failed: expected default found %s", colorMap[errorLabel])
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
