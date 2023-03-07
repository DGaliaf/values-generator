package service

import "go-values-generator/internal/domain/entities"

type Service interface {
	CreateValue(value entities.CreateValueDTO) (string, entities.Value, error)
	GetValueByID(id string) ([]byte, error)
}
