package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/jsonutils"
)

func (api *Api) handleSubscribeToAuction(w http.ResponseWriter, r *http.Request) {

	rawProductID := chi.URLParam(r, "product_id")

	productID, err := uuid.Parse(rawProductID)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"message":    "invalid product id - must be a valid uuid",
			"product_id": productID,
		})
		return
	}

}
