package TogglRichPresence

import (
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TogglWrapper struct {
	client http.Client
	apiKey string
}

type Timer struct {
	startTime   time.Time
	description string
	project     string
	tags        []string
}

// Initialize basically constructs the TogglWrapper.
//
// Adds a client struct to the TogglWrapper, and sets
// the API key for Toggl.
func InitializeTogglWrapper(apiKey string) TogglWrapper {
	var w TogglWrapper
	w.client = http.Client{
		Timeout: 10 * time.Second,
	}
	w.apiKey = apiKey
	return w
}

func (w TogglWrapper) getProjectName(projectId int64) string {
	req, _ := http.NewRequest("GET", "https://www.toggl.com/api/v8/projects/"+strconv.Itoa(int(projectId)), nil)
	req.SetBasicAuth(w.apiKey, "api_token")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request for project: %s", err)
	}
	defer resp.Body.Close()

	requestString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading request body of project: %s", err)
	}

	name, err := jsonparser.GetString(requestString, "data", "name")
	if err != nil {
		log.Fatalf("Error extracting name of project: %s", err)
	}
	return name
}

func (w TogglWrapper) CurrentTimer() Timer {
	req, _ := http.NewRequest("GET", "https://www.toggl.com/api/v8/time_entries/current", nil)
	req.SetBasicAuth(w.apiKey, "api_token")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request for current timer: %s", err)
	}
	defer resp.Body.Close()

	requestString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading request body of current timer: %s", err)
	}

	description, err := jsonparser.GetString(requestString, "data", "description")
	if err != nil {
        return Timer{
		startTime:   time.Now(),
		description: "",
		project:     "",
		tags:        []string{""},
        }
	}

	projectId, err := jsonparser.GetInt(requestString, "data", "pid")
	if err != nil {
		log.Fatalf("Error extracting project ID of timer entry: %s", err)
	}

	startTimeString, err := jsonparser.GetString(requestString, "data", "start")
	if err != nil {
		log.Fatalf("Error extracting start time of timer entry: %s", err)
	}
	startTime, err := time.Parse(time.RFC3339, startTimeString)
	if err != nil {
		log.Fatalf("Error parsing time string from time entry: %s", err)
	}

	var tags []string
	_, err = jsonparser.ArrayEach(requestString, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tags = append(tags, string(value))
	}, "data", "tags")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		log.Fatalf("Error iterating through tags array in JSON of current timer: %s", err)
	} else if err == jsonparser.KeyPathNotFoundError {
		tags = nil;
	}

	runningTimer := Timer{
		startTime:   startTime,
		description: description,
		project:     w.getProjectName(projectId),
		tags:        tags,
	}

	return runningTimer
}

func (w TogglWrapper) currentTimerID() int64 {
	req, _ := http.NewRequest("GET", "https://www.toggl.com/api/v8/time_entries/current", nil)
	req.SetBasicAuth(w.apiKey, "api_token")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request for current timer ID: %s", err)
	}
	defer resp.Body.Close()

	requestString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading request body of current timer: %s", err)
	}

	currentTimerID, err := jsonparser.GetInt(requestString, "data", "id")
	if err != nil {
		log.Fatalf("Error extracting ID of current timer: %s", err)
	}

	return currentTimerID
}

func (w TogglWrapper) StopTimer() {
	projectId := w.currentTimerID()
	req, _ := http.NewRequest("PUT", "https://www.toggl.com/api/v8/time_entries/"+strconv.Itoa(int(projectId))+"/stop", nil)
	req.SetBasicAuth(w.apiKey, "api_token")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to stop timer: %s", err)
	}
	defer resp.Body.Close()
}
