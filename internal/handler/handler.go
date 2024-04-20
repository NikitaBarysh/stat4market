package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/NikitaBarysh/stat4market/internal/model"
	"github.com/NikitaBarysh/stat4market/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouters(router chi.Router) {
	router.Post("/api/event", h.postData)
	router.Get("/get/{time}/{type}", h.Get)
}

func (h *Handler) postData(rw http.ResponseWriter, r *http.Request) {
	var input model.Event

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &input); err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		http.Error(rw, "error ahaha", http.StatusBadRequest)
		return
	}

	err = h.service.Set(r.Context(), input)
	if err != nil {
		log.Printf("Error setting event: %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("OK"))
}

func (h *Handler) Get(rw http.ResponseWriter, r *http.Request) {
	timeURl := chi.URLParam(r, "time")
	eventType := chi.URLParam(r, "type")

	timeSplited := strings.Split(timeURl, "--")
	if len(timeSplited) != 2 {
		log.Printf("Error parsing time: %v", timeSplited)
		http.Error(rw, "Invalid time", http.StatusBadRequest)
		return
	}

	timeBefore, err := time.Parse("2006-01-02", timeSplited[0])
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	timeAfter, err := time.Parse("2006-01-02", timeSplited[1])
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.Get(r.Context(), eventType, timeBefore, timeAfter)
	if err != nil {
		log.Printf("Error getting event: %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}
