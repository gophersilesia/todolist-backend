package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gopherskatowice/todolist-backend/task"
	"github.com/julienschmidt/httprouter"
)

var tm *task.TaskManager

func init() {
	tm = task.NewTaskManager()
}

// RegisterHandlers registers httprouter handlers
func RegisterHandlers() *httprouter.Router {
	rt := httprouter.New()

	rt.GET("/tasks", errorHandler(ListTasks))
	rt.PATCH("/tasks/:id", errorHandler(PatchTask))
	rt.PUT("/tasks/:id", errorHandler(PatchTask))
	rt.POST("/tasks", errorHandler(CreateTask))
	rt.DELETE("/tasks", errorHandler(DeleteTasks))
	rt.DELETE("/tasks/:id", errorHandler(DeleteTask))

	return rt
}

// ListTasks handles GET requests on /tasks
func ListTasks(w http.ResponseWriter, req *http.Request, p httprouter.Params) error {
	res := tm.All()

	handleSuccess(w, http.StatusOK, res)
	return nil
}

// CreateTask handles POST request
func CreateTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) error {
	tsk := task.Task{}
	var err error

	// Unmarshal JSON to Task Object
	if err = json.NewDecoder(req.Body).Decode(&tsk); err != nil {
		return err
	}

	// Create task
	t, err := tm.Save(&tsk)
	if err != nil {
		return err
	}

	handleSuccess(w, http.StatusOK, t)
	return nil
}

// PatchTask updates a property for the given task
func PatchTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) error {
	tid, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	tsk, err := tm.Find(tid)
	if err != nil {
		return err
	}

	var msg map[string]interface{}
	err = json.NewDecoder(req.Body).Decode(&msg)

	var ok bool

	_, ok = msg["label"]
	if ok {
		tm.Patch(tsk.ID, "label", msg["label"].(string))
	}

	_, ok = msg["completed"]
	if ok {
		tm.Patch(tsk.ID, "completed", msg["completed"].(bool))
	}

	handleSuccess(w, http.StatusOK, nil)
	return nil
}

// DeleteTask removes the given task from the stack
func DeleteTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) error {
	tid, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		return err
	}

	tsk, err := tm.Find(tid)
	if err != nil {
		return err
	}
	tm.Delete(tsk.ID)
	handleSuccess(w, http.StatusOK, nil)
	return nil
}

// DeleteTasks remove all the tasks
func DeleteTasks(w http.ResponseWriter, req *http.Request, p httprouter.Params) error {
	tm.DeleteAll()
	handleSuccess(w, http.StatusOK, nil)
	return nil
}

// handleSuccess handles the response for each endpoint.
// It follows the JSEND standard for JSON response.
// See https://labs.omniti.com/labs/jsend
// handleOutput handles the response for each endpoint.
// It follows the JSEND standard for JSON response.
// See https://labs.omniti.com/labs/jsend
func handleSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	success := false
	if code == 200 {
		success = true
	}

	// JSend has three possible statuses: success, fail and error
	// In case of error, there is no data sent, only an error message.
	status := "success"
	msgType := "data"
	if !success {
		status = "error"
		msgType = "message"
	}

	res := map[string]interface{}{"status": status}
	if data != nil {
		res[msgType] = data
	}

	json.NewEncoder(w).Encode(res)
}
