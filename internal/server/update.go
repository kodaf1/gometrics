package server

import (
	"errors"
	myErrors "github.com/kodaf1/gometrics/internal/errors"
	"github.com/kodaf1/gometrics/internal/utils"
	"net/http"
)

func (s *Server) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dto, err := utils.ParseURLToMetric(r.URL.Path)

	switch {
	case errors.Is(err, myErrors.IncorrectParamsCount):
		w.WriteHeader(http.StatusNotFound)
		return
	case errors.Is(err, nil):
		break
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.Storage.SaveData(dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
