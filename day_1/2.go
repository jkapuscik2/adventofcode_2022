package day_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solve2() {
	path := fmt.Sprintf("day_1/input.txt")
	f, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)

	topScores := make([]int, 3)

	current := 0
	for sc.Scan() {
		val := sc.Text()
		if val == "" {
			topScores = updateTopScores(topScores, current)
			current = 0
		} else {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalln(err)
			}
			current = current + num
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Solution for day 1, task 1 is %d\n", sumScores(topScores))
}

func updateTopScores(scores []int, newScore int) []int {
	var lowest int
	var lowestIdx int

	for idx, score := range scores {
		if idx == 0 || score < lowest {
			lowest = score
			lowestIdx = idx
		}
	}

	if lowest < newScore {
		scores[lowestIdx] = newScore
	}

	return scores
}

func sumScores(scores []int) int {
	var sum int

	for _, v := range scores {
		sum = sum + v
	}

	return sum
}
