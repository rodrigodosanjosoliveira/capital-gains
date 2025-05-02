package main

import (
	"encoding/json"
	"os/exec"
	"strings"
	"testing"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func TestCapitalGains_CLI(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin = strings.NewReader(`[{"operation":"buy", "unit-cost":10.00, "quantity":100},{"operation":"sell", "unit-cost":15.00, "quantity":50}]`)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("execution failed: %v\noutput: %s", err, out)
	}

	var actual []models.Tax
	if err := json.Unmarshal(out, &actual); err != nil {
		t.Fatalf("failed to parse output: %v\noutput: %s", err, out)
	}

	expected := []models.Tax{{Tax: 0.0}, {Tax: 0.0}}

	if len(actual) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("at index %d: expected %+v, got %+v", i, expected[i], actual[i])
		}
	}
}
