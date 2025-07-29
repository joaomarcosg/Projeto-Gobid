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
	}

	payLoad := map[string]any{
		"product_name": "produto de teste",
		"description":  "descrição de teste",
		"baseprice":    100.00,
		"auction_end":  "2025-08-30T15:00:00Z",
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

	if resBody["id"] != "id" {
		t.Errorf("id differs; got %q | want 'id'", resBody["id"])
	}

}
