package public

import (
	"animal-play/internal/adapters/primary/web/api/v1/public/animals"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	animalRouter *animals.Router
}

func NewRouter(db *sqlx.DB) *Router {
	return &Router{
		animalRouter: animals.NewRouter(db),
	}
}

func (r *Router) Routes() http.Handler {
	chiRouter := chi.NewRouter()
	chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from public get"))
	})

	chiRouter.Mount("/animals", r.animalRouter.Routes())

	return chiRouter
}
