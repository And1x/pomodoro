/*
Pomodoro timer
Set the time with flag at end of command - in Minutes
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"pomodoro/helpers"
	"pomodoro/sound"
	"strings"
	"time"
)

const taskFile = ".taskList.json" // todo: add weekly name

func main() {

	d := flag.Int("d", 25, "set the duration of one pomodoro round")
	taskName := flag.String("t", "not specified", "set the task you gonna do")
	printStats := flag.String("print", "n", "Print Stats about pomodoro sessions. -d for daily -m for montly")
	flag.Parse()

	f, err := os.OpenFile(taskFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	taskList := &helpers.TaskList{}

	if err := taskList.Load(f); err != nil {

		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}

	switch {
	// ## Print Tasklist - Daily or Monthly
	case *printStats == "d" || *printStats == "m":
		taskList.PrintStats(*printStats)
		os.Exit(0)
		// ## Exit if entered Task is to long.
	case len(*taskName) > 100:
		fmt.Println("Task entered is to long - please stay under the 100 character Limit.")
		os.Exit(1)
		// ## RUN THE POMODORO TIMER
	default:
		startTime := time.Now()
		runTimer(*d)

		task := &helpers.Task{
			Name:       *taskName,
			Duration:   *d,
			StartedAt:  startTime,
			FinishedAt: time.Now(),
		}

		f, err := os.OpenFile(taskFile, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		taskList.Add(task)
		taskList.Save(f)

		// Add Break Timer - Break is 1/5 of the Pomodoro time
		fmt.Printf("do you want a break for %dmin? y/n \n", *d/5)
		if makeBreak() {
			runTimer(*d / 5)
			os.Exit(0)
		}
	}
}

func runTimer(duration int) {

	tik := time.NewTicker(1 * time.Millisecond)
	i, j := 0, 0

	for range tik.C {
		j++
		i++

		switch {
		case i > (duration * 60):
			tik.Stop()
			sound.PlaySound()
			fmt.Println()
			return
		case i%60 == 0:
			j = 0
			fmt.Print("\x1B[2K") // [2K erases complete line
		}

		// 🟩 ⬛ ▀▄ █
		fmt.Printf("\r%s%s[%s]", strings.Repeat("█", j), strings.Repeat(" ", 60-j), printElapsedTime(i)) // \r brings cursor to start of line
	}
}

// printElapsedTime takes i as seconds to convert it into a time.String in the Format "72h3m0.5s"
func printElapsedTime(sec int) string {
	d, err := time.ParseDuration(fmt.Sprintf("%ds", sec))
	if err != nil {
		log.Println(err)
	}
	return d.String()
}

// makeBreak prompts if break is wanted. Pomodoro breaks are 1/5 of time per Phase.
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
