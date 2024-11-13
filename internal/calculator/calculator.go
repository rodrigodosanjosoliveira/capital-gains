package calculator

import "github.com/rodrigodosanjosoliveira/capital-gains/internal/models"

func CalculateCapitalGains(operations []models.Operation) []models.Tax {
	var taxes []models.Tax
	var totalQuantity int
	var weightedAverage float64
	var accumulatedLoss float64

	for _, op := range operations {
		switch op.Operation {
		case "buy":
			totalCost := weightedAverage*float64(totalQuantity) + op.UnitCost*float64(op.Quantity)
			totalQuantity += op.Quantity
			weightedAverage = totalCost / float64(totalQuantity)
			taxes = append(taxes, models.Tax{Tax: 0.0})

		case "sell":
			totalValue := op.UnitCost * float64(op.Quantity)

			if totalValue <= 20000 {
				profit := totalValue - (float64(op.Quantity) * weightedAverage)
				if profit < 0 {
					accumulatedLoss += -profit
				}
				taxes = append(taxes, models.Tax{Tax: 0.0})
				totalQuantity -= op.Quantity
				continue
			}

			profit := totalValue - (float64(op.Quantity) * weightedAverage)

			if profit > 0 {
				if accumulatedLoss > 0 {
					if profit <= accumulatedLoss {
						accumulatedLoss -= profit
						profit = 0
					} else {
						profit -= accumulatedLoss
						accumulatedLoss = 0
					}
				}
				tax := profit * 0.20
				taxes = append(taxes, models.Tax{Tax: tax})
			} else {
				accumulatedLoss += -profit
				taxes = append(taxes, models.Tax{Tax: 0.0})
			}

			totalQuantity -= op.Quantity
		}
	}

	return taxes
}
