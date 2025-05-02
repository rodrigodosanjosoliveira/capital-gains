package io

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func ReadInputFromReaderOrArg(r io.Reader) ([]models.Operation, error) {
	if len(os.Args) > 1 {
		return parseOperations(os.Args[1])
	}

	var input string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		input += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parseOperations(input)
}

func ReadInputFlexible() ([]models.Operation, error) {
	return ReadInputFromReaderOrArg(os.Stdin)
}

func parseOperations(input string) ([]models.Operation, error) {
	if input == "" {
		return nil, errors.New("input is empty")
	}
	var operations []models.Operation
	if err := json.Unmarshal([]byte(input), &operations); err != nil {
		return nil, err
	}
	return operations, nil
}
