package four

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Util Functions

func readFileToList(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
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

// Puzzle Functions

type Card struct {
	cardNumber     int
	winningNumbers []int
	givenNumbers   []int
}

func (card Card) getMatchedWinningNumbers() []int {
	var matchedWinningNumbers []int
	for _, givenNumber := range card.givenNumbers {
		if slices.Contains(card.winningNumbers, givenNumber) {
			matchedWinningNumbers = append(matchedWinningNumbers, givenNumber)
		}
	}
	return matchedWinningNumbers
}

func (card Card) getWinningPoints() int {
	matchedWinningNumbers := card.getMatchedWinningNumbers()
	if len(matchedWinningNumbers) == 0 {
		return 0
	}

	return int(math.Pow(2, float64(len(matchedWinningNumbers)-1)))
}

func (card Card) getWinningCopiedCardNumbers() []int {
	var copiedCardNumbers []int
	winningMatchedNumbers := card.getMatchedWinningNumbers()
	for index, _ := range winningMatchedNumbers {
		index += 1
		copiedCardNumbers = append(copiedCardNumbers, card.cardNumber+index)
	}

	return copiedCardNumbers
}

func parseCardString(cardString string) (Card, error) {
	var card Card

	if strings.HasPrefix(cardString, "Card") == false {
		fmt.Println("Error: Card string does not start with 'Card'")
		return card, errors.New("Error: Card string does not start with 'Card'")
	}

	cardGameSplit := strings.Split(cardString, ":")
	if len(cardGameSplit) != 2 {
		fmt.Println("cardGameSplit:", cardGameSplit)
		fmt.Println("Error: Card string does not have 2 parts")
		return card, errors.New("Error: Card string does not have 2 parts")
	}

	// Handles Card Number
	cardNumberString := cardGameSplit[0]
	findDigits := regexp.MustCompile(`\d+`)
	cardNumberString = findDigits.FindString(cardNumberString)
	cardNumber, err := strconv.Atoi(cardNumberString)
	if err != nil {
		fmt.Println("Error: Card number could not be parsed")
		fmt.Println("cardNumberString:", cardString)
		return card, err
	}

	// Split Game into parts
	gameNumbersString := cardGameSplit[1]
	gameNumbersStringArray := strings.Split(gameNumbersString, "|")
	winningNumbersString := gameNumbersStringArray[0]
	givenNumbersString := gameNumbersStringArray[1]

	// Handles Winning Numbers
	winningNumbersString = strings.TrimPrefix(winningNumbersString, " ")
	winningNumbersArray := strings.Split(winningNumbersString, " ")
	var winningNumbers []int

	for _, winningNumberString := range winningNumbersArray {
		if winningNumberString == "" || winningNumberString == " " {
			{
				continue
			}
		}

		winningNumber, err := strconv.Atoi(winningNumberString)
		if err != nil {
			fmt.Println("Error: Winning number could not be parsed")
			return card, err
		}
		winningNumbers = append(winningNumbers, winningNumber)
	}

	// Handles Given Numbers
	givenNumbersString = strings.TrimPrefix(givenNumbersString, " ")
	givenNumbersArray := strings.Split(givenNumbersString, " ")
	var givenNumbers []int
	for _, givenNumberString := range givenNumbersArray {
		if givenNumberString == "" || givenNumberString == " " {
			{
				continue
			}
		}

		givenNumber, err := strconv.Atoi(givenNumberString)
		if err != nil {
			fmt.Println("Error: Given number could not be parsed")
			return card, err
		}
		givenNumbers = append(givenNumbers, givenNumber)
	}

	card.cardNumber = cardNumber
	card.winningNumbers = winningNumbers
	card.givenNumbers = givenNumbers

	return card, nil
}

func parseCardStringArray(cardStringArray []string) ([]Card, error) {
	var cards []Card
	for _, cardString := range cardStringArray {
		card, err := parseCardString(cardString)
		if err != nil {
			fmt.Println("Error: Card string could not be parsed")
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

type CardCounter struct {
	originalWinningNumbers []int
	copiedCardNumbers      []int
}

func (cardCounter CardCounter) countTotalCards() int {
	return len(cardCounter.originalWinningNumbers) + len(cardCounter.copiedCardNumbers)
}

func getCardCounter(cards []Card, originalNumbers []int, copiedNumbers []int) CardCounter {
	var cardCounter CardCounter

	var tempOriginalNumbers []int
	var tempCopiedNumbers []int

	for _, card := range cards {
		wonCopiedNumbers := card.getWinningCopiedCardNumbers()

		if slices.Contains(tempOriginalNumbers, card.cardNumber) == false {
			tempOriginalNumbers = append(tempOriginalNumbers, card.cardNumber)
		}
		for _, wonCopiedNumber := range wonCopiedNumbers {
			tempCopiedNumbers = append(tempCopiedNumbers, wonCopiedNumber)
		}

	}

	var stackedCards []Card
	for _, card := range cards {
		if slices.Contains(tempCopiedNumbers, card.cardNumber) == true {

			stackedCards = append(stackedCards, card)
		}
	}
	// recursive call through won cards
	if len(stackedCards) > 0 {
		tempCardCounter := getCardCounter(stackedCards, tempOriginalNumbers, tempCopiedNumbers)

		tempCopiedNumbers = append(tempCopiedNumbers, tempCardCounter.copiedCardNumbers...)
		for _, tempOriginalNumber := range tempCardCounter.originalWinningNumbers {
			if slices.Contains(tempOriginalNumbers, tempOriginalNumber) == false {
				tempOriginalNumbers = append(tempOriginalNumbers, tempOriginalNumber)
			}
		}
	}

	cardCounter.originalWinningNumbers = tempOriginalNumbers
	cardCounter.copiedCardNumbers = tempCopiedNumbers

	return cardCounter
}

func main() {

	file_path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		return
	}
	file_path += "/day4_puzzle.txt"

	lines, err := readFileToList(file_path)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	cards, err := parseCardStringArray(lines)
	if err != nil {
		fmt.Println("Error parsing cards")
		return
	}

	var winningPoints int
	for _, card := range cards {
		winningPoints += card.getWinningPoints()
	}

	fmt.Println("puzzle 1:", winningPoints)

	// puzzle 2

	cardCount := getCardCounter(cards, []int{}, []int{})
	actualCardCount := cardCount.countTotalCards()
	fmt.Println("puzzle 2:", actualCardCount)

}
