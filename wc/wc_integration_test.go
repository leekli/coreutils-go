package main

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLI_FlagsForFileInput(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantCount string
	}{
		{name: "bytes", args: []string{"-c", "test_files/test_ascii.txt"}, wantCount: "15088"},
		{name: "lines", args: []string{"-l", "test_files/test_ascii.txt"}, wantCount: "204"},
		{name: "words", args: []string{"-w", "test_files/test_ascii.txt"}, wantCount: "2000"},
		{name: "chars", args: []string{"-m", "test_files/test_ascii.txt"}, wantCount: "15088"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "."}, tt.args...)...)
			cmd.Dir = "."
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go run failed: %v\n%s", err, string(out))
			}
			output := string(out)
			if !strings.Contains(output, tt.wantCount) {
				t.Fatalf("expected output to contain %q, got %q", tt.wantCount, output)
			}
			if !strings.Contains(output, "test_ascii.txt") {
				t.Fatalf("expected output to include the file name, got %q", output)
			}
		})
	}
}

func TestCLI_NoFlagsForFileInput(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "test_files/test_ascii.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}

	output := string(out)
	for _, want := range []string{"204", "2000", "15088"} {
		if !strings.Contains(output, want) {
			t.Fatalf("expected output to contain %q, got %q", want, output)
		}
	}
}

func TestCLI_StdinInputWithFlags(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		input     string
		wantCount string
	}{
		{name: "bytes", args: []string{"-c"}, input: "hello world\n", wantCount: "12"},
		{name: "lines", args: []string{"-l"}, input: "a\nb\nc\n", wantCount: "3"},
		{name: "words", args: []string{"-w"}, input: "one two three\n", wantCount: "3"},
		{name: "chars", args: []string{"-m"}, input: "héllo🙂", wantCount: "6"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "."}, tt.args...)...)
			cmd.Dir = "."
			cmd.Stdin = strings.NewReader(tt.input)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go run failed: %v\n%s", err, string(out))
			}
			if !strings.Contains(string(out), tt.wantCount) {
				t.Fatalf("expected output to contain %q, got %q", tt.wantCount, string(out))
			}
		})
	}
}

func TestCLI_StdinInputWithoutFlags(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = "."
	cmd.Stdin = strings.NewReader("alpha beta\n")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}

	output := string(out)
	for _, want := range []string{"1", "2", "11"} {
		if !strings.Contains(output, want) {
			t.Fatalf("expected output to contain %q, got %q", want, output)
		}
	}
}

func TestCLI_InvalidFlag(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-x", "test_files/test_ascii.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}

	if !strings.Contains(string(out), "Invalid argument flag given.") {
		t.Fatalf("expected invalid flag message, got %q", string(out))
	}
}

func TestCLI_UsesFilePathInOutput(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-c", filepath.Join("test_files", "test_ascii.txt"))
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}

	if !strings.Contains(string(out), "test_ascii.txt") {
		t.Fatalf("expected file name in output, got %q", string(out))
	}
}

func TestCLI_EmptyFileInput(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-w", "test_files/empty_test.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}
	if !strings.Contains(string(out), "0") {
		t.Fatalf("expected zero word count for empty file, got %q", string(out))
	}
}

func TestCLI_UnicodeInputFile(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-m", "test_files/unicode_sample.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}
	if !strings.Contains(string(out), "7") {
		t.Fatalf("expected unicode character count output, got %q", string(out))
	}
}

func TestCLI_WhitespaceOnlyInputFile(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-w", "test_files/whitespace_only.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}
	if !strings.Contains(string(out), "0") {
		t.Fatalf("expected zero words for whitespace-only file, got %q", string(out))
	}
}

func TestCLI_NonexistentFilePanics(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-c", "test_files/does_not_exist.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected command to fail for missing file, output was %q", string(out))
	}
	if !strings.Contains(string(out), "File does not exist") {
		t.Fatalf("expected panic message for missing file, got %q", string(out))
	}
}

func TestCLI_StdinWithNoFlagsAndUnicodeInput(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = "."
	cmd.Stdin = strings.NewReader("héllo🙂\n")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\n%s", err, string(out))
	}
	if !strings.Contains(string(out), "1") || !strings.Contains(string(out), "11") {
		t.Fatalf("expected line and byte counts for unicode stdin input, got %q", string(out))
	}
}
