package third

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type PartFinder struct {
	possiblePartNumber int
	targetRowIndex     int
	targetColIndex     int
	targetRow          []string
	rowAbove           []string
	rowBelow           []string
}

func (pf PartFinder) isPartNumber() bool {

	// check if row above has any value other than `.`
	for _, value := range pf.rowAbove {
		if value != "." {
			return true
		}
	}

	// check if row below has any value other than `.`
	for _, value := range pf.rowBelow {
		if value != "." {
			return true
		}
	}

	targetRowValue := ""
	for _, value := range pf.targetRow {
		targetRowValue += value
	}

	possibleNumberString := strconv.Itoa(pf.possiblePartNumber)
	targetRowValue = strings.ReplaceAll(targetRowValue, possibleNumberString, ".")
	for _, value := range targetRowValue {
		if string(value) != "." {
			return true
		}
	}

	return false

}

func (pf PartFinder) getGearRelativeLocation() [][]int {
	// returns [[]] if not gear
	// Gear location is relative layout of * icon, can return multple gear locations
	// Layout of relative follows pattern below:
	// Col0, Col1, Col2
	// [0, 1, 2]: row -1
	// [3, 4, 5]: row 0
	// [6, 7, 8]: row +1

	var gearLocations [][]int
	adjustRelativeCol := -1
	possibleLocationRelativeCheck := pf.targetRow[0]
	_, err := strconv.Atoi(possibleLocationRelativeCheck)
	if err == nil {
		adjustRelativeCol = 0
	}

	if slices.Contains(pf.rowAbove, "*") == true {
		for index, value := range pf.rowAbove {
			if value == "*" {
				gearLocation := []int{-1, index + adjustRelativeCol}
				gearLocations = append(gearLocations, gearLocation)
			}
		}
	}

	if slices.Contains(pf.targetRow, "*") == true {
		for index, value := range pf.targetRow {
			if value == "*" {
				gearLocation := []int{0, index + adjustRelativeCol}
				gearLocations = append(gearLocations, gearLocation)
			}
		}
	}

	if slices.Contains(pf.rowBelow, "*") == true {
		for index, value := range pf.rowBelow {
			if value == "*" {
				gearLocation := []int{1, index + adjustRelativeCol}
				gearLocations = append(gearLocations, gearLocation)
			}
		}
	}

	return gearLocations
}

func (pf PartFinder) getGearActualLocation() [][]int {

	var gearLocations [][]int

	// get relative gear locations
	relativeGearLocations := pf.getGearRelativeLocation()

	// convert relative gear locations to actual gear locations
	for _, relativeGearLocation := range relativeGearLocations {
		columnDistaince := relativeGearLocation[1] + pf.targetColIndex
		gearLocation := []int{pf.targetRowIndex + relativeGearLocation[0], columnDistaince}
		gearLocations = append(gearLocations, gearLocation)
	}

	return gearLocations
}

func readFileToList(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func parseSchematicStringArray(engineStringArray []string) []PartFinder {

	var partFinderArray []PartFinder

	// regex to find numbers
	re := regexp.MustCompile(`(\d+)`)
	max_index := len(engineStringArray) - 1
	for target_index, target_row := range engineStringArray {

		row_above := ""
		if target_index > 0 {
			row_above = engineStringArray[target_index-1]
		}
		row_below := ""
		if target_index < max_index {
			row_below = engineStringArray[target_index+1]
		}

		found_possible_numbers := re.FindAllStringIndex(target_row, -1)
		if found_possible_numbers == nil {
			continue
		}
		for _, found_possible_number := range found_possible_numbers {

			// Finds the possible number and converts to int
			possible_number := target_row[found_possible_number[0]:found_possible_number[1]]
			possible_number_int, _ := strconv.Atoi(possible_number)

			// Handles for slices above and below with out of bounds
			row_above_slice := ""
			if row_above != "" {
				start_index := found_possible_number[0] - 1
				if start_index < 0 {
					start_index = 0
				}
				end_index := found_possible_number[1] + 1
				if end_index > len(row_above) {
					end_index = len(row_above)
				}
				row_above_slice = row_above[start_index:end_index]
			}
			row_below_slice := ""
			if row_below != "" {
				start_index := found_possible_number[0] - 1
				if start_index < 0 {
					start_index = 0
				}
				end_index := found_possible_number[1] + 1
				if end_index > len(row_below) {
					end_index = len(row_below)
				}
				row_below_slice = row_below[start_index:end_index]
			}

			// Converts slices to arrays
			row_above_slice_array := strings.Split(row_above_slice, "")
			row_below_slice_array := strings.Split(row_below_slice, "")

			// Finds the target slice
			target_row_slice := ""
			start_index := found_possible_number[0] - 1
			if start_index < 0 {
				start_index = 0
			}
			end_index := found_possible_number[1] + 1
			if end_index > len(target_row) {
				end_index = len(target_row)
			}
			target_row_slice = target_row[start_index:end_index]
			target_row_slice_array := strings.Split(target_row_slice, "")

			// Adds the part finder to the array
			partSchema := PartFinder{
				possiblePartNumber: possible_number_int,
				targetRowIndex:     target_index,
				targetColIndex:     found_possible_number[0],
				targetRow:          target_row_slice_array,
				rowAbove:           row_above_slice_array,
				rowBelow:           row_below_slice_array,
			}
			partFinderArray = append(partFinderArray, partSchema)
		}
	}
	return partFinderArray
}

func isGearRatio(s PartFinder, v PartFinder) bool {
	if s.targetRowIndex != v.targetRowIndex {
		return false
	}
	if s.targetColIndex == v.targetColIndex {
		return false
	}
	if s.possiblePartNumber == v.possiblePartNumber {
		return false
	}
	return true
}

func sumIntArray(intArray []int) int {
	var sum int
	for _, number := range intArray {
		sum += number
	}
	return sum
}

func maxIntArray(intArray []int) int {
	var max int
	for _, number := range intArray {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {

	file_path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		return
	}
	file_path += "/day3_puzzle.txt"

	lines, err := readFileToList(file_path)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	partFinderArray := parseSchematicStringArray(lines)
	var trueParts []int
	var truePartFinders []PartFinder
	for _, part := range partFinderArray {
		if part.isPartNumber() == true {
			trueParts = append(trueParts, part.possiblePartNumber)
			truePartFinders = append(truePartFinders, part)
		}
	}
	fmt.Println("puzzle 1: ", sumIntArray(trueParts))

	// puzzle 2 - Find all gear ratios and sum them up
	var gearParts []PartFinder
	var matchedGearParts map[string][]PartFinder

	for _, part := range truePartFinders {
		gearLocations := part.getGearActualLocation()
		if len(gearLocations) > 0 {
			gearParts = append(gearParts, part)
			for _, gearLocation := range gearLocations {
				key := strconv.Itoa(gearLocation[0]) + "," + strconv.Itoa(gearLocation[1])
				if matchedGearParts == nil {
					matchedGearParts = make(map[string][]PartFinder)
				}
				matchedGearParts[key] = append(matchedGearParts[key], part)
			}
		}
	}

	// For matches that have exactly two valid parts
	var validGearParts []PartFinder
	var gearPartRatios []int
	for _, gearPart := range matchedGearParts {
		if len(gearPart) == 2 {
			validGearParts = append(validGearParts, gearPart...)
			gearPartRatios = append(gearPartRatios, gearPart[0].possiblePartNumber*gearPart[1].possiblePartNumber)
		}
	}

	fmt.Println("puzzle 2: ", sumIntArray(gearPartRatios))

}
