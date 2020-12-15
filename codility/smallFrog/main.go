package main

import "fmt"

func Solution(X int, Y int, D int) int {
	distance := Y - X

	i := distance % D
	j := (distance + D) % D
	switch {
	case i == 0:
		return distance / D
	case i < j:
		return distance / D
	case i > j:
		return (distance + D) / D
	case i == j:
		return (distance + D) / D
	}

	fmt.Println(i, j)

	return distance
}

func main() {
	fmt.Println(Solution(10, 85, 30))
}
