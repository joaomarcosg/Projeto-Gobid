package api

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/services"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

type mockProductStoreHandler struct{}

func (m *mockProductStoreHandler) CreateProduct(
	ctx context.Context,
	sellerID uuid.UUID,
	productName,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	productID := uuid.New()
	return productID, nil
}

func (m *mockProductStoreHandler) GetProductById(ctx context.Context, productID uuid.UUID) (store.Product, error) {
	return store.Product{
		ID:          productID,
		SellerID:    uuid.New(),
		ProductName: "Producto de teste",
		Description: "Descrição do produto",
		Baseprice:   99.99,
		AuctionEnd:  time.Now().Add(2 * time.Hour),
	}, nil
}

type mockBidService struct{}

func (m *mockBidService) PlaceBid(
	ctx context.Context,
	productID,
	bidderID uuid.UUID,
	amount float64,
) (store.Bid, error) {
	return store.Bid{}, nil
}

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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusBadRequest)
	}

	var resBody map[string]any
	err := json.Unmarshal(rec.Body.Bytes(), &resBody)

	if err != nil {
		t.Fatalf("failed to parse response body:%s\n", err.Error())
	}

	if resBody["message"] != "invalid product id - must be a valid uuid" {
		t.Errorf("expected 'invalid product id - must be a valid uuid, got %q", resBody["message"])
	}

}

func TestHandleSubscribeToAuctionWithValidUUID(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()

	api := Api{
		ProductService: *services.NewProductService(&mockProductStoreHandler{}),
		Sessions:       sessionManager,
		AuctionLobby:   services.AuctionLobby{Rooms: make(map[uuid.UUID]*services.AuctionRoom)},
	}

	validUUID := uuid.New().String()
	userID := uuid.New()

	req := httptest.NewRequest("GET", "/api/v1/products/ws/subscribe/"+validUUID, nil)

	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("product_id", validUUID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))

	ctx, _ := sessionManager.Load(req.Context(), "")
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()
	handler := api.Sessions.LoadAndSave(http.HandlerFunc(api.handleSubscribeToAuction))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code == http.StatusBadRequest {
		t.Fatalf("invalid operation, expected a different code than %q", http.StatusBadRequest)
	}

	if rec.Code == http.StatusNotFound {
		t.Fatal("no product with the given id")
	}

	if rec.Code == http.StatusInternalServerError {
		t.Fatal("unexpected error, try again later")
	}

}
