package main

import (
	"github.com/kodaf1/gometrics/internal/errors"
	"github.com/kodaf1/gometrics/internal/urlparser"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return http.ListenAndServe(`localhost:8080`, http.HandlerFunc(webhook))
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := urlparser.Parse(r.URL.Path)

	switch err {
	case errors.IncorrectParamsCount:
		w.WriteHeader(http.StatusNotFound)
		return
	case errors.UnknownType:
	case errors.UnknownMethod:
		w.WriteHeader(http.StatusBadRequest)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
