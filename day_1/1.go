package day_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solve1() {
	path := fmt.Sprintf("day_1/input.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)

	max := 0
	current := 0
	for sc.Scan() {
		if current > max {
			max = current
		}
		val := sc.Text()
		if val == "" {
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
	log.Printf("Solution for day 1, task 1 is %d\n", max)
}
