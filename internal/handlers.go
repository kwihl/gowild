package internal

import "net/http"

type ForestHandler interface {
	Animals(w http.ResponseWriter, req *http.Request)
	Plants(w http.ResponseWriter, req *http.Request)
}

type MountainHandler interface {
	Animals(w http.ResponseWriter, req *http.Request)
	Plants(w http.ResponseWriter, req *http.Request)
	Peaks(w http.ResponseWriter, req *http.Request)
}
