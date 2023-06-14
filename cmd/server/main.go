package main

import (
	"encoding/json"
	"github.com/alirezaazadi/asinodb.git"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	d := asinodb.New()

	r := mux.NewRouter()
	r.Handle("/{key}", getHandler(d)).Methods(http.MethodGet)
	r.Handle("/{key}", setHandler(d)).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

func getHandler(d *asinodb.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := mux.Vars(r)
		value, err := d.Get(pathVars["key"])

		if err != nil {
			if err == asinodb.ErrNothing {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}

			writeJSON(w, &struct {
				Message string `json:"message"`
			}{Message: err.Error()})
			return
		}

		writeJSON(w, &struct {
			Value interface{} `json:"value"`
		}{Value: value})
	}
}

func setHandler(d *asinodb.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
