package io

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func TestReadInputFlexible_WithArgument(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd", `[{"operation":"buy","unit-cost":10.0,"quantity":100}]`}

	ops, err := ReadInputFlexible()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []models.Operation{{Operation: "buy", UnitCost: 10.0, Quantity: 100}}

	if !reflect.DeepEqual(ops, expected) {
		t.Errorf("expected %v, got %v", expected, ops)
	}
}

func TestReadInputFlexible_FromStdin(t *testing.T) {
	input := `[{"operation":"sell","unit-cost":15.0,"quantity":50}]`
	r := strings.NewReader(input)

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"cmd"}

	ops, err := ReadInputFromReaderOrArg(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []models.Operation{{Operation: "sell", UnitCost: 15.0, Quantity: 50}}

	if !reflect.DeepEqual(ops, expected) {
		t.Errorf("expected %v, got %v", expected, ops)
	}
}

func TestReadInputFlexible_InvalidJSON(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd", `{"invalid":"json"}`} // inválido: não é array

	_, err := ReadInputFlexible()
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestReadInputFlexible_EmptyInput(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd", ""}

	_, err := ReadInputFlexible()
	if err == nil {
		t.Error("expected error for empty input, got nil")
	}
}
