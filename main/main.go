package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/cody-smarty/calc-lib"
)

func main() {
	handler := NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}

func (hand *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[1])
	}
	result := hand.calculator.Calculate(a, b)
	_, err = fmt.Fprint(hand.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: '%s'", errOutputFailure, args[0])
	}
	return err
}

var (
	errWrongArgCount   = errors.New("usage: calculator <a> <b>")
	errInvalidArgument = errors.New("invalid argument")
	errOutputFailure   = errors.New("output failure")
)
