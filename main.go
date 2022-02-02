package main

import (
	"fmt"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func main() {
	fmt.Println("asd")

	n := 9
	row := 7

	position := Position{
		X: 2,
		Y: 3,
	}

	beforePosition := Position{
		X: 2,
		Y: 2,
	}

	alert := ""

	for i := row; i >= 0; i-- {
		fmt.Print(i)
		k := 0
		for j := 1; j <= n; j++ {
			k = k + 1
			if i == 0 || i == row {
				fmt.Print(k)
			} else {
				print := "#"
				if j == n {
					print = strconv.Itoa(k)
				} else {

					if i == 2 && (j == 2 || j == 4 || j == 5 || j == 6 || j == 7) {
						print = "."
					}

					if i == 3 && (j == 2 || j == 3 || j == 4 || j == 6) {
						print = "."
					}

					if i == 4 && (j == 2 || j == 6 || j == 7) {
						print = "."
					}

					if i == 5 && (j != 1 && j != 8) {
						print = "."
					}

					// fmt.Print("------")
					// fmt.Print(print)
					// fmt.Print("------")

					if print == "#" {
						alert = "tidak bisa"
						beforePosition.X = position.X
						beforePosition.Y = position.Y

					} else {
						alert = ""
					}
					if position.X == i && position.Y == j {
						print = "X"
					}

				}

				fmt.Print(print)
			}

		}
		fmt.Println("")
	}

	fmt.Println(alert)

}
