package calculator

import "testing"

func TestUpdateWeightedAverage(t *testing.T) {
	tests := []struct {
		name           string
		currentAvg     float64
		currentQty     int
		newUnitCost    float64
		newQty         int
		expectedAvg    float64
		expectedNewQty int
	}{
		{
			name:           "First purchase",
			currentAvg:     0,
			currentQty:     0,
			newUnitCost:    10.0,
			newQty:         100,
			expectedAvg:    10.0,
			expectedNewQty: 100,
		},
		{
			name:           "Weighted average of two purchases",
			currentAvg:     10.0,
			currentQty:     100,
			newUnitCost:    20.0,
			newQty:         100,
			expectedAvg:    15.0,
			expectedNewQty: 200,
		},
		{
			name:           "Additional smaller batch",
			currentAvg:     15.0,
			currentQty:     200,
			newUnitCost:    30.0,
			newQty:         50,
			expectedAvg:    18.0,
			expectedNewQty: 250,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avg, qty := updateWeightedAverage(tt.currentAvg, tt.currentQty, tt.newUnitCost, tt.newQty)
			if avg != tt.expectedAvg || qty != tt.expectedNewQty {
				t.Errorf("expected (%.2f, %d), got (%.2f, %d)", tt.expectedAvg, tt.expectedNewQty, avg, qty)
			}
		})
	}
}
