package helper

import (
	"fmt"
	"log"
	"time"
)

// PrintTimePretty formats a specified time & unit like h,m,s  into a time.String in the Format "72h3m0.5s"
func PrintTimePretty(amnt int, unit string) string {
	d, err := time.ParseDuration(fmt.Sprintf("%d%s", amnt, unit))
	if err != nil {
		log.Println(err)
	}
	return d.String()
}
