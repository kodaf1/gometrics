package storage

import (
	"github.com/kodaf1/gometrics/internal/errors"
	"github.com/kodaf1/gometrics/internal/urlparser"
	"strconv"
)

type MapStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

var storage MapStorage = MapStorage{
	Gauge:   make(map[string]float64),
	Counter: make(map[string]int64),
}

func SaveData(data *urlparser.UrlMetric) error {
	switch data.Type {
	case "gauge":
		result, err := strconv.ParseFloat(data.Value, 64)
		if err != nil {
			return err
		}

		storage.Gauge[data.Name] = result

		return nil
	case "counter":
		result, err := strconv.ParseInt(data.Value, 10, 64)
		if err != nil {
			return err
		}

		storage.Counter[data.Name] += result

		return nil
	default:
		return errors.UnknownType
	}
}
