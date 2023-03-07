package service

import (
	"encoding/json"
	"errors"
	"go-values-generator/internal/adapters/db/redis"
	"go-values-generator/internal/domain/entities"
	"go-values-generator/pkg/utils/generators"
	"strconv"
)

type ValueService struct {
	storage *redis.ValueStorage
}

func NewValueService(storage *redis.ValueStorage) *ValueService {
	return &ValueService{storage: storage}
}

func (v ValueService) CreateValue(valueData entities.CreateValueDTO) (string, entities.Value, error) {
	valueData = valueData.Validate()

	if !generators.IsSupportedType(valueData.Type) {
		return "", entities.Value{}, errors.New("unsupported type")
	}

	value, err := v.generateValue(valueData)
	if err != nil {
		return "", entities.Value{}, err
	}

	marshaledValue, err := value.Marshal()
	if err != nil {
		return "", entities.Value{}, err
	}

	id, mVal, err := v.storage.Create(marshaledValue)

	var val entities.Value
	if err := json.Unmarshal(mVal, &val); err != nil {
		return "", entities.Value{}, err
	}

	return id, val, err
}

func (v ValueService) GetValueByID(id string) (entities.Value, error) {
	response, err := v.storage.Get(id)
	if err != nil {
		return entities.Value{}, err
	}

	var value entities.Value
	if err := json.Unmarshal(response, &value); err != nil {
		return entities.Value{}, err
	}

	return value, nil
}

func (v ValueService) generateValue(valueData entities.CreateValueDTO) (entities.Value, error) {
	var val string

	switch valueData.Type {
	case "string":
		val = generators.GenerateString(valueData.Len)
		break
	case "int":
		valInt, err := generators.GenerateInt(valueData.Len)
		if err != nil {
			return entities.Value{}, err
		}

		val = strconv.Itoa(valInt)
		break
	case "uuid":
		val = generators.GenerateUUID()
	}

	return entities.Value{
		Val:       val,
		RequestID: valueData.RequestID,
	}, nil
}
