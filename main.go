package main

import (
	"fmt"
)

func seasonGenerator(month int) string {
	switch month {
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Fall"
	case 12, 1, 2:
		return "Winter"
	default:
	}
	return "none"
}

func main() {
	fmt.Println(seasonGenerator(10))
}
