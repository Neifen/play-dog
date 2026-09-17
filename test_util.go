package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEq(t *testing.T, value, expected any) {
	if value != expected {
		t.Errorf("expected %+v but got %+v", expected, value)
	}
}

func assertEqMsg(t *testing.T, value, expected any, msg string, a ...any) {
	if value != expected {
		t.Errorf("expected '%+v' for %s but got '%+v'", expected, fmt.Sprintf(msg, a...), value)
	}
}

func assertNotEmpty(t *testing.T, value any, msg string, a ...any) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		if v.Len() == 0 {
			t.Errorf("expected %+v to not be empty for: %s", value, fmt.Sprintf(msg, a...))
		}
	default:
		t.Errorf("assertNotEmpty: unsupported type %T", value)
	}
}

func assertNeq(t *testing.T, value, expected any) {
	if value == expected {
		t.Errorf("expected %+v to be different from %+v", expected, value)
	}
}

func assertNil(t *testing.T, value any, desc string, a ...any) {
	if value == nil {
		return //success
	}
	v := reflect.ValueOf(value)
	if !v.IsNil() {
		t.Errorf("expected value for %s to be nil but was: %+v", fmt.Sprintf(desc, a...), value)
	}
}

func assertNotNil(t *testing.T, value any, desc string, a ...any) {
	if value != nil {
		return //success
	}
	v := reflect.ValueOf(value)
	if v.IsNil() {
		t.Errorf("expected value for %s to be not be nil", fmt.Sprintf(desc, a...))
	}
}

func assertErr(t *testing.T, err error, desc string, a ...any) {
	if err == nil {
		t.Errorf("expected error for %s but there wasn't any", fmt.Sprintf(desc, a...))
	}
}
func assertNoErr(t *testing.T, err error, desc string, a ...any) {
	if err != nil {
		t.Errorf("expected no error for %s but was %v", fmt.Sprintf(desc, a...), err)
	}
}

func assertTrue(t *testing.T, value any, desc string, a ...any) {
	if value != true {
		t.Errorf("expected true for %s but got %+v", fmt.Sprintf(desc, a...), value)
	}
}

func assertFalse(t *testing.T, value any, desc string, a ...any) {
	if value != false {
		t.Errorf("expected false for %s but got %+v", fmt.Sprintf(desc, a...), value)
	}
}
