package main

import (
	"fmt"
)

type Position struct {
	X int
	Y int
}

func main() {

	position := Position{X: 2, Y: 2}
	_, _ = runApp(position)

	tempPosition := Position{X: position.X, Y: position.Y}

	callTyping(position, tempPosition)
}

func callTyping(position Position, temp Position) {
	typing := ""

	temp.X = position.X
	temp.Y = position.Y

	fmt.Print("type up/down/right/left ?")
	fmt.Scan(&typing)

	switch typing {
	case "up":
		position.X += 1
	case "down":
		position.X -= 1
	case "right":
		position.Y += 1
	case "left":
		position.Y -= 1
	}

	success, warning := runApp(position)

	if !success && warning {
		_, _ = runApp(temp)
		callTyping(temp, temp)
	} else if !success && !warning {
		callTyping(position, temp)
	} else {
		fmt.Println("game finished")
	}
}

func runApp(position Position) (success bool, warning bool) {
	fmt.Println("")
	var pagar []Position

	treasure := Position{X: 5, Y: 6}

	n := 9
	row := 7

	for i := row; i >= 0; i-- {
		fmt.Print("")
		k := 0
		for j := 1; j <= n; j++ {
			k = k + 1
			if i == 0 || i == row {
				fmt.Print("")
			} else {
				print := "#"
				if j == n {
					// print = strconv.Itoa(k)
					print = ""
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

					if print == "#" {
						pagar = append(pagar, Position{X: i, Y: j})

					}
					if position.X == i && position.Y == j {
						if position.X == treasure.X && position.Y == treasure.Y {
							print = "$"
						} else {
							print = "X"
						}
					}

				}

				fmt.Print(print)
			}

		}
		fmt.Println("")
	}

	alert := ""
	for _, pg := range pagar {
		if pg.X == position.X && pg.Y == position.Y {
			alert = "FORBIDDEN"
		}
	}

	fmt.Println(alert)

	if alert != "" {
		return false, true
	}

	if position.X == treasure.X && position.Y == treasure.Y {
		return true, false
	}

	return false, false
}
