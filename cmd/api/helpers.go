package main

import (
	"encoding/json"
	"net/http"
)

//create an envelope type
type envelope map[string]any

func (a *applicationDependencies)writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	jsResponse, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsResponse = append(jsResponse, '\n')
	//additional headers to be set
	for key, value := range headers {
		w.Header()[key] = value
		//w.Header().Set(key, value)
	}
	//set content type header
	w.Header().Set("Content-Type", "application/json")
	//explicitly set the response status code
	w.WriteHeader(status)
	_, err = w.Write(jsResponse)
	if err != nil {
		return err
	}

	return nil
}