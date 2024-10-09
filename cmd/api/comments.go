package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	//"github.com/tchenbz/comments/internal/data"
)

func (a *applicationDependencies)createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Content string `json:"content"`
		Author string `json:"author"`
	}

	//err := json.NewDecoder(r.Body).Decode(&incomingData)
	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		//a.errorResponseJSON(w, r, http.StatusBadRequest, err.Error())
		a.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData)
}