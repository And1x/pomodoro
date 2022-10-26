package helpers

import (
	"log"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {

	now := time.Now()
	var duration time.Duration = 25 * time.Minute
	mockTaskList := TaskList{
		Task{Name: "testTask", Duration: int(duration.Minutes()), StartedAt: now, FinishedAt: now.Add(duration)},
	}

	newTask := Task{
		Name:       "add a new Task",
		Duration:   int(duration),
		StartedAt:  now,
		FinishedAt: now.Add(duration),
	}

	log.Println(mockTaskList)

	mockTaskList.Add(&newTask)

	if len(mockTaskList) != 2 {
		t.Errorf("got %d, wanted %d", len(mockTaskList), 2)
	}
}

func TestSave(t *testing.T) {

}
