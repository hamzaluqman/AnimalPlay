package web

import (
	"animal-play/internal/adapters/primary/web/api"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Router interface {
	MountRoutes() http.Handler
}

type WebHandler struct {
	router Router
}

func NewWebHandler(db *sqlx.DB) *WebHandler {
	router := api.NewRouter(db)
	return &WebHandler{
		router: router,
	}
}

func (wH *WebHandler) MountRoutes() http.Handler {
	router := wH.router.MountRoutes()
	return router
}
