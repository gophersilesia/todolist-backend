package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gopherskatowice/todolist-backend/task"
	"github.com/julienschmidt/httprouter"
)

var tasks task.Tasks

func init() {
	tasks = task.Tasks{}
}

// RegisterHandlers registers httprouter handlers
func RegisterHandlers() *httprouter.Router {
	rt := httprouter.New()

	rt.GET("/tasks", ListTasks)
	rt.PATCH("/tasks/:id", PatchTask)
	rt.POST("/tasks", CreateTask)
	rt.DELETE("/tasks", DeleteTasks)
	rt.DELETE("/tasks/:id", DeleteTask)

	return rt
}

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

// errorHandler wraps a function returning an error by handling the error and returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as described for every type.
// If the error is of another type, it is considered as an internal error and its message is logged.
func errorHandler(f func(w http.ResponseWriter, r *http.Request, p httprouter.Params) error) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := f(w, r, p)
		if err == nil {
			return
		}

		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

// ListTasks handles GET requests on /tasks
func ListTasks(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	res := tasks.All()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(res)
}

// CreateTask handles POST request
func CreateTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tsk := task.Task{}
	var err error

	// Unmarshal JSON to Task Object
	if err = json.NewDecoder(req.Body).Decode(&tsk); err != nil {
		w.WriteHeader(400)
		return
	}

	// Create task
	t := tasks.Create(tsk)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(t)
}

func PatchTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tid := p.ByName("id")
	tsk := tasks.Find(tid)
	if tsk == nil {

	}
}

func DeleteTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tid := p.ByName("id")
	tsk := tasks.Find(tid)
	if tsk == nil {

	}
}

func DeleteTasks(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

}
