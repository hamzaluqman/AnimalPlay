package api

import (
	v1 "animal-play/internal/adapters/primary/web/api/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	db       *sqlx.DB
	v1Router *v1.Router
}

func NewRouter(db *sqlx.DB) *Router {
	v1Router := v1.NewRouter(db)
	return &Router{
		db:       db,
		v1Router: v1Router,
	}
}

func (r *Router) MountRoutes() http.Handler {
	chiRouter := chi.NewRouter()
	chiRouter.Mount("/api/v1", r.v1Router.Routes())

	return chiRouter
}
