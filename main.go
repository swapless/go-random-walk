package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var choices [4][2]int = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type RandomWalk struct {
	length     int
	iterations int
}

func newRandomWalk(l int, iter int) *RandomWalk {
	rw := RandomWalk{
		length:     l,
		iterations: iter,
	}
	return &rw
}

// for starting position (0,0) make a step, till step reaches length
func (rw RandomWalk) Walk() {
	// repeat for i iterations
	TotalWalkLength := 0
	for i := 0; i < rw.iterations; i++ {
		// for every iteration
		x, y := 0, 0
		for step := 0; step < rw.length; step++ {
			fmt.Println(x, y, choices[rand.Intn(len(choices))])
			choice := choices[rand.Intn(len(choices))]
			x += choice[0]
			y += choice[1]
		}

		distance := math.Max(math.Abs(float64(0-x)), math.Abs(float64(0-y)))

		TotalWalkLength += int(distance)

	}

	avgWalkLength := TotalWalkLength / rw.iterations

	fmt.Println(avgWalkLength)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("Random Walk Tester")
	randomWalk := newRandomWalk(1000, 100000)
	fmt.Println("Random Walk Length : ", randomWalk.length)
	randomWalk.Walk()
}
