package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/cody-smarty/calc-lib"
)

func TestHandler_WrongNumberOfArgs(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	if !errors.Is(err, errWrongArgCount) {
		t.Errorf("want :%v, got: %v", errWrongArgCount, err)
	}
}

func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"invalid", "42"})
	if !errors.Is(err, errInvalidArgument) {
		t.Errorf("want :%v, got: %v", errInvalidArgument, err)
	}
}

func TestHandler_SecondArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"42", "invalid"})
	if !errors.Is(err, errInvalidArgument) {
		t.Errorf("want :%v, got: %v", errInvalidArgument, err)
	}
}

func TestHandler_ResultWrittenToOutput(t *testing.T) {
	stdout := &bytes.Buffer{}
	handler := NewHandler(stdout, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	if err != nil {
		t.Errorf("want: %v, got: %v", nil, err)
	}
	if stdout.String() != "5" {
		t.Errorf("want: %v, got: %v", "5", stdout.String())
	}
}
