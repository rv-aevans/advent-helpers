package helpers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

type Input struct {
	ColOne   string `csv:"column_one"`
	ColTwo   string `csv:"column_two"`
	ColThree string `csv:"column_three"`
	ColFour  string `csv:"column_four"`
}

func GetAlphaMap() map[string]int {
	return map[string]int{
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
}

func GetLargestInt(ints []int) (largest int) {
	for _, i := range ints {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func MakeInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("** STR TO INT ERR **")
		fmt.Println(err)
		panic(err)
	}
	return n
}

func UnmarshalCsv() (output []Input) {
	f, err := os.Open("input_emile.csv")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %w", err))
	}
	defer f.Close()

	inputStruct := []Input{}

	if err := gocsv.UnmarshalFile(f, &inputStruct); err != nil {
		fmt.Println(fmt.Errorf("error unmarshalling file: %w", err))
	}

	return inputStruct
}
