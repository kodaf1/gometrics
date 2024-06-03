package storage

import (
	"github.com/kodaf1/gometrics/internal/dto"
	"github.com/kodaf1/gometrics/internal/errors"
	"strconv"
)

type MapStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
}

func (s *MapStorage) SaveData(data *dto.UpdateMetricDTO) error {
	switch data.Type {
	case "gauge":
		result, err := strconv.ParseFloat(data.Value, 64)
		if err != nil {
			return err
		}

		s.Gauge[data.Name] = result

		return nil
	case "counter":
		result, err := strconv.ParseInt(data.Value, 10, 64)
		if err != nil {
			return err
		}

		s.Counter[data.Name] += result

		return nil
	default:
		return errors.UnknownType
	}
}
