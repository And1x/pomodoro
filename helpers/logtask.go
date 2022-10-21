package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Task gets added after time is elapsed

// create Tasks and marshal them in JSON- no need to read from file
// every JSON entry is a Run -> to calc Total runs just loop trough JSON elements
// same could be done with TOTAL-Time

type Task struct {
	Name       string
	Duration   int
	Date       string // Needed??
	StartedAt  time.Time
	FinishedAt time.Time
}

type TaskList []Task

// todo: use allot of single Params??
// func (t *Tasks) Add(name string, duration int, date, start, end time.Time) {
func (t *TaskList) Add(nTask *Task) {

	newTask := Task{
		Name:       nTask.Name,
		Duration:   nTask.Duration,
		Date:       nTask.Date,
		StartedAt:  nTask.StartedAt,
		FinishedAt: nTask.FinishedAt,
	}

	*t = append(*t, newTask)
}

func (t *TaskList) Save(filename string) error {
	d, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, d, 0644)
}

func (t *TaskList) Load(filename string) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { // check this err when func gets called like in main??
			return nil
		}
		return nil
	}

	if len(f) == 0 {
		return err
	}

	err = json.Unmarshal(f, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskList) TotalRuns() int {
	total := 0
	for i := range *t { // todo: rewrite to use only 1var to do this
		total = i
	}
	return total
}
