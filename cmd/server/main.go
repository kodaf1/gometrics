package main

import (
	"errors"
	"github.com/kodaf1/gometrics/internal/urlparser"
	"net/http"
)

// TODO: escape renaming
import (
	myErrors "github.com/kodaf1/gometrics/internal/errors"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	http.Handle("/update/", http.HandlerFunc(webhook))
	return http.ListenAndServe(`localhost:8080`, nil)
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := urlparser.Parse(r.URL.Path)

	switch {
	case errors.Is(err, myErrors.IncorrectParamsCount):
		w.WriteHeader(http.StatusNotFound)
		return
	case errors.Is(err, myErrors.UnknownType):
		w.WriteHeader(http.StatusBadRequest)
		return
	case errors.Is(err, nil):
		break
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
