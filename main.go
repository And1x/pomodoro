/*
Pomodoro timer
Set the time with flag at end of command - in Minutes
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"pomodoro/helpers"
	"pomodoro/sound"
	"strconv"
	"time"
)

func main() {

	timeframe := setTimer() * 60 //*60 hence we calculate in seconds - timeframe gets time in minutes
	pause := timeframe / 5

	tik := time.NewTicker(1 * time.Millisecond)

	// i is our controll variable which represents the seconds elapsed
	i := 0

	// startTime := fmt.Sprintf("Start: %s", time.Now())
	startTime := time.Now()

	// phaseEnd is used to control how often the first if statement should be used
	var phaseEnd int

	for range tik.C {
		i++

		fmt.Printf("%s", "#")
		if i%5 == 0 {

			fmt.Printf("\x1b[2J") // clear whole screen
			fmt.Print("\033[H")   // jump at the start of the screen
			fmt.Printf("%vs: ", i)
		}

		if i >= timeframe && phaseEnd == 0 {
			phaseEnd += 1
			//tik.Stop()
			fmt.Println("xx", phaseEnd)
			fmt.Println("start", startTime)
			endTime := time.Now()
			fmt.Println("Ende ", endTime)
			sound.PlaySound() // play sound at the end
			helpers.Logger(startTime, endTime)
			//break
		} else if i >= (timeframe + pause) {
			tik.Stop()
			fmt.Println("")
			fmt.Println("startBREAK", startTime)
			endTime := time.Now()
			fmt.Println("EndeBREAK ", endTime)
			sound.PlaySound() // play sound at the end
			helpers.Logger(startTime, endTime)
			break
		}
	}

}

// setTimer gets the timeframe as cli- arg eg. // eg. "go run main.go 15"
// flags without "-" or "--"
func setTimer() (timeframe int) {
	flag.Parse()
	timeframe, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Println(err)
	}
	return
}
