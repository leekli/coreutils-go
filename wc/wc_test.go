package main

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"wc/utils"

	"github.com/stretchr/testify/assert"
)

func readTestFile(filePath string) []byte {
	fileContents := utils.ReadFile(filePath)

	return fileContents
}

func TestCountBytes_ReturnsCorrectBytesCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountBytes(input)

	assert.Equal(test, 0, output)
}

func TestCountBytes_ReturnsCorrectBytesCountForUTF8File(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountBytes(input)

	assert.Equal(test, 342190, output)
}

func TestCountBytes_ReturnsCorrectBytesCountForASCIIFile(test *testing.T) {
	input := readTestFile("test_files/test_ascii.txt")

	output := CountBytes(input)

	assert.Equal(test, 15088, output)
}

func TestCountLines_ReturnsCorrectLineCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountLines(input)

	assert.Equal(test, 0, output)
}

func TestCountLines_ReturnsCorrectLineCountForFileWithContents(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountLines(input)

	assert.Equal(test, 7145, output)

	input = readTestFile("test_files/test_ascii.txt")

	output = CountLines(input)

	assert.Equal(test, 204, output)
}

func TestCountLines_ReturnsCorrectLineCountForVariousNewLineChars(test *testing.T) {
	input := readTestFile("test_files/test_newlines.txt")

	output := CountLines(input)

	assert.Equal(test, 8, output)
}

func TestCountWords_ReturnsCorrectWordCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountWords(input)

	assert.Equal(test, 0, output)
}

func TestCountWords_ReturnsCorrectWordCountForFileWithContents(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountWords(input)

	assert.Equal(test, 58164, output)

	input = readTestFile("test_files/test_ascii.txt")

	output = CountWords(input)

	assert.Equal(test, 2000, output)
}

func TestCountChars_ReturnsCorrectCharCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountChars(input)

	assert.Equal(test, 0, output)
}

func TestCountChars_ReturnsCorrectCharCountForFileWithContents(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountChars(input)

	assert.Equal(test, 339292, output)

	input = readTestFile("test_files/test_ascii.txt")

	output = CountChars(input)

	assert.Equal(test, 15088, output)
}

func TestCountWithNoFlags_ReturnsCorrectStringOfCountsForUTF8File(test *testing.T) {
	fileName := "test_files/test.txt"

	input := readTestFile(fileName)

	outputStr := CountWithNoFlags()
	byteCount := CountBytes(input)
	lineCount := CountLines(input)
	wordCount := CountWords(input)

	byteCountStr := strconv.Itoa(byteCount)
	lineCountStr := strconv.Itoa(lineCount)
	wordCountStr := strconv.Itoa(wordCount)

	if strings.Contains(outputStr, byteCountStr) {
        test.Errorf("Expected: %s to be present, but it was not.", byteCountStr)
    }

    if strings.Contains(outputStr, lineCountStr) {
        test.Errorf("Expected: %s to be present, but it was not.", lineCountStr)
    }

    if strings.Contains(outputStr, wordCountStr) {
        test.Errorf("Expected: %s to be present, but it was not.", wordCountStr)
    }

    if strings.Contains(outputStr, fileName) {
        test.Errorf("Expected: %s to be present, but it was not.", fileName)
    }
}

func TestCountWithNoFlags_ReturnsCorrectStringOfCountsForASCIIFile(test *testing.T) {
	fileName := "test_files/test_ascii.txt"

	input := readTestFile(fileName)

	outputStr := CountWithNoFlags()
	byteCount := CountBytes(input)
	lineCount := CountLines(input)
	wordCount := CountWords(input)

	byteCountStr := strconv.Itoa(byteCount)
	lineCountStr := strconv.Itoa(lineCount)
	wordCountStr := strconv.Itoa(wordCount)

    if strings.Contains(outputStr, byteCountStr) {
        test.Errorf("Expected: %s to be present, but it was not.", byteCountStr)
    }

    if strings.Contains(outputStr, lineCountStr) {
        test.Errorf("Expected: %s to be present, but it was not.", lineCountStr)
    }

    if strings.Contains(outputStr, wordCountStr) {
        test.Errorf("Expected: %s to be present, but it was not.", wordCountStr)
    }

    if strings.Contains(outputStr, fileName) {
        test.Errorf("Expected: %s to be present, but it was not.", fileName)
    }
}

func TestCountChars_MultibyteCharacters(test *testing.T) {
    input := []byte("é😊漢") // 3 runes, multiple bytes
    chars := CountChars(input)
    bytes := CountBytes(input)

	assert.Equal(test, 3, chars)
	assert.NotEqual(test, bytes, chars, "Expected bytes and chars to differ for multibyte runes")
}

func TestCountLines_NoTrailingNewline(test *testing.T) {
    input := []byte("one line without newline at end")
    output := CountLines(input)

	assert.Equal(test, 0, output)
}

func TestCountLines_OnlyNewlines(test *testing.T) {
    input := []byte("\n\n\n")
    output := CountLines(input)

	assert.Equal(test, 3, output)
}

func TestCountWords_TabsAndNewlinesAreSeparators(test *testing.T) {
    input := []byte("one\ttwo\nthree  four")
    output := CountWords(input)

	assert.Equal(test, 4, output)
}

func TestCountWords_PunctuationIsPartOfWords(test *testing.T) {
    input := []byte("hello, world! this.is")
    output := CountWords(input)

	assert.Equal(test, 3, output)
}

func TestCountWords_OnlyWhitespaceReturnsZero(test *testing.T) {
    input := []byte("   \t\n  ")
    output := CountWords(input)

	assert.Equal(test, 0, output)
}

func TestCountBytesAndChars_DifferentForEmoji(test *testing.T) {
    input := []byte("😊")
    bytes := CountBytes(input)
    chars := CountChars(input)

	assert.Greater(test, bytes, chars, "Expected bytes > chars for emoji")
	assert.Equal(test, 1, chars)
}

func TestCountWords_DifferentDelimiters_SpacesTabsNewlines(test *testing.T) {
    spaceInput := []byte("one two three")
    tabInput := []byte("one\ttwo\tthree")
    newlineInput := []byte("one\ntwo\nthree")

	assert.Equal(test, 3, CountWords(spaceInput), "Expected 3 words for spaces")
	assert.Equal(test, 3, CountWords(tabInput), "Expected 3 words for tabs")
	assert.Equal(test, 3, CountWords(newlineInput), "Expected 3 words for newlines")
}

func TestCountWords_ApostrophesAndHyphens_AreSingleWords(test *testing.T) {
    input := []byte("don't stop well-known mother-in-law")
	assert.Equal(test, 4, CountWords(input), "Expected 4 words for apostrophes/hyphens")
}

func TestCountChars_OnlyWhitespace(test *testing.T) {
    input := []byte("  \t\n  ") // 2 spaces, tab, newline, 2 spaces = 6 runes
	assert.Equal(test, 6, CountChars(input), "Expected 6 chars for whitespace-only input")
}

func TestMain_WithArgs_CountsFileBytesUsingGoRun(test *testing.T) {
    // run the program in a subprocess to avoid interfering with test process flags/state
    cmd := exec.Command("go", "run", ".", "-c", "test_files/test_ascii.txt")
    cmd.Dir = "."
    out, err := cmd.CombinedOutput()
	assert.NoError(test, err, "running go run failed: %s", string(out))
	assert.Contains(test, string(out), "15088", "Expected output to contain byte count 15088")
}

func TestMain_ReadsFromStdin_CountsWords(test *testing.T) {
    // pipe input into the program and ask for word count
    cmd := exec.Command("bash", "-c", "printf 'one two three\n' | go run . -w")
    cmd.Dir = "."
    out, err := cmd.CombinedOutput()
	assert.NoError(test, err, "running go run with stdin failed: %s", string(out))
	assert.Contains(test, string(out), "3", "Expected output to contain word count 3")
}