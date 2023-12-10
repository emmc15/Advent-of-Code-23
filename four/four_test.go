package four

import (
	"fmt"
	"os"
	"slices"
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

func TestExampleCasePuzzle1(t *testing.T) {
	file_path, err := os.Getwd()
	if err != nil {
		t.Error("Error getting current directory")
	}
	file_path += "/example.txt"
	example, _ := readFileToList(file_path)

	output, err2 := parseCardStringArray(example)
	if err2 != nil {
		t.Error("Error parsing cards")
		return
	}

	// Checks output for card game 1
	card1 := output[0]
	matchedNumbers := card1.getMatchedWinningNumbers()
	expectedMatchedNumbers := []int{48, 83, 17, 86}

	for _, expectedMatchedNumber := range expectedMatchedNumbers {
		if slices.Contains(matchedNumbers, expectedMatchedNumber) == false {
			t.Error("Expected matched number not found")
		}
	}

	var actualWinningPoints int
	for _, card := range output {
		actualWinningPoints += card.getWinningPoints()
	}

	expectedWinningPoints := 13
	if actualWinningPoints != expectedWinningPoints {
		t.Error("Expected winning points not found")
	}
}

func TestExampleCasePuzzle2(t *testing.T) {
	file_path, err := os.Getwd()
	if err != nil {
		t.Error("Error getting current directory")
	}
	file_path += "/example.txt"
	example, _ := readFileToList(file_path)

	output, err2 := parseCardStringArray(example)
	if err2 != nil {
		t.Error("Error parsing cards")
		return
	}
	cardCount := getCardCounter(output, []int{}, []int{})
	actualCardCount := cardCount.countTotalCards()

	if actualCardCount != 30 {
		fmt.Println(actualCardCount)
		fmt.Println(cardCount)
		t.Error("Expected total cards not found")
	}
}

func TestMain(t *testing.T) {
	main()
}
