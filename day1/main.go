package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	cal int
}

type elfArr []elf

func (e elfArr) Len() int {
	return len(e)
}
func (e elfArr) Less(i, j int) bool {
	return e[i].cal < e[j].cal
}

func (e elfArr) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	file, err := os.Open("data")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(file)

	elfs := make([]elf, 0, 2000)
	var current elf

loop:
	for {
		line, _, err := buf.ReadLine()

		switch {
		case err == nil:
		case errors.Is(err, io.EOF):
			break loop
		default:
			panic(err)
		}

		if len(line) == 0 {
			elfs = append(elfs, current)
			current = elf{}
			continue
		}

		cal, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}

		current.cal += cal
	}

	sort.Sort(elfArr(elfs))

	fmt.Printf("%+v\n", elfs[len(elfs)-1])
}
