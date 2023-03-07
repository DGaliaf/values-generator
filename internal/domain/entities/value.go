package entities

import (
	"encoding/json"
	"go-values-generator/internal/config"
)

type Value struct {
	Val       string `json:"val,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

func (v Value) Marshal() ([]byte, error) {
	return json.Marshal(v)
}

type CreateValueDTO struct {
	Type      string `json:"type,omitempty"`
	Len       int    `json:"len,omitempty"`
	RequestID string
}

func (v CreateValueDTO) Validate() CreateValueDTO {
	if v.Type == "" {
		v.Type = config.DefaultType
	}

	if v.Len == 0 {
		v.Len = config.DefaultLen
	}

	return v
}
