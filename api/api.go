package api

import "net/http"

func NewHandler() (http.Handler, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", createClock)
	return mux, nil
}

func createClock(resp http.ResponseWriter, req *http.Request) {
	// Create the circle
	// Need to parse the parameters from the request body
	// Get the height and width of the svg
	// Get the hour and minute for the hands
	// We need to get the hour and minute
}
