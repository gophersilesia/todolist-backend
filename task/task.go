package task

import "errors"

var (
	errItemNotFound = errors.New("The given item was not found")
)

// Task type
type Task struct {
	ID        int64  `json:"id"`
	Label     string `json:"label"`
	Completed bool   `json:"completed"`
}

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	tasks  []*Task
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *Task) (*Task, error) {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return task, nil
	}

	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			return m.tasks[i], nil
		}
	}

	return nil, errItemNotFound
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []*Task {
	return m.tasks
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(ID int64) (*Task, error) {
	for _, t := range m.tasks {
		if t.ID == ID {
			return t, nil
		}
	}
	return nil, errItemNotFound
}

func (m *TaskManager) Delete(ID int64) {
	for i, tsk := range m.tasks {
		if tsk.ID == ID {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			continue
		}
	}
}

func (m *TaskManager) DeleteAll() {
	m.tasks = []*Task{}
}

func (m *TaskManager) Patch(ID int64, prop string, val interface{}) error {
	tsk, err := m.Find(ID)
	if err != nil {
		return err
	}

	switch prop {
	case "label":
		tsk.Label = val.(string)
		break
	case "completed":
		tsk.Completed = val.(bool)
	}

	return nil
}
