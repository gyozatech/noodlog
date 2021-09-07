package noodlog

import (
	"fmt"
	"strings"
	"testing"
)

var jsnSlashes = "{\"param1\": 1, \"param2\": \"secret\", \"param3\": {\"param4\": \"secret\"}}"
var jsnRaw = `{"param1": 1, "param2": "secret", "param3": {"param4": "secret"}}`
var jsnMap = map[string]interface{}{
	"param1": 1,
	"param2": "secret",
	"param3": map[string]interface{}{"param4": "secret"},
}

func TestPointerOfString(t *testing.T) {
	actual := pointerOfString("red")
	if *actual != *Red {
		t.Errorf("TestPointerOfString failed: expected %s, got %s", *Red, *actual)
	}
}

func TestPointerOfBool(t *testing.T) {
	actual := pointerOfBool(true)
	if *actual != *Enable {
		t.Errorf("TestPointerOfString failed: expected %t, got %t", *Enable, *actual)
	}
}

func TestStringify(t *testing.T) {
	expected := "hello to everyone!"
	actual := stringify([]interface{}{"hello", "to", "everyone!"})
	if actual != expected {
		t.Errorf("TestStringify failed: expected %s, got %s", expected, actual)
	}
}

func TestObscureParam(t *testing.T) {
	testData := map[string]string{
		jsnSlashes: "{\"param1\": 1, \"param2\": \"**********\", \"param3\": {\"param4\": \"secret\"}}",
		jsnRaw:     `{"param1": 1, "param2": "**********", "param3": {"param4": "secret"}}`,
	}
	for input, expected := range testData {
		if actual := obscureParam(input, "param2"); actual != expected {
			t.Errorf("TestObscureParam failed: expected %s, got %s", expected, actual)
		}
	}
}

func TestObscureParams(t *testing.T) {
	testData := map[string]string{
		jsnSlashes: "{\"param1\": 1, \"param2\": \"**********\", \"param3\": {\"param4\": \"**********\"}}",
		jsnRaw:     `{"param1": 1, "param2": "**********", "param3": {"param4": "**********"}}`,
	}
	for input, expected := range testData {
		if actual := obscureParams(input, []string{"param2", "param4"}); actual != expected {
			t.Errorf("TestObscureParams failed: expected %s, got %s", expected, actual)
		}
	}
}

func TestStrToObj(t *testing.T) {
	simpleMessageStr := "simple-message"
	var simpleMessage interface{} = simpleMessageStr
	testData := map[string]interface{}{
		jsnSlashes:       jsnMap,
		jsnRaw:           jsnMap,
		simpleMessageStr: simpleMessage,
	}
	for input, expected := range testData {
		if actual := strToObj(input); fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
			t.Errorf("TestStrToObj failed: expected %s, got %s", expected, actual)
		}
	}
}

func TestTraceCaller(t *testing.T) {

	errFormat := "TestTraceCaller failed: expected %s, got %s"

	file, function := traceCaller(5)
	expectedFile := ":0"
	expectedFunction := ""

	if file != expectedFile {
		t.Errorf(errFormat, expectedFile, file)
	}
	if function != expectedFunction {
		t.Errorf(errFormat, expectedFunction, function)
	}

	expectedFilePortion := "runtime"
	expectedFunction = "runtime.goexit"
	file, function = caller2()

	if !strings.Contains(file, expectedFilePortion) {
		t.Errorf(errFormat, expectedFilePortion, file)
	}
	if function != expectedFunction {
		t.Errorf(errFormat, expectedFunction, function)
	}
}

func caller1() (string, string) {
	return traceCaller(6)
}

func caller2() (string, string) {
	return caller1()
}
