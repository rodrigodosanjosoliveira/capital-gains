package models

import "fmt"

type (
	Operation struct {
		Operation string  `json:"operation"`
		UnitCost  float64 `json:"unit-cost"`
		Quantity  int     `json:"quantity"`
	}

	Tax struct {
		Tax float64 `json:"-"`
	}
)

func (t Tax) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"tax":%.1f}`, t.Tax)), nil
}
