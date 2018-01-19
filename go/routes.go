package main

import (
)

type route struct {
Name string
	Method string
		Pattern string
	HandlerFunc http.HandlerFunc
}

type transcriptionJobData struct {AudioURL string `json:"audioURL"`
	EmailAddresses []string `json:"emailAddresses"`}

var routes = []route{
	
	route{
	"hello", "GET",
	"/hello/{name}",
	helloHandler,
},
				route{
					"add_job",
					"POST",
					"/add_job",
					initiateTranscriptionJobHandler,
				},

route{"health", "GET", "/health", healthHandler}, route{
	"job_status",
	"GET",
	"/job_status/{id}",
	jobStatusHandler},
}

func helloHandler(w http.ResponseWriter,
	r *http.Request) {
		args := mux.Vars(r)
				fmt.Fprintf(w, "Hello %s!", args["name"])
}

// initiateTranscriptionJobHandle takes a POST request containing a json object,
// decodes it into an audioData struct, and returns appropriate message.
func initiateTranscriptionJobHandler(w http.ResponseWriter,
	r *http.Request) {
	var jsonData transcriptionJobData

	// unmarshal from the response body directly into our struct
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Accepted!")
}

func healthHandler(w http.ResponseWriter,
	r *http.Request) {
	w.Write([]byte("healthy!"))
}

// jobStatusHandler returns the status of a task with given id.
func jobStatusHandler(w http.ResponseWriter,
	r *http.Request) {
	args := mux.Vars(r)
	id := args["id"]

	status := tasks.DefaultTaskExecuter.GetTaskStatus(id)
	w.Write([]byte(status.String()))
}
