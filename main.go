package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/and1x/pomodoro/logtask"
)

// every new month a new file to save pomodoro sessions gets created
var taskFile string = fmt.Sprintf(".%v.json", time.Now().Month())

func main() {

	duration := flag.Int("d", 25, "set the duration of one pomodoro round (in Minutes).")
	taskName := flag.String("t", "not specified", "set the task you gonna do")
	printStats := flag.String("print", "n", "Print Stats about pomodoro sessions. -d for daily -m for montly")
	flag.Parse()

	f, err := os.OpenFile(taskFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	taskList := &logtask.TaskList{}

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
		runTimer(*duration)

		task := &logtask.Task{
			Name:       *taskName,
			Duration:   *duration,
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
		fmt.Printf("do you want a break for %dmin? y/n \n", *duration/5)
		if makeBreak() {
			runTimer(*duration / 5)
			os.Exit(0)
		}
	}
}
