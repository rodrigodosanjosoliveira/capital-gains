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
				{"buy", 10.00, 100},
				{"sell", 15.00, 50},
				{"sell", 15.00, 50},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #2",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"sell", 20.00, 5000},
				{"sell", 5.00, 5000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 10000.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #1 + Case #2",
			operations: []models.Operation{
				{"buy", 10.00, 100},
				{"sell", 15.00, 50},
				{"sell", 15.00, 50},
				{"buy", 10.00, 10000},
				{"sell", 20.00, 5000},
				{"sell", 5.00, 5000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
				{Tax: 0.00}, {Tax: 10000.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #3",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"sell", 5.00, 5000},
				{"sell", 20.00, 3000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 1000.00},
			},
		},
		{
			name: "Case #4",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"buy", 25.00, 5000},
				{"sell", 15.00, 10000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #5",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"buy", 25.00, 5000},
				{"sell", 15.00, 10000},
				{"sell", 25.00, 5000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 10000.00},
			},
		},
		{
			name: "Case #6",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"sell", 2.00, 5000},
				{"sell", 20.00, 2000},
				{"sell", 20.00, 2000},
				{"sell", 25.00, 1000},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 3000.00},
			},
		},
		{
			name: "Case #7",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"sell", 2.00, 5000},
				{"sell", 20.00, 2000},
				{"sell", 20.00, 2000},
				{"sell", 25.00, 1000},
				{"buy", 20.00, 10000},
				{"sell", 15.00, 5000},
				{"sell", 30.00, 4350},
				{"sell", 30.00, 650},
			},
			expected: []models.Tax{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 3000.00},
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 3700.00}, {Tax: 0.00},
			},
		},
		{
			name: "Case #8",
			operations: []models.Operation{
				{"buy", 10.00, 10000},
				{"sell", 50.00, 10000},
				{"buy", 20.00, 10000},
				{"sell", 50.00, 10000},
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
