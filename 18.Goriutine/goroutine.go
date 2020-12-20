package main

import (
	"fmt"
	"math"
)

const min = 2

func main() {
	primeNum := []int{}

	// * The Channel is how to communicate between main() function and GoRoutine.
	// * It's How to Communicate by putting Channel as a Parameter instead of returning Something.
	channel := make(chan int)

	// * This is the GoRoutine.
	// * The GoRoutine makes process asynchronous.
	// * but GoRoutine is valid for program is running.
	// * And the main() function doesn't wait the GoRoutine.
	// * So if main() function has nothing to do except GoRoutine, the program ends.

	go generatePrimeNumber(primeNum, min, 10000000, channel)
	go generatePrimeNumber(primeNum, min, 10000000, channel)

	result := <-channel // * but main() function waits until the Channel sends message.

	fmt.Println(result)
	fmt.Println(<-channel) // * It can also use it like this.

	// time.Sleep(time.Second * 25)
}

func generatePrimeNumber(primeNum []int, start, endPoint int, channel chan int) {
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
		// time.Sleep(time.Second)
	}
	channel <- len(primeNum) //* Usage
	// return len(primeNum)
}
