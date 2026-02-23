package api

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

type mockBidStore struct{}

func (m *mockBidStore) CreateBid(
	ctx context.Context,
	productID,
	bidderID uuid.UUID,
	bidAmount float64,
) (store.Bid, error) {
	return store.Bid{
		ID:        uuid.New(),
		ProductID: productID,
		BidderID:  bidderID,
		BidAmount: bidAmount,
	}, nil
}

func (m *mockBidStore) GetBidsByProduct(ctx context.Context, productID uuid.UUID) ([]store.Bid, error) {
	return []store.Bid{}, nil
}

func (m *mockBidStore) GetHighestBidByProductId(ctx context.Context, productID uuid.UUID) (store.Bid, error) {
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
		BidService:     services.BidService{Store: &mockBidStore{}},
	}

	validUUID := uuid.New().String()
	productID := uuid.MustParse(validUUID)
	userID := uuid.New()

	api.AuctionLobby.Rooms[productID] = &services.AuctionRoom{
		Id:          productID,
		Context:     context.Background(),
		Broadcast:   make(chan services.Message),
		Register:    make(chan *services.Client),
		Unregister:  make(chan *services.Client),
		Clients:     make(map[uuid.UUID]*services.Client),
		BidsService: &api.BidService,
	}

	router := chi.NewRouter()

	router.Get("/api/v1/products/ws/subscribe/{product_id}", api.handleSubscribeToAuction)

	srv := httptest.NewServer(api.Sessions.LoadAndSave(router))
	defer srv.Close()

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http") + "/api/v1/products/ws/subscribe/" + validUUID

	dialer := websocket.DefaultDialer

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	token, expiry, err := sessionManager.Commit(ctx)
	if err != nil {
		t.Fatal(err)
	}
	sessionManager.WriteSessionCookie(ctx, w, token, expiry)

	cookie := w.Result().Cookies()[0].String()

	header := http.Header{}
	header.Add("Cookie", cookie)

	conn, resp, err := dialer.Dial(wsURL, header)
	if err != nil {
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("failed to connect websocket: %v | status: %d | body: %s",
				err, resp.StatusCode, string(body))
		}
		t.Fatalf("failed to connect websocket: %v", err)
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte("hello auction"))
	if err != nil {
		t.Fatalf("fail to write message: %v", err)
	}

}
