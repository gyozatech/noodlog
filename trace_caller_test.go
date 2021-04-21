package noodlog

import "testing"

func TestEnableDisableTraceCaller(t *testing.T) {
	errFormat := "TestEnableDisableTraceCaller failed: expected traceCallerEnabled %t, got %t"
	EnableTraceCaller()
	if !traceCallerEnabled {
		t.Errorf(errFormat, true, traceCallerEnabled)
	}
	DisableTraceCaller()
	if traceCallerEnabled {
		t.Errorf(errFormat, false, traceCallerEnabled)
	}
}

func TestEnableDisableSinglePointTracing(t *testing.T) {
	errFormat := "TestEnableDisableSinglePointTracing failed: expected traceCallerLevel %d, got %d"
	EnableSinglePointTracing()
	if traceCallerLevel != 6 {
		t.Errorf(errFormat, 6, traceCallerLevel)
	}
	DisableSinglePointTracing()
	if traceCallerLevel != 5 {
		t.Errorf(errFormat, 5, traceCallerLevel)
	}
}

func TestTraceCaller(t *testing.T) {

}

func TestTraceCallerSinglePointTracing(t *testing.T) {

}
