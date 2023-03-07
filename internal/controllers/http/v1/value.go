package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-values-generator/internal/controllers/http/dto"
	"go-values-generator/internal/domain/entities"
	"go-values-generator/internal/domain/service"
	"io"
	"net/http"
)

type ValueHandler struct {
	service *service.ValueService
}

var (
	baseURL          = "/api/v1"
	generateValueURL = "/generate"
	retrieveValueURL = "/retrieve"
)

func NewValueHandler(service *service.ValueService) *ValueHandler {
	return &ValueHandler{service: service}
}

func (v ValueHandler) Register(router chi.Router) {
	router.Route(baseURL, func(r chi.Router) {
		r.Post(generateValueURL, v.createValue)
		r.Post(retrieveValueURL, v.retrieveValue)
	})
}

func (v ValueHandler) createValue(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var createValueDTO dto.CreateValueDTO
	if err := json.Unmarshal(body, &createValueDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createValueDTO.RequestID = r.Context().Value(middleware.RequestIDKey).(string)

	id, value, err := v.service.CreateValue(entities.CreateValueDTO(createValueDTO))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := dto.ShowValueDTO{
		ID:    id,
		Value: value.Val,
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response.Marshal())
}

func (v ValueHandler) retrieveValue(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var getValueDTO dto.GetValueDTO
	if err := json.Unmarshal(body, &getValueDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := v.service.GetValueByID(getValueDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := dto.ShowValueDTO{
		Value: value.Val,
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response.Marshal())
}
