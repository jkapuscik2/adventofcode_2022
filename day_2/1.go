package day_2

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rockPremium     = 1
	paperPremium    = 2
	scissorsPremium = 3

	win  = 6
	draw = 3
	lost = 0

	playerRock     = "X"
	playerPaper    = "Y"
	playerScissors = "Z"

	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"
	gameLen          = 2
)

func Solve1() {
	path := fmt.Sprintf("day_2/input.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var totalScore int

	for sc.Scan() {
		game := strings.Split(sc.Text(), " ")

		if len(game) != gameLen {
			log.Fatalln("Invalid input data")
		}
		gameScore, err := getGameScore(game[1], game[0])
		if err != nil {
			log.Fatalln(err)
		}
		totalScore = totalScore + gameScore

		premiumScore, err := getPremium(game[1])
		if err != nil {
			log.Fatalln(err)
		}
		totalScore = totalScore + premiumScore

	}

	log.Printf(fmt.Sprintf("Solution for day 2, task 1 is %d\n", totalScore))
}

func getGameScore(player string, opponent string) (int, error) {
	switch player {
	case playerRock:
		if opponent == opponentPaper {
			return lost, nil
		}
		if opponent == opponentScissors {
			return win, nil
		}
		if opponent == opponentRock {
			return draw, nil
		}
	case playerPaper:
		if opponent == opponentPaper {
			return draw, nil
		}
		if opponent == opponentScissors {
			return lost, nil
		}
		if opponent == opponentRock {
			return win, nil
		}
	case playerScissors:
		if opponent == opponentPaper {
			return win, nil
		}
		if opponent == opponentScissors {
			return draw, nil
		}
		if opponent == opponentRock {
			return lost, nil
		}
	default:
		return 0, errors.New(fmt.Sprintf("Unknow played by player %s", player))
	}

	return 0, errors.New(fmt.Sprintf("Unknow played by opponent %s", opponent))
}

func getPremium(picked string) (int, error) {
	switch picked {
	case playerRock:
		return rockPremium, nil
	case playerPaper:
		return paperPremium, nil
	case playerScissors:
		return scissorsPremium, nil
	}

	return 0, errors.New(fmt.Sprintf("Unknown played %s", picked))
}
