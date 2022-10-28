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
	return json.Unmarshal(buf, t)
}

// PrintStats pretty prints a table with Tasks done - possible to show just daily or whole month - default is daily
func (t *TaskList) PrintStats(dayOrMonth string) {
	var isMonth bool
	if dayOrMonth == "m" {
		isMonth = true
	}

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

	dailyRuns, dailyTime, totalRuns, totalTime := 0, 0, 0, 0

	for i, task := range *t {
		i++

		ta := []*simpletable.Cell{
			{Text: fmt.Sprint(i)},
			{Text: green(task.Name)},
			{Align: simpletable.AlignCenter, Text: blue(fmt.Sprint(task.Duration))},
			{Align: simpletable.AlignCenter, Text: task.StartedAt.Format("02-01")},
			{Text: task.StartedAt.Format("15:04")},
			{Text: task.FinishedAt.Format("15:04")},
		}

		// add only task done at the same day to the cells // todo: change to switch and let user choose specific day
		if !isMonth && time.Now().Day() == task.StartedAt.Day() {
			dailyRuns++
			dailyTime += task.Duration
			table.Body.Cells = append(table.Body.Cells, ta)
		} else if isMonth {
			totalRuns++
			totalTime += task.Duration
			table.Body.Cells = append(table.Body.Cells, ta)
		}
	}

	dailyTimePretty := helper.PrintTimePretty(dailyTime, "m")
	totalTimePretty := helper.PrintTimePretty(totalTime, "m")

	if !isMonth {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Span: 2, Text: red(fmt.Sprintf("%d", dailyRuns))},
			{Align: simpletable.AlignCenter, Text: red(dailyTimePretty[:len(dailyTimePretty)-2])}, // slice 0s out of daily duration
			{Align: simpletable.AlignCenter, Text: red(time.Now().Weekday().String())},
			{},
			{},
		}}
	} else {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Span: 2, Text: red(fmt.Sprintf("%d", totalRuns))},
			{Align: simpletable.AlignCenter, Text: red(totalTimePretty[:len(totalTimePretty)-2])},
			{Align: simpletable.AlignCenter, Text: red(time.Now().Month().String())},
			{},
			{},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
