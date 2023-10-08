package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Sleeper() {
	rand.Seed(time.Now().UnixNano())

	minDuration := 100 * time.Millisecond
	maxDuration := 400 * time.Millisecond
	randomDuration := minDuration + time.Duration(rand.Float64()*(float64(maxDuration-minDuration)))

	fmt.Printf("Random Duration: %s\n", randomDuration)
	time.Sleep(randomDuration)

	fmt.Println("Sleep timer finished. Continue with the program.")
}
