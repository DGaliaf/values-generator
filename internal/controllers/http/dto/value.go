package dto

import "encoding/json"

type CreateValueDTO struct {
	Type      string `json:"type,omitempty"`
	Len       int    `json:"len,omitempty"`
	RequestID string
}

type ShowValueDTO struct {
	ID    string `json:"value_id"`
	Value string `json:"value"`
}

type GetValueDTO struct {
	ID string `json:"value_id"`
}

func (s ShowValueDTO) Marshal() []byte {
	m, _ := json.Marshal(s)
	return m
}
