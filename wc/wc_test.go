package main

import (
	"os/exec"
	"strconv"
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

	oldFileContents := fileContents
	oldFilePathGiven := filePathGiven
	fileContents = input
	filePathGiven = fileName
	defer func() {
		fileContents = oldFileContents
		filePathGiven = oldFilePathGiven
	}()

	outputStr := CountWithNoFlags()
	byteCount := CountBytes(input)
	lineCount := CountLines(input)
	wordCount := CountWords(input)

	assert.Contains(test, outputStr, strconv.Itoa(byteCount))
	assert.Contains(test, outputStr, strconv.Itoa(lineCount))
	assert.Contains(test, outputStr, strconv.Itoa(wordCount))
	assert.Contains(test, outputStr, fileName)
}

func TestCountWithNoFlags_ReturnsCorrectStringOfCountsForASCIIFile(test *testing.T) {
	fileName := "test_files/test_ascii.txt"
	input := readTestFile(fileName)

	oldFileContents := fileContents
	oldFilePathGiven := filePathGiven
	fileContents = input
	filePathGiven = fileName
	defer func() {
		fileContents = oldFileContents
		filePathGiven = oldFilePathGiven
	}()

	outputStr := CountWithNoFlags()
	byteCount := CountBytes(input)
	lineCount := CountLines(input)
	wordCount := CountWords(input)

	assert.Contains(test, outputStr, strconv.Itoa(byteCount))
	assert.Contains(test, outputStr, strconv.Itoa(lineCount))
	assert.Contains(test, outputStr, strconv.Itoa(wordCount))
	assert.Contains(test, outputStr, fileName)
}

func TestCountChars_MultibyteCharacters(test *testing.T) {
	input := []byte("é😊漢")
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

func TestCountLines_MixedNewlineStyles(test *testing.T) {
	input := []byte("one\ntwo\r\nthree\rfour")
	output := CountLines(input)

	assert.Equal(test, 2, output)
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

func TestCountWords_MultipleConsecutiveWhitespace(test *testing.T) {
	input := []byte("a   b\t\t c\n\nd\r\ne")
	output := CountWords(input)

	assert.Equal(test, 5, output)
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
	input := []byte("  \t\n  ")
	assert.Equal(test, 6, CountChars(input), "Expected 6 chars for whitespace-only input")
}

func TestCountChars_UnicodeAndASCIIMix(test *testing.T) {
	input := []byte("AΩ€🙂")
	assert.Equal(test, 4, CountChars(input))
}

func TestCountBytes_EmptyInputAndZeroLengthSlice(test *testing.T) {
	assert.Equal(test, 0, CountBytes(nil))
	assert.Equal(test, 0, CountBytes([]byte{}))
}

func TestCountLines_EmptyInputAndZeroLengthSlice(test *testing.T) {
	assert.Equal(test, 0, CountLines(nil))
	assert.Equal(test, 0, CountLines([]byte{}))
}

func TestCountWords_EmptyInputAndZeroLengthSlice(test *testing.T) {
	assert.Equal(test, 0, CountWords(nil))
	assert.Equal(test, 0, CountWords([]byte{}))
}

func TestCountChars_EmptyInputAndZeroLengthSlice(test *testing.T) {
	assert.Equal(test, 0, CountChars(nil))
	assert.Equal(test, 0, CountChars([]byte{}))
}

func TestCountWithNoFlags_UsesCurrentGlobalState(test *testing.T) {
	oldFileContents := fileContents
	oldFilePathGiven := filePathGiven
	fileContents = []byte("a b\nc")
	filePathGiven = "custom.txt"
	defer func() {
		fileContents = oldFileContents
		filePathGiven = oldFilePathGiven
	}()

	output := CountWithNoFlags()
	assert.Contains(test, output, "1")
	assert.Contains(test, output, "3")
	assert.Contains(test, output, "5")
	assert.Contains(test, output, "custom.txt")
}

func TestMain_WithArgs_CountsFileBytesUsingGoRun(test *testing.T) {
	cmd := exec.Command("go", "run", ".", "-c", "test_files/test_ascii.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	assert.NoError(test, err, "running go run failed: %s", string(out))
	assert.Contains(test, string(out), "15088", "Expected output to contain byte count 15088")
}

func TestMain_ReadsFromStdin_CountsWords(test *testing.T) {
	cmd := exec.Command("bash", "-c", "printf 'one two three\n' | go run . -w")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	assert.NoError(test, err, "running go run with stdin failed: %s", string(out))
	assert.Contains(test, string(out), "3", "Expected output to contain word count 3")
}

func TestMain_UsesFileInputAndNoFlagsWithRelativePath(test *testing.T) {
	cmd := exec.Command("go", "run", ".", "test_files/empty_test.txt")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	assert.NoError(test, err, "running go run failed: %s", string(out))
	assert.Contains(test, string(out), "test_files/empty_test.txt")
}