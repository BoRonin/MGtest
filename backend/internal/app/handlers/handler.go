package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"mgtest/internal/models"
	"mgtest/internal/service"
	"mgtest/pkg/utils"
	"net/http"
)

type Handler struct {
	s   *service.Service
	Log *slog.Logger
}

func New(s *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		s:   s,
		Log: log,
	}
}

func (h *Handler) AddPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data := models.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	result, err := h.s.InsertProfile(ctx, data)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, result)
}

func (h *Handler) EditPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data := models.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	result, err := h.s.UpdateProfile(ctx, data)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, result)
}

func (h *Handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := r.URL.Query().Get("id")
	data, err := h.s.PS.GetProfile(ctx, id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, data)
}
