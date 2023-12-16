package application

import (
	"time"

	"github.com/dafailyasa/golang-template/pkg/logger/models"
	"github.com/dafailyasa/golang-template/pkg/logger/ports"
	"go.uber.org/zap"
)

type Logger struct {
	repo ports.LoggerRepository
	zap  *zap.SugaredLogger
}

var _ ports.LoggerApplication = (*Logger)(nil)

func NewLogger(repo ports.LoggerRepository) *Logger {
	logger, _ := zap.NewProduction()

	sugar := logger.Sugar()
	return &Logger{
		zap:  sugar,
		repo: repo,
	}
}

func (l Logger) Close() error {
	return l.zap.Sync()
}

func (l Logger) Debug(msg string, args any) {
	l.zap.Debug(msg, args)

	log := models.Log{
		Level:     "Debug",
		Message:   msg,
		CreatedAt: time.Now(),
	}

	err := l.repo.Save(&log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}

func (l Logger) Warn(msg string, args any) {
	l.zap.Warn(msg, args)

	log := models.Log{
		Level:     "Warning",
		Message:   msg,
		CreatedAt: time.Now(),
	}

	err := l.repo.Save(&log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}

func (l Logger) Info(msg string, args any) {
	l.zap.Info(msg, args)

	log := models.Log{
		Level:     "Info",
		Message:   msg,
		CreatedAt: time.Now(),
	}

	err := l.repo.Save(&log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}

func (l Logger) Error(msg string, args any) {
	l.zap.Error(msg, args)

	log := models.Log{
		Level:     "Error",
		Message:   msg,
		CreatedAt: time.Now(),
	}

	err := l.repo.Save(&log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}
