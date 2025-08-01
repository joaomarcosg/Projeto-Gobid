package api

import (
	"errors"
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
			"message": "invalid product id - must be a valid uuid",
		})
		return
	}

	_, err = api.ProductService.GetProductById(r.Context(), productID)
	if err != nil {
		if errors.Is(err, services.ErrProductNotFound) {
			jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{
				"message": "no product with given id",
			})
			return
		}

		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "unexpected error, try again later",
		})
		return
	}

}
