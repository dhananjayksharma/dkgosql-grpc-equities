package handlers

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
