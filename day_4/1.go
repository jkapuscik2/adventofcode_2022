package day_4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	elvesSeparator = ","
	rangeSeparator = "-"
	rangeLen       = 2
	elvesNum       = 2
)

type sectorRange struct {
	Start int
	End   int
}

func Solve1() {

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
		if covers(range1, range2) {
			total = total + 1
		}
	}
	log.Printf("Solution for day 4, task 1 is %d \n", total)
}

func covers(range1 sectorRange, range2 sectorRange) bool {
	if range1.Start <= range2.Start && range1.End >= range2.End {
		return true
	}
	if range2.Start <= range1.Start && range2.End >= range1.End {
		return true
	}

	return false
}

func convertRange(r []string) sectorRange {
	if len(r) != rangeLen {
		log.Fatalf("Invalid data provided %v", r)
	}
	start, err := strconv.Atoi(r[0])
	if err != nil {
		log.Fatalln(err)
	}
	end, err := strconv.Atoi(r[1])
	if err != nil {
		log.Fatalln(err)
	}

	return sectorRange{start, end}
}
