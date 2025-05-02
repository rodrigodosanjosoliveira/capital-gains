package io

import (
	"encoding/json"
	"fmt"
	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func WriteOutput(taxes []models.Tax) error {
	return WriteOutputWithEncoder(taxes, json.Marshal)
}

func WriteOutputWithEncoder(taxes []models.Tax, encode func(v any) ([]byte, error)) error {
	output, err := encode(taxes)
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
