package animals

import (
	"animal-play/internal/adapters/primary/web/api/handler_response"
	mysql_repository "animal-play/internal/adapters/secondary/mysql/repositories"
	"animal-play/internal/core/animal"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	animalService *animal.Service
}

func NewHandler(db *sqlx.DB) *Handler {
	animalRepo := mysql_repository.NewAnimalRepository(db)
	animalService := animal.NewService(animalRepo)

	return &Handler{
		animalService: animalService,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	animals, err := h.animalService.GetAll()

	if err != nil {
		// Add logging
		errMsg := fmt.Sprintf("Failed to get animals. Error: %v", err)
		handler_response.WriteErrorResponse(w, http.StatusInternalServerError, errMsg)
		return
	}

	handler_response.WriteSuccessResponse(w, animals, http.StatusOK)
}
