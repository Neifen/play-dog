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
