package main

import (
	"fmt"
	"math"
	"time"
)

const min = 2

func main() {
	primeNum := []int{}

	go generatePrimeNumber(primeNum, min, 10000000)
	generatePrimeNumber(primeNum, min, 10000000)
}

func generatePrimeNumber(primeNum []int, start, endPoint int) {
	isPrime := true
	end := start + endPoint
	for i := start; i < end; i++ {
		for j := min; j < int(math.Sqrt(float64(end))); j++ {
			if i != j && i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeNum = append(primeNum, i)
		}
		isPrime = true
	}
	fmt.Println(len(primeNum))
	time.Sleep(time.Second)
}
