package main

import (
	//"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *applicationDependencies)routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(a.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthCheckHandler) 
	return a.recoverPanic(router)
}