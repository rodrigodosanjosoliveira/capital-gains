package calculator

import (
	"reflect"
	"testing"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func TestCalculateCapitalGains(t *testing.T) {
	tests := []struct {
		name       string
		operations []models.Operation
		expected   []models.Tax
	}{
		{
			name: "Case #1",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 100},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #2",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 10000.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #1 + Case #2",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 100},
				{Operation: "sell", UnitCost: 5.00, Quantity: 50},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
				{Tax: 0.00}, {Tax: 10000.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #3",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 3000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 1000.00},
			},
		},
		{
			name: "Case #4",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #5",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 5000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 10000.00},
			},
		},
		{
			name: "Case #6",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 3000.00},
			},
		},
		{
			name: "Case #7",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 30.00, Quantity: 4350},
				{Operation: "sell", UnitCost: 30.00, Quantity: 650},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 3000.00},
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 3700.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #8",
			operations: []models.Operation{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 80000.00}, {Tax: 0.00}, {Tax: 60000.00},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateCapitalGains(tt.operations)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
