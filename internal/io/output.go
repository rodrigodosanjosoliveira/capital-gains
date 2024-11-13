package io

import (
	"encoding/json"
	"fmt"
	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func WriteOutput(taxes []models.Tax) error {
	output, err := json.Marshal(taxes)
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
