package logtask

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/and1x/pomodoro/helper"
)

type Task struct {
	Name       string
	Duration   int
	StartedAt  time.Time
	FinishedAt time.Time
}

type TaskList []Task

// Add appends a new Task to the in Memory TaskList.
func (t *TaskList) Add(nTask *Task) {
	newTask := Task{
		Name:       nTask.Name,
		Duration:   nTask.Duration,
		StartedAt:  nTask.StartedAt,
		FinishedAt: nTask.FinishedAt,
	}

	*t = append(*t, newTask)
}

// Save writes the whole TaskList in JSON to a file.
func (t *TaskList) Save(w io.Writer) error {
	d, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}

// Load loads the whole TaskList into t (in Memory).
func (t *TaskList) Load(r io.Reader) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	// if file is empty (after create) - return here
	if len(buf) == 0 {
		return nil
	}

	var temp TaskList
	err = json.Unmarshal(buf, &temp)
	if err != nil {
		return err
	}
	*t = append(*t, temp...)
	return nil
}

// PrintStats pretty prints a table with Tasks done - interval Options are "d","m", "all", "January","February"...
func (t *TaskList) PrintStats(interval string) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Duration"},
			{Align: simpletable.AlignCenter, Text: "Date"},
			{Align: simpletable.AlignCenter, Text: "Start"},
			{Align: simpletable.AlignCenter, Text: "End"},
		},
	}

	totalRuns, totalTime := 0, 0

	for i, task := range *t {
		i++

		ta := []*simpletable.Cell{
			{Text: fmt.Sprint(i)},
			{Text: green(task.GetShortString(58))},
			{Align: simpletable.AlignCenter, Text: blue(fmt.Sprint(task.Duration))},
			{Align: simpletable.AlignCenter, Text: task.StartedAt.Format("02-01")},
			{Text: task.StartedAt.Format("15:04")},
			{Text: task.FinishedAt.Format("15:04")},
		}

		if interval == "d" && time.Now().Day() == task.StartedAt.Day() {
			interval = time.Now().Weekday().String()
			totalRuns++
			totalTime += task.Duration
			table.Body.Cells = append(table.Body.Cells, ta)
		} else if interval != "d" {
			if interval == "m" {
				interval = time.Now().Month().String()
			}
			totalRuns++
			totalTime += task.Duration
			table.Body.Cells = append(table.Body.Cells, ta)
		}
	}

	totalTimePretty := helper.PrintTimePretty(totalTime, "m")

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Span: 2, Text: red(fmt.Sprintf("%d", totalRuns))},
		{Align: simpletable.AlignCenter, Text: red(totalTimePretty[:len(totalTimePretty)-2])},
		{Align: simpletable.AlignCenter, Text: red(interval)},
		{},
		{},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

// get the string shortend to max value -> limit length when it gets displayed in a table
func (t Task) GetShortString(max int) string {
	if len(t.Name) <= max {
		return t.Name
	} else {
		return t.Name[:max-3] + "..."
	}

}
