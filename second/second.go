package second

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ElfGame struct {
	game       int
	blueCubes  []int
	redCubes   []int
	greenCubes []int
}

func (re ElfGame) getBlueCubeCount() int {
	return sumIntArray(re.blueCubes)
}

func (re ElfGame) getRedCubeCount() int {
	return sumIntArray(re.redCubes)
}

func (re ElfGame) getGreenCubeCount() int {
	return sumIntArray(re.greenCubes)
}

func readFileToList(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func parseElfGameString(elfString string) *ElfGame {
	// check if string stars with Game
	if !strings.HasPrefix(elfString, "Game ") {
		return nil
	}

	// split the string into the game number and the cube string
	elfString = strings.TrimPrefix(elfString, "Game ")
	gameNameString := strings.Split(elfString, ":")[0]
	re := regexp.MustCompile("(\\d+)")
	gameNumberString := re.FindString(gameNameString)
	gameNumberInt, err := strconv.Atoi(gameNumberString)

	if err != nil {
		fmt.Println("Error converting string to int")
		return nil
	}

	cubeString := strings.Split(elfString, ":")[1]
	cubeArray := strings.Split(cubeString, ",")

	var blueCubes []int
	var redCubes []int
	var greenCubes []int

	for _, cube := range cubeArray {

		// regex out digit
		re := regexp.MustCompile("(\\d+)")
		countString := re.FindString(cube)
		countInt, err := strconv.Atoi(countString)
		if err != nil {
			fmt.Println("Error converting string to int")
			return nil
		}
		// fmt.Println(cube, countString, countInt)

		if strings.HasSuffix(cube, "blue") {

			blueCubes = append(blueCubes, countInt)

		} else if strings.HasSuffix(cube, "red") {
			countInt, err := strconv.Atoi(countString)
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil
			}
			redCubes = append(redCubes, countInt)

		} else if strings.HasSuffix(cube, "green") {
			countInt, err := strconv.Atoi(countString)
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil
			}
			greenCubes = append(greenCubes, countInt)
		}
	}

	// Add the parse values to the struct
	game := ElfGame{game: gameNumberInt, blueCubes: blueCubes, redCubes: redCubes, greenCubes: greenCubes}
	return &game

}

func parseElfGameStringArray(numberStringArray []string) ([]ElfGame, error) {
	var games []ElfGame
	for _, numberString := range numberStringArray {
		game := parseElfGameString(numberString)

		if game == nil {
			return nil, fmt.Errorf("Error parsing game string")
		}
		games = append(games, *game)

	}
	return games, nil
}

func sumIntArray(intArray []int) int {
	var sum int
	for _, number := range intArray {
		sum += number
	}
	return sum
}

func filterElfGamesByCubeCounts(games []ElfGame, blueCubeCount int, redCubeCount int, greenCubeCount int) []ElfGame {
	var filteredGames []ElfGame
	for _, game := range games {

		if game.getBlueCubeCount() <= blueCubeCount && game.getRedCubeCount() <= redCubeCount && game.getGreenCubeCount() <= greenCubeCount {
			filteredGames = append(filteredGames, game)
		}
	}
	return filteredGames
}

func main() {

	file_path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		return
	}
	file_path += "/day2_puzzle.txt"

	lines, err := readFileToList(file_path)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	var clenasedLines []string
	for _, line := range lines {
		clenasedLines = append(clenasedLines, strings.ReplaceAll(line, ";", ","))
	}

	games, err := parseElfGameStringArray(clenasedLines)
	if err != nil {
		fmt.Println("Error parsing game string")
		return
	}

	// puzzle 1
	// filter the games by cube counts
	filteredGames := filterElfGamesByCubeCounts(games, 14, 12, 13)

	// sum all filtered game numbers
	sum := 0
	for _, game := range filteredGames {
		sum += game.game
	}
	fmt.Printf("%+v", filteredGames)
	fmt.Println("puzzle 1:", sum)

}
