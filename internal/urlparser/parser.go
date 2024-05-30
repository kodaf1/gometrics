package urlparser

import (
	"github.com/kodaf1/gometrics/internal/errors"
	"log"
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

	log.Println(len(urlSlice))

	if len(urlSlice) != 4 {
		return nil, errors.IncorrectParamsCount
	}

	if urlSlice[0] != "update" {
		return nil, errors.UnknownMethod
	}

	if inStringsArray(urlSlice[1], validMetricTypes) {
		return nil, errors.UnknownType
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