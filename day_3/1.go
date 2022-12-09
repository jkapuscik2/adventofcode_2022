package day_3

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func Solve1() {
	path := fmt.Sprintf("day_3/input.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var total int
	for sc.Scan() {
		line := sc.Bytes()
		duplicate, err := getDuplicated(line)
		if err != nil {
			log.Fatalln(err)
		}
		if val, ok := priorities[duplicate]; ok {
			total = total + val
		}
	}
	log.Printf("Solution for day 3, task 1 is %d \n", total)
}

func getDuplicated(line []byte) (string, error) {
	comp1 := line[:len(line)/2]
	comp2 := line[len(line)/2:]

	for _, elem := range comp1 {
		if contains(comp2, elem) {
			return string(elem), nil
		}
	}

	return "", errors.New(fmt.Sprintf("No matching elements in %v and %v", comp1, comp2))
}

func contains(list []byte, elem byte) bool {
	for _, l := range list {
		if l == elem {
			return true
		}
	}
	return false
}
