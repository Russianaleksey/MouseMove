package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Click shift + ctrl + q to quit")
	go event()
	for {
		secondsMin := 10
		secondsMax := 50

		lower := -300
		upper := 300
		fTime := (rand.Intn(secondsMax-secondsMin) + secondsMin)
		robotgo.MoveSmoothRelative(RandInt(lower, upper), RandInt(lower, upper))
		time.Sleep(time.Duration(fTime) * time.Second)
	}
}

func event() {
	quit := robotgo.AddEvents("q", "ctrl", "shift")
	if quit {
		os.Exit(0)
	}
}

func RandInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	rng := upper - lower
	return rand.Intn(rng) + lower
}
