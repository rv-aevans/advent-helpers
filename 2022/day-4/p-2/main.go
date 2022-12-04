package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution(unmarshalCsv()))
}

func solution(input []Input) (res int) {
	for _, v := range input {
		colOne := strings.Split(strings.TrimSpace(v.ColOne), "-")
		colTwo := strings.Split(strings.TrimSpace(v.ColTwo), "-")
		if makeInt(colOne[0]) <= makeInt(colTwo[0]) && makeInt(colOne[1]) >= makeInt(colTwo[0]) {
			res++
		} else if makeInt(colTwo[0]) <= makeInt(colOne[0]) && makeInt(colTwo[1]) >= makeInt(colOne[0]) {
			res++
		} else {
			continue
		}
	}
	return
}
