/*

 */

package helpers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// logger writes start and end of a pomodoro phase in file logPhases
func Logger(start, end time.Time, timeframe int, task string) {

	f, err := os.OpenFile("logPhases.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("cant open file bro", err)
		return
	}
	defer f.Close()

	//##########################################
	// read the file to get the Runs - and add +1 to add current run
	f.Seek(0, io.SeekStart)

	buffer := make([]byte, 30)
	n, err := f.Read(buffer)
	if err != nil {
		fmt.Println("cant read string bro", err, n)
	}

	dirtyRunsTotal := strings.Fields(string(buffer))

	// get number of runs from string and add +1
	runs := dirtyRunsTotal[1]
	runsInt, err := strconv.Atoi(runs)
	if err != nil {
		fmt.Println("cant strconv", err)
	}
	runsInt++ // add +1 for current run

	total, _ := strconv.Atoi(dirtyRunsTotal[4])
	total += timeframe / 60

	//statsSep := strings.Repeat("═", 55)

	//##########################################
	// write the Runs to file

	f.Seek(0, io.SeekStart)

	_, err = f.Write([]byte(LogHead(runsInt, total)))
	if err != nil {
		fmt.Println("cannot writeFile 2 runs", err)
	}

	//##########################################
	// append the logs the file

	f.Seek(0, io.SeekEnd)

	// Creation of log parts
	startS := start.Format("02-01-2006 | 15:04:05")
	endS := end.Format("15:04:05")
	// calculate exact time difference - could be longer as set time due to breaks
	differenceS := fmt.Sprintf(" | Passed: %.2f", end.Sub(start).Minutes())
	setS := fmt.Sprintf(" | Set: %v", strconv.Itoa(timeframe/60))
	stats := startS + "  ===>  " + endS + differenceS + setS
	maxLogLength := utf8.RuneCountInString(stats)

	task = SplitTask(task, maxLogLength)

	// build each line to add into log file
	wholeLog := []string{
		//start.String(),
		stats,
		task,
		strings.Repeat("═", maxLogLength), //▞ ═
	}
	// create []byte with newlines
	var wholeLog2 bytes.Buffer
	for _, val := range wholeLog {
		_, _ = wholeLog2.WriteString(val + "\n")
	}

	_, err = f.Write(wholeLog2.Bytes())
	// write into the log file
	if err != nil {
		log.Fatal(err)
	}

}

// Split tasks splits the task with a newline \n at the length to fit into logFrame
func SplitTask(task string, length int) string {

	var b strings.Builder
	stask := strings.Fields("Task: " + task)

	counter := 0

	for _, val := range stask {
		counter += utf8.RuneCountInString(val + " ") // add whitespace to counter aswell
		if counter > length {
			counter = 0
			counter += utf8.RuneCountInString(val + " ") // count the string which get's into the newline too
			b.WriteRune('\n')

		}
		b.WriteString(val + " ")

	}
	return b.String()
}

// Creates new Log File if none is existing
func CreateLogFileIfNotExsiting() bool {
	f, err := os.Open("logPhases.txt")
	if os.IsNotExist(err) {
		f, err := os.Create("logPhases.txt")
		if err != nil {
			log.Println(err)
		}
		_, err = f.Write([]byte(LogHead(0, 0)))
		if err != nil {
			log.Println(err)
		}
		f.Close()
		return true
	} else {
		f.Close()
		return false
	}
}

// Build Default string for Log Head (total + runs)
func LogHead(runs int, total int) string {
	var statsSep string = strings.Repeat("═", 62)
	return fmt.Sprintf("Runs: %v \nTotal time: %v min\n%s\n%s\n", runs, total, statsSep, statsSep)
}
