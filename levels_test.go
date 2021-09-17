package noodlog

import (
	"testing"
)

func TestGetLogLevel(t *testing.T) {

	logLevelTestMap := map[string]int{
		"non-existing-level": infoLevel,
		"trace":              traceLevel,
		"debug":              debugLevel,
		"info":               infoLevel,
		"warn":               warnLevel,
		"error":              errorLevel,
		"panic":              panicLevel,
		"fatal":              fatalLevel,
	}

	for input, expected := range logLevelTestMap {
		actual := getLogLevel(input)
		if actual != expected {
			t.Errorf("TestGetLogLevel failed: expected %d, got %d", expected, actual)
		}
	}

}
