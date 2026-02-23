package api

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/services"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

type mockProductStore struct{}

func (m *mockProductStore) CreateProduct(
	ctx context.Context,
	sellerID uuid.UUID,
	productName,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (m *mockProductStore) GetProductById(ctx context.Context, productID uuid.UUID) (store.Product, error) {
	return store.Product{}, nil
}

func TestCreateProduct(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()

	api := Api{
		ProductService: *services.NewProductService(&mockProductStore{}),
		Sessions:       sessionManager,
		AuctionLobby: services.AuctionLobby{
			Rooms: make(map[uuid.UUID]*services.AuctionRoom),
		},
	}

	auctionEnd := time.Now().Add(3 * time.Hour).UTC()

	payLoad := map[string]any{
		"product_name": "produto de teste",
		"description":  "descrição de teste",
		"baseprice":    100.00,
		"auction_end":  auctionEnd.Format(time.RFC3339),
	}

	body, err := json.Marshal(payLoad)
	if err != nil {
		t.Fatal("fail to parse request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/products/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")

	userID := uuid.New()
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleCreateProduct))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusCreated)
	}

	var resBody map[string]any
	err = json.Unmarshal(rec.Body.Bytes(), &resBody)
	if err != nil {
		t.Fatalf("failed to parse response body:%s\n", err.Error())
	}

	productIDstr, ok := resBody["id"].(string)
	if !ok {
		t.Fatalf("id is not a string: %q", resBody["id"])
	}

	productID, err := uuid.Parse(productIDstr)
	if err != nil {
		t.Fatalf("invalid uuid: %q", productIDstr)
	}

	if productID == uuid.Nil {
		t.Errorf("productID cannot be nil")
	}

}
