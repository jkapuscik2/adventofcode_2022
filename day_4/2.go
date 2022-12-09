package day_4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve2() {

	path := fmt.Sprintf("day_4/input.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var total int
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, elvesSeparator)

		if len(parts) != elvesNum {
			log.Fatalf("Invalid data provided %s", line)
		}
		range1 := convertRange(strings.Split(parts[0], rangeSeparator))
		range2 := convertRange(strings.Split(parts[1], rangeSeparator))
		if overlaps(range1, range2) {
			total = total + 1
		}
	}
	log.Printf("Solution for day 4, task 1 is %d \n", total)
}

func overlaps(range1 sectorRange, range2 sectorRange) bool {
	if range1.Start <= range2.End && range2.Start <= range1.End {
		return true
	}
	if range2.Start <= range1.End && range1.Start <= range2.End {
		return true
	}

	return false
}
