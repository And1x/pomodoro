/*
Pomodoro timer
Set the time with flag at end of command - in Minutes
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"pomodoro/helpers"
	"pomodoro/sound"
	"strings"
	"time"
)

func main() {
	// check if logPhases exists - if not create it
	if helpers.CreateLogFileIfNotExsiting() {
		fmt.Println("created new log File")
	}

	setFrames := createFrames("█", 60) // 🟩 ⬛ ▀▄ █
	//fmt.Println(setFrames)

	// get timeframe and task from user - if empty set default settings
	var timeframe int
	timeframe, task := getUserSettings()

	timeframe *= 60 //*60 hence we calculate in seconds - timeframe gets time in minutes
	pause := timeframe / 5

	//tik := time.NewTicker(1 * time.Millisecond)
	tik := time.NewTicker(1 * time.Second)

	// i is our control variable which represents the seconds elapsed, j is the animation counter
	i, j := 0, 0

	// startTime := fmt.Sprintf("Start: %s", time.Now())
	startTime := time.Now()

	// phaseEnd is used to control how often the first if statement should be used
	var phaseEnd int

	// clear screen and jump at start of screen
	fmt.Print("\033[H\033[2J")

timeLoop:
	for range tik.C {
		i++
		j++

		switch {
		case i%60 == 0:
			fmt.Print("\033[H\033[2J") // clear whole screen
			j = 0
		case i >= timeframe && phaseEnd == 0:
			phaseEnd++
			//## tik.Stop()

			// fmt.Println("start", startTime)
			endTime := time.Now()
			// fmt.Println("Ende ", endTime)
			sound.PlaySound() // play sound at the end
			helpers.Logger(startTime, endTime, timeframe, task)

			fmt.Printf("do you want a break for %dsec? y/n \n", pause)
			if makeBreak() { // check if user wants a break - if yes continue to tik else stop here
				continue
			} else {
				tik.Stop()
				break timeLoop
			}
			//## break timeLoop
		case i >= timeframe+pause && phaseEnd == 1:
			tik.Stop()
			sound.PlaySound() // play sound at the end
			break timeLoop

		default:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("\033[39m%v\n\033[31m%d\n\n", setFrames[j], i)

		}
	}

}

// getUserSettings gets the timeframe + task as cli- arg eg. // eg. >$ go run main.go 15 'write stuff'
// flags without "-" or "--"
func getUserSettings() (int, string) {

	d := flag.Int("d", 25, "set the duration of one pomodoro round")
	t := flag.String("t", "not specified", "set the task you gonna do")
	flag.Parse()
	return *d, *t //flag.Arg(0), flag.Arg(1)
}

// createFrames creates the Frames to loop trough to get an animation... Probably unnecessarry - could be improved later on.
func createFrames(symbol string, length int) []string {

	animationFrames := []string{}
	for i := 0; i <= length; i++ {
		step := strings.Repeat(symbol, i)
		animationFrames = append(animationFrames, step)
	}
	return animationFrames
}

// makeBreak ask user to make a break if he types y/Y then true -- else: false
func makeBreak() bool {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	if char == 'y' || char == 'Y' {
		return true
	} else if char == 'n' || char == 'N' {
		return false
	}
	return false
}
