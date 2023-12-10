package third

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {

	file_path, err := os.Getwd()
	if err != nil {
		t.Error("Error getting current directory")
	}
	file_path += "/day2_puzzle.txt"
	readFileToList(file_path)
}

func TestParseSchematicStringArray(t *testing.T) {
	file_path, err := os.Getwd()
	if err != nil {
		t.Error("Error getting current directory")
	}
	file_path += "/example.txt"
	example, _ := readFileToList(file_path)

	output := parseSchematicStringArray(example)

	expectedFalseParts := []int{114, 58}
	var actualTrueParts []int
	var actualFalseParts []int
	for _, part := range output {
		if part.isPartNumber() == false {
			fmt.Printf("part: %v\n", part)
			actualFalseParts = append(actualFalseParts, part.possiblePartNumber)
		} else {
			actualTrueParts = append(actualTrueParts, part.possiblePartNumber)
		}
	}
	for _, expectedFalsePart := range expectedFalseParts {
		if slices.Contains(actualFalseParts, expectedFalsePart) == false {
			t.Error("Expected false part not found")
		}
	}

	expectedTruePartSum := 4361
	actualTruePartSum := sumIntArray(actualTrueParts)
	if actualTruePartSum != expectedTruePartSum {
		t.Error("Expected true part sum not found")
	}
}

func TestGetGearActualLocation(t *testing.T) {
	file_path, err := os.Getwd()
	if err != nil {
		t.Error("Error getting current directory")
	}
	file_path += "/day3_puzzle.txt"
	example, _ := readFileToList(file_path)

	output := parseSchematicStringArray(example)

	var actualParts []PartFinder
	for _, part := range output {
		if part.isPartNumber() == true {
			actualParts = append(actualParts, part)
		}
	}

	for _, actualPart := range actualParts {
		locations := actualPart.getGearActualLocation()
		for _, location := range locations {
			for colIndex, stringFound := range strings.Split(example[location[0]], "") {
				if colIndex == location[1] {
					if stringFound != "*" {
						t.Error("Expected gear not found")
					}
				}
			}
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
