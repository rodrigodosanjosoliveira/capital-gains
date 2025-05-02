package io

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/models"
)

func TestWriteOutputWithEncoder_Success(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	taxes := []models.Tax{
		{Tax: 100.0},
		{Tax: 0.0},
	}

	err := WriteOutputWithEncoder(taxes, func(v any) ([]byte, error) {
		return []byte(`[{"tax":100},{"tax":0}]`), nil
	})

	w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	expected := `[{"tax":100},{"tax":0}]`
	if output != expected {
		t.Errorf("expected output %s, got %s", expected, output)
	}
}

func TestWriteOutputWithEncoder_Error(t *testing.T) {
	taxes := []models.Tax{{Tax: 123.45}}

	fakeEncoder := func(v any) ([]byte, error) {
		return nil, errors.New("encoder failure")
	}

	err := WriteOutputWithEncoder(taxes, fakeEncoder)
	if err == nil || err.Error() != "encoder failure" {
		t.Errorf("expected encoder failure error, got: %v", err)
	}
}
