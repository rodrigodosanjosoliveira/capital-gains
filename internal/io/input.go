package io

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func ReadInputFrom(r io.Reader) ([]models.Operation, error) {
	scanner := bufio.NewScanner(r)

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

func ReadInput() ([]models.Operation, error) {
	return ReadInputFrom(os.Stdin)
}
