package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func getOutcome(a, b int) int {
	switch {
	case a == b:
		return b + 3
	case (a == 1 && b == 3) || (a == 2 && b == 1) || (a == 3 && b == 2):
		return b
	case (a == 1 && b == 2) || (a == 2 && b == 3) || (a == 3 && b == 1):
		return b + 6
	}
	panic("")
}

func getOutcomePart2(a, b int) int {
	// 1 loose, 2 draw, 3 win
	switch {
	case b == 1:
		switch a {
		case 1:
			return 3
		case 2:
			return 1
		case 3:
			return 2
		}

	case b == 3:
		switch a {
		case 1:
			return 6 + 2
		case 2:
			return 6 + 3
		case 3:
			return 6 + 1
		}
	case b == 2:
		return 3 + a
	}

	panic("")
}

// enemy(A1 B2 C3) you(X Y Z)
func main() {
	file, err := os.Open("data")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(file)

	total := 0
	total2 := 0

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
		s := strings.Split(string(line), " ")
		var n, n2 int
		switch s[0] {
		case "A":
			n = 1
		case "B":
			n = 2
		case "C":
			n = 3
		}

		switch s[1] {
		case "X":
			n2 = 1
		case "Y":
			n2 = 2
		case "Z":
			n2 = 3
		}
		out := getOutcome(n, n2)
		out2 := getOutcomePart2(n, n2)
		fmt.Printf("%q%d %q%d - %d -%d\n", s[0], n, s[1], n2, out, out2)
		total = total + out
		total2 = total2 + out2
	}

	fmt.Println(total)
	fmt.Println(total2)
}
