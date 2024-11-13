package io

import (
	"bufio"
	"encoding/json"
	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
	"os"
)

func ReadInput() ([]models.Operation, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var input string

	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var operations []models.Operation

	if err := json.Unmarshal([]byte(input), &operations); err != nil {
		return nil, err
	}

	return operations, nil
}
