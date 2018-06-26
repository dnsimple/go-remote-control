package grc

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestShellScripts(t *testing.T) {
	scriptsDir = "./test-scripts"

	functions := map[string]func(io.Writer) error{
		"status": RunStatus,
		"start":  RunStart,
		"stop":   RunStop,
		"update": RunUpdate,
	}

	for name, function := range functions {
		var w bytes.Buffer
		err := function(&w)
		if err != nil {
			t.Errorf("Error executing %s: %s", function, err)
		}

		actual := w.String()
		expected := fmt.Sprintf("shell %s success\n", name)
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
}

func TestRubyScripts(t *testing.T) {
	scriptsDir = "./test-scripts"
	scriptExt = "rb"

	functions := map[string]func(io.Writer) error{
		"status": RunStatus,
		"start":  RunStart,
		"stop":   RunStop,
		"update": RunUpdate,
	}

	for name, function := range functions {
		var w bytes.Buffer
		err := function(&w)
		if err != nil {
			t.Errorf("Error executing %s: %s", function, err)
		}

		actual := w.String()
		expected := fmt.Sprintf("ruby %s success\n", name)
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
}
