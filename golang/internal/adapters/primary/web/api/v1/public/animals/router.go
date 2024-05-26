package animals

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	handler *Handler
}

func NewRouter(db *sqlx.DB) *Router {
	handler := NewHandler(db)

	return &Router{
		handler: handler,
	}
}

func (r *Router) Routes() http.Handler {
	chiRouter := chi.NewRouter()
	chiRouter.Get("/", r.handler.GetAll)

	return chiRouter
}
