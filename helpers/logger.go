/*
ToDo:: does any use of logPhases.txt needs  seperate os.Open()  ???
*/

package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// logger writes start and end of a pomodoro phase in file logPhases
func Logger(start, end time.Time) {

	//##########################################
	// read the file to get the Runs - and add +1 to add it in next turn
	readFile, err := os.Open("logPhases.txt")
	if err != nil {
		fmt.Println("cant open file bro", err)
	}
	defer readFile.Close()
	readRuns := bufio.NewReader(readFile)
	res, err := readRuns.ReadString('\n')
	if err != nil {
		fmt.Println("cant read string bro", err)
	}
	// get number of runs from string and add +1
	runs := strings.TrimPrefix(res, "Runs: ")
	runs = strings.TrimSuffix(runs, "\n")
	runsInt, err := strconv.Atoi(runs)
	if err != nil {
		fmt.Println("cant strconv", err)
	}
	runs = strconv.Itoa(runsInt + 1)
	total, _ := strconv.Atoi(runs)
	total *= 25

	//##########################################
	// write the Runs to file
	writeFile, err := os.OpenFile("logPhases.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("cant open file bro", err)
	}
	_, err = writeFile.Write([]byte("Runs: " + runs + "\n"))
	if err != nil {
		fmt.Println("cannot writeFile 2 runs", err)
	}
	_, err = writeFile.Write([]byte("Total time: " + strconv.Itoa(total) + "\n" + strings.Repeat("_-", 30) + "\n" + strings.Repeat("_-", 30) + "\n"))
	if err != nil {
		fmt.Println("cannot writeFile 2 total time", err)
	}
	defer writeFile.Close()

	//##########################################
	// append the logs the file
	file, err := os.OpenFile("logPhases.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
	_, err = file.Write([]byte(start.String() + "\n" + end.String() + "\n" + "Time passed: " + difference.String() + "\n" + sepLine + "\n"))
	if err != nil {
		log.Fatal(err)
	}

}
