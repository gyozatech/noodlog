package noodlog

import (
	"testing"
)

/*func TestSetConfigsEmptyObject(t *testing.T) {

	errFormat := "TestSetConfigsEmptyObject failed: param %s expected %v, got %v"

	SetConfigs(Configs{})

	if logLevel != 3 {
		t.Errorf(errFormat, "logLevel", 5, logLevel)
	}
	if JSONPrettyPrint {
		t.Errorf(errFormat, "JSONPrettyPrint", false, JSONPrettyPrint)
	}
	if traceCallerEnabled {
		t.Errorf(errFormat, "traceCallerEnabled", false, traceCallerEnabled)
	}
	if traceCallerLevel != 5 {
		t.Errorf(errFormat, "traceCallerLevel", 5, traceCallerLevel)
	}
	if colorEnabled {
		t.Errorf(errFormat, "colorEnabled", false, colorEnabled)
	}
	if colorMap[traceLabel] != colorReset {
		t.Errorf(errFormat, "colorMap[traceLabel]", colorReset, colorMap[traceLabel])
	}
	if colorMap[debugLabel] != colorGreen {
		t.Errorf(errFormat, "colorMap[debugLabel]", colorGreen, colorMap[debugLabel])
	}
	if colorMap[infoLabel] != colorReset {
		t.Errorf(errFormat, "colorMap[infoLabel]", colorReset, colorMap[infoLabel])
	}
	if colorMap[warnLabel] != colorYellow {
		t.Errorf(errFormat, "colorMap[warnLabel]", colorYellow, colorMap[warnLabel])
	}
	if colorMap[errorLabel] != colorRed {
		t.Errorf(errFormat, "colorMap[errorLabel]", colorRed, colorMap[errorLabel])
	}
	if obscureSensitiveDataEnabled {
		t.Errorf(errFormat, "obscureSensitiveDataEnabled", false, obscureSensitiveDataEnabled)
	}
	if len(sensitiveParams) != 0 {
		t.Errorf(errFormat, "sensitiveParams", 0, len(sensitiveParams))
	}
}*/

func TestSetConfigsFullObject(t *testing.T) {

	sensitiveList := []string{"password", "age"}
	errFormat := "TestSetConfigsFullObject failed: param %s expected %v, got %v"

	SetConfigs(Configs{
		LogLevel:           LevelError,
		JSONPrettyPrint:    Enable,
		TraceCaller:        Enable,
		SinglePointTracing: Enable,
		Colors:             Enable,
		CustomColors: &CustomColors{
			Trace: Purple,
			Debug: Yellow,
			Info:  Red,
			Warn:  Blue,
			Error: Cyan,
		},
		ObscureSensitiveData: Enable,
		SensitiveParams:      sensitiveList,
	})

	if logLevel != 5 {
		t.Errorf(errFormat, "logLevel", 5, logLevel)
	}
	if !JSONPrettyPrint {
		t.Errorf(errFormat, "JSONPrettyPrint", true, JSONPrettyPrint)
	}
	if !traceCallerEnabled {
		t.Errorf(errFormat, "traceCallerEnabled", true, traceCallerEnabled)
	}
	if traceCallerLevel != 6 {
		t.Errorf(errFormat, "traceCallerLevel", 6, traceCallerLevel)
	}
	if !colorEnabled {
		t.Errorf(errFormat, "colorEnabled", true, colorEnabled)
	}
	if colorMap[traceLabel] != colorPurple {
		t.Errorf(errFormat, "colorMap[traceLabel]", colorPurple, colorMap[traceLabel])
	}
	if colorMap[debugLabel] != colorYellow {
		t.Errorf(errFormat, "colorMap[debugLabel]", colorYellow, colorMap[debugLabel])
	}
	if colorMap[infoLabel] != colorRed {
		t.Errorf(errFormat, "colorMap[infoLabel]", colorRed, colorMap[infoLabel])
	}
	if colorMap[warnLabel] != colorBlue {
		t.Errorf(errFormat, "colorMap[warnLabel]", colorBlue, colorMap[warnLabel])
	}
	if colorMap[errorLabel] != colorCyan {
		t.Errorf(errFormat, "colorMap[errorLabel]", colorCyan, colorMap[errorLabel])
	}
	if !obscureSensitiveDataEnabled {
		t.Errorf(errFormat, "obscureSensitiveDataEnabled", true, obscureSensitiveDataEnabled)
	}
	if len(sensitiveParams) != 2 {
		t.Errorf(errFormat, "sensitiveParams", 2, len(sensitiveParams))
	}
}

func TestLogLevel(t *testing.T) {
	testMap := map[string]int{
		"trace":       1,
		"debug":       2,
		"info":        3,
		"warn":        4,
		"error":       5,
		"invalidName": 3,
	}

	for label, level := range testMap {
		LogLevel(label)
		if logLevel != level {
			t.Errorf("TestLogLevel failed: expected %d, got %d", level, logLevel)
		}
	}
}

func TestEnableDisableJSONPrettyPrint(t *testing.T) {
	errFormat := "TestEnableDisableJSONPrettyPrint failed: JSONPrettyPrint expected %t, got %t "
	EnableJSONPrettyPrint()
	if !JSONPrettyPrint {
		t.Errorf(errFormat, true, JSONPrettyPrint)
	}
	DisableJSONPrettyPrint()
	if JSONPrettyPrint {
		t.Errorf(errFormat, false, JSONPrettyPrint)
	}
}

//TestEnableDisableObscureSensitiveData

//TestSetSensitiveParams

//TestSetComposeLog

//TestSetComposeMessage

//TestStringify

//TestAdaptMessage

//TestStrToObj

//TestObscureSensitiveData

//TestObscureParam
