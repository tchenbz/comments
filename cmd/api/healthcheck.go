package main

import (
  "fmt"
  "net/http"
)

func (a *applicationDependencies)healthCheckHandler(w http.ResponseWriter,r *http.Request) {
    //fmt.Fprintln(w, "status: available")
    //fmt.Fprintf(w, "environment: %s\n", a.config.environment)
    //fmt.Fprintf(w, "version: %s\n", appVersion)

    jsResponse := `{"status": "available", "environment": %q, "version": %q}`
    jsResponse = fmt.Sprintf(jsResponse, a.config.environment, appVersion)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(jsResponse))
}