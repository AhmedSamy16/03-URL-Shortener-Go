package application

import (
	"github.com/AhmedSamy16/03-url-shortener-Go/handlers"
	"github.com/AhmedSamy16/03-url-shortener-Go/repository"
	"github.com/go-chi/chi/v5"
)

func (app *App) LoadRoutes() {
	router := chi.NewRouter()

	router.Route("/urls", app.LoadUrlsRoutes)

	app.router = router
}

func (app *App) LoadUrlsRoutes(router chi.Router) {
	repo := &repository.UrlRepository{
		DB: app.DB,
	}
	urlsHandler := &handlers.UrlsHandler{
		UrlsRepository: repo,
	}

	router.Get("/", urlsHandler.GetAllUrls)
	router.Post("/", urlsHandler.CreateShortUrl)
	router.Get("/{shortUrl}", urlsHandler.GetShortUrl)
}
