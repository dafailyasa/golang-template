package ports

import "github.com/dafailyasa/golang-template/pkg/logger/models"

type LoggerRepository interface {
	Save(log *models.Log) error
}
