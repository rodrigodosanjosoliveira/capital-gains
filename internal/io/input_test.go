package io

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func TestReadInputFrom(t *testing.T) {
	jsonInput := `[{"operation":"buy","unit-cost":10.0,"quantity":100},{"operation":"sell","unit-cost":15.0,"quantity":50}]`
	reader := strings.NewReader(jsonInput)

	operations, err := ReadInputFrom(reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []models.Operation{
		{Operation: "buy", UnitCost: 10.0, Quantity: 100},
		{Operation: "sell", UnitCost: 15.0, Quantity: 50},
	}

	if !reflect.DeepEqual(operations, expected) {
		t.Errorf("expected %v, got %v", expected, operations)
	}
}

type errorReader struct{}

func (e *errorReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

func TestReadInputFrom_Errors(t *testing.T) {
	t.Run("should return error when reader fails", func(t *testing.T) {
		_, err := ReadInputFrom(&errorReader{})
		if err == nil || err.Error() != "simulated read error" {
			t.Errorf("expected simulated read error, got: %v", err)
		}
	})

	t.Run("should return error when JSON is invalid", func(t *testing.T) {
		invalidJSON := `{"operation":"buy","unit-cost":10.0,"quantity":100}` // não é um array
		_, err := ReadInputFrom(strings.NewReader(invalidJSON))
		if err == nil {
			t.Errorf("expected json unmarshal error, got nil")
		}
	})

	t.Run("should return error when input is empty", func(t *testing.T) {
		_, err := ReadInputFrom(strings.NewReader(""))
		if err == nil {
			t.Errorf("expected error due to empty input, got nil")
		}
	})
}
