package urlparser

import (
	"github.com/kodaf1/gometrics/internal/errors"
	"strings"
)

var validMetricTypes = []string{"gauge", "counter"}

type UrlMetric struct {
	Type  string
	Name  string
	Value string
}

func Parse(url string) (m *UrlMetric, err error) {
	urlSlice := strings.Split(strings.Trim(url, "/"), "/")

	if len(urlSlice) != 3 {
		return nil, errors.IncorrectParamsCount
	}

	if !inStringsArray(urlSlice[1], validMetricTypes) {
		return nil, errors.UnknownType
	}

	return &UrlMetric{
		Type:  urlSlice[0],
		Name:  urlSlice[1],
		Value: urlSlice[2],
	}, nil
}

func inStringsArray(val string, array []string) bool {
	for i := range array {
		if array[i] == val {
			return true
		}
	}
	return false
}
