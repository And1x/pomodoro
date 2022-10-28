package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/and1x/pomodoro/helper"
	"github.com/and1x/pomodoro/sound"
)

func runTimer(duration int) {

	tik := time.NewTicker(1 * time.Millisecond) // todo change inProduction to seconds
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

		// print Timer
		fmt.Printf("\r%s%s[%s]", strings.Repeat("█", j), strings.Repeat(" ", 60-j), helper.PrintTimePretty(i, "s")) // \r brings cursor to start of line
	}
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
