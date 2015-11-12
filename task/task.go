package task

import "github.com/gogits/gogs/modules/uuid"

// Task type
type Task struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Completed bool   `json:"completed"`
}

// Tasks type
type Tasks map[string]*Task

// All returns the list of all the Tasks
func (t Tasks) All() []*Task {
	items := []*Task{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

// Find finds the task in the map
func (t Tasks) Find(ID string) *Task {
	if item, found := t[ID]; found {
		return item
	}
	return nil
}

// Create creates a new task
func (t Tasks) Create(item Task) *Task {
	item.ID = newID()
	t[item.ID] = &item
	return &item
}

// Update updates a task
func (t Tasks) Update(ID string, updatedItem Task) *Task {
	if item := t.Find(ID); item != nil {
		return item.update(updatedItem)
	}
	return nil
}

// DeleteAll deletes all tasks
func (t Tasks) DeleteAll() string {
	for k := range t {
		delete(t, k)
	}
	return ""
}

// Delete deletes a single task
func (t Tasks) Delete(ID string) string {
	for k := range t {
		if k == ID {
			delete(t, k)
		}
	}
	return ""
}

// newID generates a unique ID for the task
func newID() string {
	id := uuid.NewV4()
	return id.String()
}

// Internal task update
func (i *Task) update(item Task) *Task {
	i.Label = item.Label
	i.Completed = item.Completed
	return i
}
