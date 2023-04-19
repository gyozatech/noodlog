package noodlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

var errorFmt string = "%s failed: expected %v, got %v"

var defaultLogger = Logger{
	level:                infoLevel,
	logWriter:            os.Stdout,
	prettyPrint:          false,
	traceCaller:          false,
	traceCallerLevel:     5,
	obscureSensitiveData: false,
	sensitiveParams:      nil,
	colors:               false,
	colorMap:             colorMap,
}

var customLogger = Logger{
	level:                errorLevel,
	logWriter:            os.Stderr,
	prettyPrint:          true,
	traceCaller:          true,
	traceCallerLevel:     6,
	obscureSensitiveData: true,
	sensitiveParams:      []string{"password"},
	colors:               true,
	colorMap:             colorMap,
}

func toStr(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

func wildCardToRegexp(pattern string) string {
	var result strings.Builder
	for i, literal := range strings.Split(pattern, "*") {
		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}
		// Quote any regular expression meta characters in the literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return result.String()
}

func Matches(value string, pattern string) bool {
	result, _ := regexp.MatchString(wildCardToRegexp(pattern), value)
	return result
}

func TestNewLogger(t *testing.T) {

	expected := toStr(defaultLogger)
	actual := toStr(*NewLogger())

	if actual != expected {
		t.Errorf(errorFmt, "TestNewLogger", expected, actual)
	}
}

func TestSetConfigsEmptyConfigs(t *testing.T) {
	expected := toStr(defaultLogger)
	actual := toStr(*NewLogger().SetConfigs(Configs{}))

	if actual != expected {
		t.Errorf(errorFmt, "TestSetConfigsEmptyConfigs", expected, actual)
	}
}

func TestSetConfigsFullConfigsAllEnabled(t *testing.T) {
	expected := toStr(customLogger)
	actual := toStr(*NewLogger().SetConfigs(Configs{
		LogLevel:             LevelError,
		LogWriter:            os.Stderr,
		JSONPrettyPrint:      Enable,
		TraceCaller:          Enable,
		SinglePointTracing:   Enable,
		Colors:               Enable,
		CustomColors:         &CustomColors{Trace: Color{}, Debug: Green},
		ObscureSensitiveData: Enable,
		SensitiveParams:      []string{"password"},
	}))

	if actual != expected {
		t.Errorf(errorFmt, "TestSetConfigsFullConfigsAllEnabled", expected, actual)
	}
}

func TestSetConfigsFullConfigsAllDisabled(t *testing.T) {
	customLogger.obscureSensitiveData = false
	customLogger.prettyPrint = false
	customLogger.colors = false
	customLogger.traceCaller = false
	customLogger.traceCallerLevel = 5
	customLogger.obscureSensitiveData = false

	expected := toStr(customLogger)
	actual := toStr(*NewLogger().SetConfigs(Configs{
		LogLevel:             LevelError,
		LogWriter:            os.Stderr,
		JSONPrettyPrint:      Disable,
		TraceCaller:          Disable,
		SinglePointTracing:   Disable,
		Colors:               Disable,
		ObscureSensitiveData: Disable,
		SensitiveParams:      []string{"password"},
	}))

	if actual != expected {
		t.Errorf(errorFmt, "TestSetConfigsFullConfigsAllDisabled", expected, actual)
	}
}

func TestLevel(t *testing.T) {
	l := NewLogger()
	testMap := map[string]int{
		traceLabel: traceLevel,
		infoLabel:  infoLevel,
		warnLabel:  warnLevel,
		errorLabel: errorLevel,
		panicLabel: panicLevel,
		fatalLabel: fatalLevel,
	}

	for input, expected := range testMap {
		if l.Level(input).level != expected {
			t.Errorf(errorFmt, "TestLevel", expected, l.level)
		}
	}
}

func TestLogWriter(t *testing.T) {
	l := NewLogger()
	testSlice := []io.Writer{
		os.Stdin,
		os.Stdout,
		os.Stderr,
	}

	for _, expected := range testSlice {
		if l.LogWriter(expected).logWriter != expected {
			t.Errorf(errorFmt, "TestLogWriter", expected, l.logWriter)
		}
	}
}

func TestEnableDisableJSONPrettyPrint(t *testing.T) {
	l := NewLogger()
	errFormat := "TestEnableDisableJSONPrettyPrint failed: expected prettyPrint %t, got %t"

	if !l.EnableJSONPrettyPrint().prettyPrint {
		t.Errorf(errFormat, true, l.prettyPrint)
	}
	if l.DisableJSONPrettyPrint().prettyPrint {
		t.Errorf(errFormat, false, l.prettyPrint)
	}
}

func TestEnableDisableTraceCaller(t *testing.T) {
	l := NewLogger()
	errFormat := "TestEnableDisableTraceCaller failed: expected traceCaller %t, got %t"

	if !l.EnableTraceCaller().traceCaller {
		t.Errorf(errFormat, true, l.traceCaller)
	}
	if l.DisableTraceCaller().traceCaller {
		t.Errorf(errFormat, false, l.traceCaller)
	}
}

func TestEnableDisableSinglePointTracing(t *testing.T) {
	l := NewLogger()
	errFormat := "TestEnableDisableSinglePointTracing failed: expected traceCallerLevel %d, got %d"

	if l.EnableSinglePointTracing().traceCallerLevel != 6 {
		t.Errorf(errFormat, 6, l.traceCallerLevel)
	}
	if l.DisableSinglePointTracing().traceCallerLevel != 5 {
		t.Errorf(errFormat, 5, l.traceCallerLevel)
	}
}

func TestEnableDisableLoggerColors(t *testing.T) {
	l := NewLogger()
	errFormat := "TestEnableDisableLoggerColors failed: expected colors %t, got %t"

	if !l.EnableColors().colors {
		t.Errorf(errFormat, true, l.colors)
	}
	if l.DisableColors().colors {
		t.Errorf(errFormat, false, l.colors)
	}
}

func TestEnableDisableObscureSensitiveParams(t *testing.T) {

	params := []string{"password", "age"}
	errFormat := "TestEnableDisableObscureSensitiveParams failed: expected obscureSensitiveData %t, got %t"

	l := NewLogger()
	l.EnableObscureSensitiveData(params)

	if !l.obscureSensitiveData {
		t.Errorf(errFormat, true, l.obscureSensitiveData)
	}
	if !reflect.DeepEqual(l.sensitiveParams, params) {
		t.Errorf("TestEnableDisableObscureSensitiveParams failed: expected sensitiveParams %v. got %v", params, l.sensitiveParams)
	}

	l.DisableObscureSensitiveData()
	if l.obscureSensitiveData {
		t.Errorf(errFormat, false, l.obscureSensitiveData)
	}
}

func TestSetSensitiveParams(t *testing.T) {

	params := []string{"secret", "privatekey"}
	errFormat := "TestEnableDisableObscureSensitiveParams failed: expected sensitiveParams %v. got %v"

	l := NewLogger()
	l.EnableObscureSensitiveData(params)

	if !reflect.DeepEqual(l.SetSensitiveParams(params).sensitiveParams, params) {
		t.Errorf(errFormat, params, l.sensitiveParams)
	}
}

type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestInfoLogging(t *testing.T) {

	var testLoggingMap = map[interface{}]string{
		"":      `"level":"%s","message":""`,
		"hello": `"level":"%s","message":"hello"`,
		`{"name": "gyoza", "cool": true, "password": "Sup3rS3cr3t"}`: `"level":"%s","message":{"cool":true,"name":"gyoza","password":"**********"}`,
		"{\"name\": \"gyozatech\", \"repo\": \"noodlog\"}":           `"level":"%s","message":{"name":"gyozatech","repo":"noodlog"}`,
		account{"gyozatech", "Sup3rS3cr3t"}:                          `"level":"%s","message":{"password":"**********","username":"gyozatech"}`,
	}

	var b bytes.Buffer
	b.Reset()
	l := NewLogger().EnableObscureSensitiveData([]string{"password"}).LogWriter(&b)

	for input, expectedFmt := range testLoggingMap {
		if input != "" {
			l.Info(input)
		} else {
			l.Info()
		}

		actual := b.String()
		expected := fmt.Sprintf(expectedFmt, "info")
		if !strings.Contains(actual, expected) {
			t.Errorf(errorFmt, "TestInfoLogging", expected, actual)
		}
		b.Reset()
	}
}

func TestLogging(t *testing.T) {

	var b bytes.Buffer
	b.Reset()
	l := NewLogger().
		EnableObscureSensitiveData([]string{"password"}).
		EnableTraceCaller().
		EnableColors().
		LogWriter(&b)

	for _, level := range []string{"trace", "debug", "info", "warn", "error"} {

		input1 := "logging example"
		input2 := "text"

		input3 := "logging example %s"
		input4 := "logging example text"

		expected := `{"level":"*","message":"logging example text","time":"*"}`

		switch level {
		case "trace":
			l.Trace(input1, input2)
		case "debug":
			l.Debug(input1, input2)
		case "info":
			l.Info(input1, input2)
		case "warn":
			l.DisableObscureSensitiveData()
			l.Warn(input4)
		case "error":
			l.Error(input3, input2)
		}

		actual := b.String()

		if level == "debug" || level == "trace" {
			if actual != "" {
				t.Errorf(errorFmt, "TestLogging", "", actual)
			}
		} else {
			if !Matches(actual, expected) {
				t.Errorf(errorFmt, "TestLogging", expected, actual)
			}
		}

		b.Reset()
	}
}

func TestLogPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(errorFmt, "TestLogPanic", "panic", "not-panic")
		}
	}()

	l := NewLogger()
	l.Panic("Hello, I'm gonna panic!")
}

func TestLogFatal(t *testing.T) {

	var b bytes.Buffer
	b.Reset()
	l := NewLogger().EnableJSONPrettyPrint().LogWriter(&b)

	os.Setenv("EXIT_ON_FATAL_DISABLED", "true")
	l.Fatal("Hello, I'm gonna exit!")
	actual := b.String()
	expected := `{
		"level": "fatal",
		"message": "Hello, I'm gonna exit!",
		"time": "*"
	 }`
	b.Reset()
	if Matches(actual, expected) {
		t.Errorf(errorFmt, "TestLogFatal", expected, actual)
	}
}

func composeColor(color string) string {
	return "\033[" + color + "m"
}

func TestSetCustomColors(t *testing.T) {
	errFormat := "TestSetCustomColors failed: expected %s got %s"

	l := NewLogger()

	l.SetCustomColors(CustomColors{
		Trace: Blue,
		Debug: Purple,
		Info:  Yellow,
		Warn:  Green,
		Error: Default,
	})

	if blueCode := composeColor(colorBlue); l.colorMap[traceLabel] != blueCode {
		t.Errorf(errFormat, blueCode, l.colorMap[traceLabel])
	}
	if purpleCode := composeColor(colorPurple); l.colorMap[debugLabel] != purpleCode {
		t.Errorf(errFormat, purpleCode, l.colorMap[debugLabel])
	}
	if yellowCode := composeColor(colorYellow); l.colorMap[infoLabel] != yellowCode {
		t.Errorf(errFormat, yellowCode, l.colorMap[infoLabel])
	}
	if greenCode := composeColor(colorGreen); l.colorMap[warnLabel] != greenCode {
		t.Errorf(errFormat, greenCode, l.colorMap[warnLabel])
	}
	if defaultCode := composeColor(colorReset); l.colorMap[errorLabel] != defaultCode {
		t.Errorf(errFormat, defaultCode, l.colorMap[errorLabel])
	}
}

var colorTestMap = map[Color]string{
	NewColor(Blue):   composeColor(colorBlue),
	NewColor(Purple): composeColor(colorPurple),
}

func TestSetTraceColor(t *testing.T) {
	l := NewLogger()

	for color, colorCode := range colorTestMap {
		l.SetTraceColor(color)
		if actualCode := l.colorMap[traceLabel]; actualCode != colorCode {
			t.Errorf(errorFmt, "TestSetTraceColor", colorCode, actualCode)
		}
	}
}

func TestSetDebugColor(t *testing.T) {
	l := NewLogger()

	for color, colorCode := range colorTestMap {
		l.SetDebugColor(color)
		if actualCode := l.colorMap[debugLabel]; actualCode != colorCode {
			t.Errorf(errorFmt, "TestSetDebugColor", colorCode, actualCode)
		}
	}
}

func TestSetInfoColor(t *testing.T) {

	l := NewLogger()

	for color, colorCode := range colorTestMap {
		l.SetInfoColor(color)
		if actualCode := l.colorMap[infoLabel]; actualCode != colorCode {
			t.Errorf(errorFmt, "TestSetInfoColor", colorCode, actualCode)
		}
	}
}

func TestSetWarnColor(t *testing.T) {
	l := NewLogger()

	for color, colorCode := range colorTestMap {
		l.SetWarnColor(color)
		if actualCode := l.colorMap[warnLabel]; actualCode != colorCode {
			t.Errorf(errorFmt, "TestSetWarnColor", colorCode, actualCode)
		}
	}
}

func TestSetErrorColor(t *testing.T) {
	l := NewLogger()

	for color, colorCode := range colorTestMap {
		l.SetErrorColor(color)
		if actualCode := l.colorMap[errorLabel]; actualCode != colorCode {
			t.Errorf(errorFmt, "TestSetErrorColor", colorCode, actualCode)
		}
	}
}

func TestAdaptMessage(t *testing.T) {

	testMap := map[interface{}]string{
		struct{ test string }{"Hello test"}:  "{Hello test}",
		"Hi message":                         "Hi message",
		`{"name": "John", "surname": "Doe"}`: `map[name:John surname:Doe]`,
		fmt.Errorf("Nice error!"):            "Nice error!",
	}

	var b bytes.Buffer
	b.Reset()
	l := NewLogger().LogWriter(&b)

	for input, expected := range testMap {
		actual := l.adaptMessage(input)
		if toStr(actual) != expected {
			t.Errorf(errorFmt, "TestAdaptMessage", expected, actual)
		}
	}

}
