package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	input := `[{"operation":"buy", "unit-cost":10.00, "quantity":100},{"operation":"sell", "unit-cost":15.00, "quantity":50},{"operation":"sell", "unit-cost":15.00, "quantity":50}]`

	expectedOutput := `[{"tax":0.0},{"tax":0.0},{"tax":0.0}]`

	oldStdin := os.Stdin

	oldStdout := os.Stdout

	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	stdinReader, stdinWriter, _ := os.Pipe()

	stdoutReader, stdoutWriter, _ := os.Pipe()

	os.Stdin = stdinReader

	os.Stdout = stdoutWriter

	go func() {
		defer func(stdinWriter *os.File) {
			err := stdinWriter.Close()
			if err != nil {
				t.Errorf("Error closing stdin writer: %v", err)
			}
		}(stdinWriter)
		_, _ = stdinWriter.Write([]byte(input + "\n"))
	}()

	var outputBuffer bytes.Buffer
	go func() {
		defer func(stdoutReader *os.File) {
			err := stdoutReader.Close()
			if err != nil {
				t.Errorf("Error closing stdout reader: %v", err)
			}
		}(stdoutReader)
		_, _ = outputBuffer.ReadFrom(stdoutReader)
	}()

	main()

	_ = stdoutWriter.Close()

	output := strings.TrimSpace(outputBuffer.String())
	if output != expectedOutput {
		t.Errorf("Output mismatch:\nExpected: %s\nGot: %s", expectedOutput, output)
	}
}
