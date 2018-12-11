package day9

import (
	"strings"
	"regexp"
	"log"
	"strconv"
	"fmt"
	"errors"
)

func getWinner(scoreBoard map[int]int) (elf int, score int) {
	for k, v := range scoreBoard {
		if v < score {
			continue
		}

		elf = k
		score = v
	}

	return elf, score
}

func parseInput(input string) (int, int, error) {
	inputSlice := strings.Split(input, ";")
	if len(inputSlice) != 2 {
		return 0, 0, errors.New("incorrect input format")
	}

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	playersString := reg.ReplaceAllString(inputSlice[0], "")
	pointsString := reg.ReplaceAllString(inputSlice[1], "")

	players, err := strconv.Atoi(playersString)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse players. error - [%s]", err.Error())
	}

	lastMarbleScore, err := strconv.Atoi(pointsString)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse last marble score. error - [%s]", err.Error())
	}

	return players, lastMarbleScore, nil
}

func createScoreBoard(amountOfPlayers int) map[int]int {
	scoreBoard := make(map[int]int)

	for i := 1; i <= amountOfPlayers; i++ {
		scoreBoard[i] = 0
	}

	return scoreBoard
}

func addScore(scoreBoard map[int]int, players, turn, score int) {
	player := getPlayer(players, turn)

	scoreBoard[player] = scoreBoard[player] + score + turn
}

func getPlayer(players, turn int) int {
	player := turn - ((turn / players) * players)

	if player == 0 {
		return players
	}
	return player
}
