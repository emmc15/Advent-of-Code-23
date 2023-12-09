package second

import (
	"os"
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

func TestParseElfGameFromString(t *testing.T) {
	game := parseElfGameString("Game 12: 1 green, 2 blue, 13 red, 2 blue, 3 green, 4 green, 14 red")
	if game.game != 12 {
		t.Error("Number not parsed correctly")
	}
	var expectedCubes []int
	expectedCubes = append(expectedCubes, 2)
	expectedCubes = append(expectedCubes, 2)
	if len(game.blueCubes) != len(expectedCubes) {
		t.Error("Number of blue cubes not parsed correctly")
	}
	for i, cube := range game.blueCubes {
		if cube != expectedCubes[i] {
			t.Error("Blue cube not parsed correctly")
		}
	}

	if game.getGreenCubeCount() != 8 {
		t.Error("Blue cube count not calculated correctly")
	}
	if game.getRedCubeCount() != 27 {
		t.Error("Blue cube count not calculated correctly")
	}
	if game.getBlueCubeCount() != 4 {
		t.Error("Blue cube count not calculated correctly")
	}

}

// func TestFiltetElfGamesByCubeCounts(t, *testing.T) {
//
// 	// puzzle 1
// 	// filter the games by cube counts
// 	filteredGames := filterElfGamesByCubeCounts(games, 14, 12, 13)
// 	if len(filteredGames) != 1 {
// 		t.Error("Filtered games not calculated correctly")
// 	}
// 	if filteredGames[0].game != 1 {
// 		t.Error("Filtered games not calculated correctly")
// 	}
//
// 	// puzzle 2
// 	// filter the games by cube counts
// 	filteredGames = filterElfGamesByCubeCounts(games, 14, 12, 13)
// 	if len(filteredGames) != 1 {
// 		t.Error("Filtered games not calculated correctly")
// 	}
// 	if filteredGames[0].game != 1 {
// 		t.Error("Filtered games not calculated correctly")
// 	}
// }

func TestMain(t *testing.T) {
	main()
}
