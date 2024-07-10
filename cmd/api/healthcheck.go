package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Maps to hold information
	data := map[string]string{
		"status":      "available",
		"enivrnoment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and cloud not process your request", http.StatusInternalServerError)
		return
	}

	// For terminal view add a new line
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}
