package task

import "testing"

func getTaskManagerMock() *TaskManager {
	tm := NewTaskManager()
	tm.tasks = []*Task{
		{ID: 1, Label: "First task", Completed: false},
		{ID: 2, Label: "Second task", Completed: true},
		{ID: 3, Label: "Third task", Completed: false},
	}
	return tm
}

func TestGetAll(t *testing.T) {
	tm := getTaskManagerMock()
	tsks := tm.All()

	if len(tsks) != len(tm.tasks) {
		t.Errorf("expected %d tasks, got %d", len(tm.tasks), len(tsks))
	}
}

func TestCreate(t *testing.T) {
	tm := getTaskManagerMock()

	lbl := "Foo"
	_, err := tm.Save(&Task{Label: lbl, Completed: false})

	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	tm := getTaskManagerMock()

	tsk, err := tm.Find(2)
	if err != nil {
		t.Error(err)
	}

	if tsk.ID != 2 {
		t.Errorf("expected id %d, got %d", 2, tsk.ID)
	}
	if tsk.Label != "Second task" {
		t.Errorf("expected label \"%s\", got \"%s\"", "Second task", tsk.Label)
	}
	if tsk.Completed != true {
		t.Errorf("expected task completed to be %v, got %v", true, tsk.Completed)
	}
}

func TestDeleteOne(t *testing.T) {
	tm := getTaskManagerMock()

	tm.Delete(2)

	if len(tm.tasks) != 2 {
		t.Errorf("expected %d task, got %d", 2, len(tm.tasks))
	}
}

func TestDeleteAll(t *testing.T) {
	tm := getTaskManagerMock()

	tm.DeleteAll()

	if len(tm.tasks) != 0 {
		t.Errorf("expected to delete all tasks, got %d task left", len(tm.tasks))
	}
}

func TestPatchLabel(t *testing.T) {
	tm := getTaskManagerMock()

	var itemID int64 = 2
	newLbl := "Finish todo app"

	tm.Patch(itemID, "label", newLbl)

	item, err := tm.Find(itemID)
	if err != nil {
		t.Error(err)
	}

	if item.Label != newLbl {
		t.Errorf("expected label \"%s\", got \"%s\"", "Second task", item.Label)
	}
}

func TestPatchCompleted(t *testing.T) {
	tm := getTaskManagerMock()

	var itemID int64 = 3
	newStatus := true

	tm.Patch(itemID, "completed", newStatus)

	item, err := tm.Find(itemID)
	if err != nil {
		t.Error(err)
	}

	if item.Completed != newStatus {
		t.Error("expected completed task")
	}
}
