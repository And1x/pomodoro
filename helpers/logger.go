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
)

// logger writes start and end of a pomodoro phase in file logPhases
func Logger(start, end time.Time, timeframe int, task string) {

	f, err := os.OpenFile("logPhases.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("cant open file bro", err)
	}
	defer f.Close()

	//##########################################
	// read the file to get the Runs - and add +1 to add current run
	buffer := make([]byte, 30)
	n, err := f.Read(buffer)
	if err != nil {
		fmt.Println("cant read string bro", err, n)
	}
	//fmt.Println("BUFFERRR::", string(buffer))

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

	statsSep := strings.Repeat("═", 55)

	//##########################################
	// write the Runs to file
	f.Seek(0, io.SeekStart)
	addTotalRuns := "Runs: " + strconv.Itoa(runsInt) + " \nTotal time: " + strconv.Itoa(total) + " min\n" + statsSep + "\n" + statsSep + "\n" // " \n" whitespace before newline needed so split function works as intended
	_, err = f.Write([]byte(addTotalRuns))
	if err != nil {
		fmt.Println("cannot writeFile 2 runs", err)
	}

	//##########################################
	// append the logs the file
	f.Seek(0, io.SeekEnd)

	// get length to create a separation line - with the exact same length as the logs
	var length int
	if len(start.String()) > len(end.String()) {
		length = len(start.String())

	} else {
		length = len(end.String())
	}
	sepLine := strings.Repeat("═", length) //▞ ═

	// get time difference and write at end of file with separator line
	difference := end.Sub(start)

	// build each line to add into log file
	wholeLog := []string{
		start.String(),
		end.String(),
		"Time passed: " + difference.String(),
		"time set: " + strconv.Itoa(timeframe/60),
		"Task: " + task,
		sepLine,
	}
	// create []byte with newlines
	var wholeLog2 bytes.Buffer
	for _, val := range wholeLog {
		_, _ = wholeLog2.WriteString(val + "\n")
	}

	// write into the log file
	//_, err = f.Write([]byte(start.String() + "\n" + end.String() + "\n" + "Time passed: " + difference.String() + "\n" + "Time set: " + strconv.Itoa(timeframe/60) + "\n" + "Task: " + task + "\n" + sepLine + "\n"))
	_, err = f.Write(wholeLog2.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
