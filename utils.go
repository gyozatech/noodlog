package noodlog

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
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

// IsEmpty function checks for the input if it is empty / vice versa
func IsEmpty(i interface{}) bool {
	switch v := i.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return true
		}
	default:
		if v == nil {
			return true
		}
	}
	return false

}

// getTime function retreives the time in the custom time format or default if not defined
func getTime(t EncodeTime) (now string) {
	isEmptyZone := IsEmpty(t.TimeZone)
	isEmptyFormat := IsEmpty(t.Format)
	now = strings.Split(time.Now().String(), "m")[0]
	if isEmptyZone && isEmptyFormat {
		return
	}

	if !isEmptyZone {
		tz, err := time.LoadLocation(t.TimeZone)
		if err != nil {
			return
		}
		if isEmptyFormat {
			now = time.Now().In(tz).String()
			return
		}
		now = time.Now().In(tz).Format(t.Format)
	} else {
		if isEmptyFormat {
			return
		}
		now = time.Now().Format(t.Format)
	}

	return
}
