package noodlog

import "testing"

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
