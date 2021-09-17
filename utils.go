package noodlog

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

func pointerOfString(v string) *string {
	return &v
}

func pointerOfBool(v bool) *bool {
	return &v
}

func stringify(message []interface{}) string {
	var b strings.Builder
	for _, m := range message {
		if m != nil {
			fmt.Fprintf(&b, "%v ", m)
		}
	}
	msg := b.String()
	return msg[:len(msg)-1]
}

func obscureParams(jsn string, sensitiveParams []string) string {
	for _, param := range sensitiveParams {
		jsn = obscureParam(jsn, param)
	}
	return jsn
}

func obscureParam(jsn string, param string) string {
	rWithSlash := *regexp.MustCompile(`\\"` + param + `\\":.*?"(.*?)\\"`)
	jsn = rWithSlash.ReplaceAllString(jsn, `\"`+param+`\": \"**********\"`)

	rWithoutSlash := *regexp.MustCompile(`"` + param + `":.*?"(.*?)"`)
	return rWithoutSlash.ReplaceAllString(jsn, `"`+param+`": "**********"`)
}

func strToObj(strMsg string) interface{} {
	if byteMsg := []byte(strMsg); json.Valid(byteMsg) {
		var obj interface{}
		_ = json.Unmarshal(byteMsg, &obj)
		return obj
	}
	return strMsg
}

// traceCaller function retrieves the filename of the function which wants to log
func traceCaller(traceCallerLevel int) (file, function string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(traceCallerLevel, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return fmt.Sprintf("%s:%d", frame.File, frame.Line),
		frame.Function
}
