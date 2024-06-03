package utils

import (
	"github.com/kodaf1/gometrics/internal/dto"
	"github.com/kodaf1/gometrics/internal/errors"
	"strings"
)

func ParseURLToMetric(url string) (m *dto.UpdateMetricDTO, err error) {
	urlSlice := strings.Split(strings.Trim(url, "/"), "/")

	if len(urlSlice) != 4 {
		return nil, errors.IncorrectParamsCount
	}

	return &dto.UpdateMetricDTO{
		Type:  urlSlice[1],
		Name:  urlSlice[2],
		Value: urlSlice[3],
	}, nil
}
