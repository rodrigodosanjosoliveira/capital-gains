package calculator

import (
	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
	"testing"
)

func TestHandleSellOperation(t *testing.T) {
	tests := []struct {
		name            string
		op              models.Operation
		weightedAvg     float64
		totalQty        int
		accumulatedLoss float64
		expectedTax     float64
		expectedQty     int
		expectedNewLoss float64
	}{
		{
			name:            "Sell below 20k is tax-free",
			op:              models.Operation{"sell", 10.0, 1000},
			weightedAvg:     5.0,
			totalQty:        10000,
			accumulatedLoss: 0,
			expectedTax:     0.0,
			expectedQty:     9000,
			expectedNewLoss: 0,
		},
		{
			name:            "Loss is accumulated",
			op:              models.Operation{Operation: "sell", UnitCost: 5.0, Quantity: 1000},
			weightedAvg:     10.0,
			totalQty:        10000,
			accumulatedLoss: 0,
			expectedTax:     0.0,
			expectedQty:     9000,
			expectedNewLoss: 5000.0,
		},
		{
			name:            "Profit absorbed by loss",
			op:              models.Operation{"sell", 20.0, 1000},
			weightedAvg:     10.0,
			totalQty:        10000,
			accumulatedLoss: 10000.0,
			expectedTax:     0.0,
			expectedQty:     9000,
			expectedNewLoss: 10000.0, // <- corrigido!
		},
		{
			name:            "Profit generates tax after loss used",
			op:              models.Operation{"sell", 20.0, 1000},
			weightedAvg:     10.0,
			totalQty:        10000,
			accumulatedLoss: 5000.0,
			expectedTax:     0.00, // CORRIGIDO
			expectedQty:     9000,
			expectedNewLoss: 5000.00, // CORRIGIDO
		},
		{
			name:            "Full profit tax when no loss",
			op:              models.Operation{"sell", 30.0, 1000},
			weightedAvg:     10.0,
			totalQty:        10000,
			accumulatedLoss: 0.0,
			expectedTax:     4000.0, // lucro 20k * 20%
			expectedQty:     9000,
			expectedNewLoss: 0.0,
		},
		{
			name:            "Negative profit still adds loss",
			op:              models.Operation{"sell", 8.0, 1000},
			weightedAvg:     10.0,
			totalQty:        10000,
			accumulatedLoss: 200.0,
			expectedTax:     0.0,
			expectedQty:     9000,
			expectedNewLoss: 2200.0, // acumulado + 2k de prejuízo
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tax, qty, newLoss := handleSellOperation(tt.op, tt.weightedAvg, tt.totalQty, tt.accumulatedLoss)

			if tax != tt.expectedTax || qty != tt.expectedQty || newLoss != tt.expectedNewLoss {
				t.Errorf("expected (tax=%.2f, qty=%d, loss=%.2f), got (tax=%.2f, qty=%d, loss=%.2f)",
					tt.expectedTax, tt.expectedQty, tt.expectedNewLoss,
					tax, qty, newLoss)
			}
		})
	}
}
