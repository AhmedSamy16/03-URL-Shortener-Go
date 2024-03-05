package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AhmedSamy16/03-url-shortener-Go/repository"
	"github.com/AhmedSamy16/03-url-shortener-Go/types"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UrlsHandler struct {
	UrlsRepository *repository.UrlRepository
}

func (handler *UrlsHandler) GetAllUrls(w http.ResponseWriter, r *http.Request) {
	data, err := handler.UrlsRepository.GetAllUrls(r.Context())
	if err != nil {
		respondWithError(w, 500, "Failed to get users")
		return
	}
	respondWithJson(w, 200, data)
}

func (handler *UrlsHandler) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrlParams := chi.URLParam(r, "shortUrl")
	shortUrl, err := uuid.Parse(shortUrlParams)
	if err != nil {
		respondWithError(w, 400, "Invalid url id")
		return
	}
	data, err := handler.UrlsRepository.GetShortUrl(r.Context(), shortUrl)
	if err != nil {
		log.Println(err)
		respondWithError(w, 404, "url not found")
		return
	}
	respondWithJson(w, 200, data)
}

func (handler *UrlsHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	urlToCreate := types.CreateShortUrl{}
	if err := json.NewDecoder(r.Body).Decode(&urlToCreate); err != nil {
		respondWithError(w, 400, "Invalid data")
		return
	}
	defer r.Body.Close()

	data, err := handler.UrlsRepository.CreateShortUrl(r.Context(), &urlToCreate)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	respondWithJson(w, 200, data)
}
