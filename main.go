package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/and1x/pomodoro/logtask"
)

var Months = [12]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func isValidMonth(str string) bool {
	for _, v := range Months {
		if str == v {
			return true
		}
	}
	return false
}

func getTaskList(taskList *logtask.TaskList, fileName string, flag int) error {
	f, err := os.OpenFile(fileName, flag, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return taskList.Load(f)
}

func main() {
	timeFlag := flag.Int("d", 25, "set the duration of one pomodoro round (in Minutes).")
	nameFlag := flag.String("t", "not specified", "set the task you gonna do")
	printFlag := flag.String("l", "", "List Stats about pomodoro sessions.\nOptions:  'd' (daily), 'm' (monthly)") // ## Options: d, m, all, January, February ...
	flag.Parse()

	// ## every new month a new file to save pomodoro sessions gets created
	fileName := fmt.Sprintf(".%v.json", time.Now().Month())
	taskList := &logtask.TaskList{}

	switch {
	case *printFlag == "":
		if err := getTaskList(taskList, fileName, os.O_RDONLY|os.O_CREATE); err != nil {
			log.Println(err)
			os.Exit(1)
		}

		startTime := time.Now()
		runTimer(*timeFlag)

		task := &logtask.Task{
			Name:       *nameFlag,
			Duration:   *timeFlag,
			StartedAt:  startTime,
			FinishedAt: time.Now(),
		}

		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		taskList.Add(task)
		taskList.Save(f)

		// Add Break Timer - Break is 1/5 of the Pomodoro time
		fmt.Printf("do you want a break for %dmin? y/n \n", *timeFlag/5)
		if makeBreak() {
			runTimer(*timeFlag / 5)
			os.Exit(0)
		}
	case *printFlag == "d" || *printFlag == "m":
		if err := getTaskList(taskList, fileName, os.O_RDONLY|os.O_CREATE); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		taskList.PrintStats(*printFlag)
		os.Exit(0)
	case isValidMonth(*printFlag):
		fileName = fmt.Sprintf(".%v.json", *printFlag)
		if err := getTaskList(taskList, fileName, os.O_RDONLY); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		taskList.PrintStats(*printFlag)
		os.Exit(0)

	case *printFlag == "all":
		for i := 0; i < 12; i++ {
			fileName = fmt.Sprintf(".%v.json", Months[i])

			if err := getTaskList(taskList, fileName, os.O_RDONLY); err != nil {
				if !errors.Is(err, os.ErrNotExist) {
					log.Println(err)
					os.Exit(1)
				}
			}
		}
		taskList.PrintStats(*printFlag)
		os.Exit(0)
	default:
		fmt.Println("Invalid Flag/Option\nMore info with: pom -h or --help")

	}
}
