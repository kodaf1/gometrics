package urlparser

import (
	"github.com/kodaf1/gometrics/internal/errors"
	"strings"
)

type UrlMetric struct {
	Type  string
	Name  string
	Value string
}

func Parse(url string) (m *UrlMetric, err error) {
	urlSlice := strings.Split(strings.Trim(url, "/"), "/")

	if len(urlSlice) != 4 {
		return nil, errors.IncorrectParamsCount
	}

	return &UrlMetric{
		Type:  urlSlice[1],
		Name:  urlSlice[2],
		Value: urlSlice[3],
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
