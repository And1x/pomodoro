package logtask

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func createMockTaskList() *TaskList {
	now := time.Now()
	var duration time.Duration = 25 * time.Minute
	mockTaskList := TaskList{
		Task{Name: "testTask", Duration: int(duration.Minutes()), StartedAt: now, FinishedAt: now.Add(duration)},
	}
	return &mockTaskList
}

func TestAdd(t *testing.T) {
	mockTaskList := createMockTaskList()
	newTask := Task{
		Name:       "add a new Task",
		Duration:   25,
		StartedAt:  time.Time{},
		FinishedAt: time.Time{},
	}
	mockTaskList.Add(&newTask)

	// log.Println((*mockTaskList)[1])

	if len(*mockTaskList) != 2 {
		t.Errorf("got %d, wanted %d", len(*mockTaskList), 2)
	}
}

// write to a buffer instead of a file
func TestSave(t *testing.T) {
	mockTaskList := createMockTaskList()
	buf := new(bytes.Buffer)

	err := mockTaskList.Save(buf)
	if err != nil {
		t.Error(err)
	}

	if buf.Len() == 0 {
		t.Errorf("got %d, but buffer should be > 0 after writing to it", buf.Len())
	}
}

// load from a buffer istead a file
func TestLoad(t *testing.T) {

	emptyTaskList := &TaskList{}
	mockTaskList := createMockTaskList()
	buf := new(bytes.Buffer)
	err := mockTaskList.Save(buf)
	if err != nil {
		t.Errorf("failed to save into buffer: %v", err)
	}

	err = emptyTaskList.Load(buf)
	if err != nil {
		t.Errorf("failed to load from buffer: %v", err)
	}

	if !cmp.Equal(emptyTaskList, mockTaskList) {
		t.Errorf("got %v, wanted %v", emptyTaskList, mockTaskList)
	}

	// for i := range emptyTaskList {
	// 	if emptyTaskList[i].Name != (*mockTaskList)[i].Name {
	// 		t.Errorf("got %v, want %v", emptyTaskList[i].Name, (*mockTaskList)[i].Name)
	// 	}
	// }
}
