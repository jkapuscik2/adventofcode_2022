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
	playerLoose = "X"
	playerDraw  = "Y"
	playerWin   = "Z"
)

func Solve2() {
	path := fmt.Sprintf("day_2/input.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)

	var totalScore int

	for sc.Scan() {
		game := strings.Split(sc.Text(), " ")
		if len(game) != gameLen {
			log.Fatalln("Invalid input data")
		}

		gameScore, err := getOutcomeScore(game[1])
		if err != nil {
			log.Fatalln(err)
		}
		totalScore = totalScore + gameScore

		best, err := getSelectedType(game[0], game[1])
		if err != nil {
			log.Fatalln(err)
		}

		premiumScore, err := getPremium(best)
		if err != nil {
			log.Fatalln(err)
		}
		totalScore = totalScore + premiumScore
	}

	log.Printf("Solution for day 3, task 1 is %d \n", totalScore)
}

func getSelectedType(opponent string, outcome string) (string, error) {
	switch opponent {
	case opponentRock:
		if outcome == playerLoose {
			return playerScissors, nil
		}
		if outcome == playerWin {
			return playerPaper, nil
		}
		if outcome == playerDraw {
			return playerRock, nil
		}
	case opponentPaper:
		if outcome == playerDraw {
			return playerPaper, nil
		}
		if outcome == playerLoose {
			return playerRock, nil
		}
		if outcome == playerWin {
			return playerScissors, nil
		}
	case opponentScissors:
		if outcome == playerWin {
			return playerRock, nil
		}
		if outcome == playerDraw {
			return playerScissors, nil
		}
		if outcome == playerLoose {
			return playerPaper, nil
		}
	default:
		return "", errors.New(fmt.Sprintf("Unknow played by opponent %s", opponent))
	}

	return "", errors.New(fmt.Sprintf("Unknow game outcome %s", outcome))
}

func getOutcomeScore(outcome string) (int, error) {
	switch outcome {
	case playerLoose:
		return lost, nil
	case playerWin:
		return win, nil
	case playerDraw:
		return draw, nil
	}

	return 0, errors.New(fmt.Sprintf("Unknown game outcome %s", outcome))
}
