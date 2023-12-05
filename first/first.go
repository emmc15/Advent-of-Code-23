package first

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFileToList(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func parseElfNumberFromString(numberString string) (int, error) {

	regex_string := "[\\d+]"
	re := regexp.MustCompile(regex_string)
	// number_groups := re.FindStringSubmatch(numberString)
	// Extract all the number groups from the string
	number_groups := re.FindAllString(numberString, -1)

	digits := len(number_groups)

	numberString = number_groups[0]
	numberString += number_groups[digits-1]
	// numberString = strings.Join(number_groups, "")
	// convert numberString to int
	number, err := strconv.Atoi(numberString)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func parseElfNumberFromStringArray(numberStringArray []string) ([]int, error) {
	var numbers []int
	for _, numberString := range numberStringArray {
		number, err := parseElfNumberFromString(numberString)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func convertEnglishDigitFromString(digitString string) (string, error) {
	// dictionary of english numbers
	englishNumbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// replace any isntances of digit words with ints
	for key, value := range englishNumbers {
		digitString = strings.ReplaceAll(digitString, key, strconv.Itoa(value))
	}

	return digitString, nil
}

func parseEnglishDigitFromStringArray(digitStringArray []string) ([]string, error) {
	var digits []string
	for _, digitString := range digitStringArray {
		digit, err := convertEnglishDigitFromString(digitString)
		if err != nil {
			return nil, err
		}
		digits = append(digits, digit)
	}
	return digits, nil
}

func sumIntArray(intArray []int) int {
	var sum int
	for _, number := range intArray {
		sum += number
	}
	return sum
}

func main() {

	file_path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		return
	}
	file_path += "/day1_puzzle.txt"

	lines, err := readFileToList(file_path)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	numbers, err := parseElfNumberFromStringArray(lines)
	if err != nil {
		fmt.Println("Error parsing numbers")
		return
	}
	sum := sumIntArray(numbers)
	fmt.Println("puzzle 1:", sum)

	// puzzle 2
	// convert all the numbers to english
	englishNumbers, err := parseEnglishDigitFromStringArray(lines)
	if err != nil {
		fmt.Println("Error parsing numbers")
		return
	}
	// convert all the english numbers to ints
	elfNumbers, err := parseElfNumberFromStringArray(englishNumbers)
	if err != nil {
		fmt.Println("Error parsing numbers")
		return
	}
	fmt.Println(englishNumbers)
	sumPuzzle2 := sumIntArray(elfNumbers)
	fmt.Println("puzzle 2:", sumPuzzle2)

}
