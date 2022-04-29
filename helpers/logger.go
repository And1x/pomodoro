/*

 */

package helpers

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// logger writes start and end of a pomodoro phase in file logPhases
func Logger(start, end time.Time) {

	f, err := os.OpenFile("logPhases.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("cant open file bro", err)
	}
	defer f.Close()

	//##########################################
	// read the file to get the Runs - and add +1 to add it in next turn
	buffer := make([]byte, 10)
	n, err := f.Read(buffer)
	if err != nil {
		fmt.Println("cant read string bro", err, n)
	}
	fmt.Println("BUFFERRR::", string(buffer))
	// get number of runs from string and add +1
	runs := strings.TrimPrefix(string(buffer), "Runs: ")
	runs = trimRuns(runs)
	runsInt, err := strconv.Atoi(runs)
	if err != nil {
		fmt.Println("cant strconv", err)
	}
	runs = strconv.Itoa(runsInt + 1)
	total, _ := strconv.Atoi(runs)
	total *= 25

	//##########################################
	// write the Runs to file
	f.Seek(0, io.SeekStart)
	addTotalRuns := "Runs: " + runs + "\nTotal time: " + strconv.Itoa(total) + "\n" + strings.Repeat("_-", 30) + "\n"
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
	sepLine := strings.Repeat("#", length)

	// get time difference and write at end of file with separator line
	difference := end.Sub(start)
	_, err = f.Write([]byte(start.String() + "\n" + end.String() + "\n" + "Time passed: " + difference.String() + "\n" + sepLine + "\n"))
	if err != nil {
		log.Fatal(err)
	}

}

func trimRuns(str string) string {
	for i, val := range str {
		if !unicode.IsNumber(val) {
			return str[:i]
		}
	}
	return str
}
