package first

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {

	file_path, err := os.Getwd()
	if err != nil {
		t.Error("Error getting current directory")
	}
	file_path += "/day1_puzzle.txt"
	readFileToList(file_path)
}

func TestParseElfNumberFromString(t *testing.T) {
	number, err := parseElfNumberFromString("1Fds23")
	if err != nil {
		t.Error("Error from parsing function")
	}
	if number != 13 {
		t.Error("Number not parsed correctly")
	}
}

func TestParseElfNumberFromStringArray(t *testing.T) {
	numberArray, err := parseElfNumberFromStringArray([]string{"1Fds23", "1Fds23"})
	if err != nil {
		t.Error("Error from parsing function")
	}
	if numberArray[0] != 13 {
		t.Error("Number not parsed correctly")
	}
	if numberArray[1] != 13 {
		t.Error("Number not parsed correctly")
	}
}

func TestSumIntArray(t *testing.T) {
	sum := sumIntArray([]int{1, 2, 3})
	if sum != 6 {
		t.Error("Sum not calculated correctly")
	}
}

func TestConvertEnglishDigitFromString(t *testing.T) {
	digit, err := convertEnglishDigitFromString("onetwentysixsomething5")
	if err != nil {
		t.Error("Error from parsing function")
	}
	if digit != "1twenty6something5" {
		t.Error("Digit not parsed correctly")
	}
}

func TestMain(t *testing.T) {
	main()
}
