package v1

import (
	"animal-play/internal/adapters/primary/web/api/v1/public"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	publicRouter *public.Router
}

func NewRouter(db *sqlx.DB) *Router {
	return &Router{
		publicRouter: public.NewRouter(db),
	}
}

func (r *Router) Routes() http.Handler {
	chiRouter := chi.NewRouter()
	chiRouter.Mount("/public", r.publicRouter.Routes())

	return chiRouter
}
