package day_3

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const groupLen = 3

func Solve2() {
	path := fmt.Sprintf("day_3/input.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var total int
	var elvesGroup [][]byte
	for sc.Scan() {
		line := sc.Bytes()
		elvesGroup = append(elvesGroup, line)

		if len(elvesGroup) == groupLen {
			badge, err := getBadges(elvesGroup)
			if err != nil {
				log.Fatalln(err)
			}
			if val, ok := priorities[badge]; ok {
				total = total + val
			}

			elvesGroup = [][]byte{}
		}
	}
	log.Printf("Solution for day 3, task 2 is %d \n", total)
}

func getBadges(group [][]byte) (string, error) {
	if len(group) != groupLen {
		return "", errors.New(fmt.Sprintf("Group must consist of %d memeber", groupLen))
	}
	comp1 := group[0]
	comp2 := group[1]
	comp3 := group[2]

	for _, elem := range comp1 {
		if contains(comp2, elem) {
			if contains(comp3, elem) {
				return string(elem), nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("No matching elements in %v and %v and %v", comp1, comp2, comp3))
}
