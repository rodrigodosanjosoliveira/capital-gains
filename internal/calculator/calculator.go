package calculator

import (
	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func CalculateCapitalGains(operations []models.Operation) []models.Tax {
	var (
		taxes           []models.Tax
		totalQuantity   int
		weightedAverage float64
		accumulatedLoss float64
	)

	for _, op := range operations {
		switch op.Operation {
		case "buy":
			weightedAverage, totalQuantity = updateWeightedAverage(weightedAverage, totalQuantity, op.UnitCost, op.Quantity)
			taxes = append(taxes, models.Tax{Tax: 0.0})

		case "sell":
			tax, updatedQuantity, updatedLoss := handleSellOperation(op, weightedAverage, totalQuantity, accumulatedLoss)
			taxes = append(taxes, models.Tax{Tax: tax})
			totalQuantity = updatedQuantity
			accumulatedLoss = updatedLoss
		}
	}

	return taxes
}

func updateWeightedAverage(currentAverage float64, currentQty int, unitCost float64, qty int) (float64, int) {
	totalCost := currentAverage*float64(currentQty) + unitCost*float64(qty)
	newQty := currentQty + qty
	newAverage := totalCost / float64(newQty)
	return newAverage, newQty
}

func handleSellOperation(
	op models.Operation,
	weightedAverage float64,
	totalQuantity int,
	accumulatedLoss float64,
) (float64, int, float64) {
	saleValue := op.UnitCost * float64(op.Quantity)
	profit := saleValue - float64(op.Quantity)*weightedAverage
	newQuantity := totalQuantity - op.Quantity

	if saleValue <= 20000 {
		if profit < 0 {
			accumulatedLoss += -profit
		}
		return 0.0, newQuantity, accumulatedLoss
	}

	if profit <= 0 {
		accumulatedLoss += -profit
		return 0.0, newQuantity, accumulatedLoss
	}

	// lucro > 0
	if accumulatedLoss > 0 {
		if profit <= accumulatedLoss {
			accumulatedLoss -= profit
			return 0.0, newQuantity, accumulatedLoss
		}
		profit -= accumulatedLoss
		accumulatedLoss = 0
	}

	tax := profit * 0.20
	return tax, newQuantity, accumulatedLoss
}
