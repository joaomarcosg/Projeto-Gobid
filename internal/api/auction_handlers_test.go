package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandleSubscribeToAuctionWithInvalidUUID(t *testing.T) {

	api := Api{}

	invalidUUID := "0123456789abc"
	req := httptest.NewRequest("GET", "/api/v1/products/ws/subscribe/"+invalidUUID, nil)

	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("product_id", invalidUUID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.handleSubscribeToAuction)
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

}
