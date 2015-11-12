package task

import (
	"testing"
)

var tasks Tasks

func init() {
	tasks = Tasks{}
}

func TestCreate(t *testing.T) {
	lbl := "Foo"
	task := tasks.Create(Task{Label: lbl, Completed: false})

	if task.Label != lbl {
		t.Errorf("expected text %q, got %q", lbl, task.Label)
	}
	if task.Completed {
		t.Errorf("a new instance should not return as completed")
	}
}

func TestDeleteOne(t *testing.T) {

}

func TestDeleteAll(t *testing.T) {

}

func TestPatch(t *testing.T) {

}
