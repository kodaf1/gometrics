package repositories

import "github.com/kodaf1/gometrics/internal/dto"

type StorageRepository interface {
	SaveData(*dto.UpdateMetricDTO) error
}
