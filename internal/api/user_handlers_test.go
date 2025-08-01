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

type mockUserStore struct{}

func (m *mockUserStore) CreateUser(
	ctx context.Context,
	userName,
	email string,
	password []byte,
	bio string) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (m *mockUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (m *mockUserStore) GetUserByEmail(ctx context.Context, email string) (store.User, error) {
	return store.User{}, nil
}

func (m *mockUserStore) GetUserById(ctx context.Context, id uuid.UUID) (store.User, error) {
	return store.User{}, nil
}

func TestSignupUser(t *testing.T) {

	api := Api{
		UserService: *services.NewUserService(&mockUserStore{}),
	}

	payload := map[string]any{
		"user_name": "marcos",
		"email":     "marcos.santana@gmail.com",
		"password":  "Marcos@2025",
		"bio":       "testing my api",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("fail to parse request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/users/signupuser", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.handleSignupUser)
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

	if _, ok := resBody["user_id"]; !ok {
		t.Errorf("expected 'user_id' in response, got %q", resBody)
	}

}

func TestLoginUser(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		UserService: *services.NewUserService(&mockUserStore{}),
		Sessions:    sessionManager,
	}

	payLoad := map[string]any{
		"email":    "marcos.santana@gmail.com",
		"password": "Marcos@2025",
	}

	body, err := json.Marshal(payLoad)
	if err != nil {
		t.Fatal("fail to parse request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/users/loginuser", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := api.Sessions.LoadAndSave(http.HandlerFunc(api.handleLoginUser))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusOK {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusOK)
	}

	var resBody map[string]any
	err = json.Unmarshal(rec.Body.Bytes(), &resBody)
	if err != nil {
		t.Fatalf("failed to parse response body:%s\n", err.Error())
	}

	if resBody["message"] != "logged in sucessfully" {
		t.Errorf("message differs; got %q | want 'logged in sucessfully'", resBody["message"])
	}

}

func TestLogoutUser(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		Sessions: sessionManager,
	}

	req := httptest.NewRequest("POST", "/api/v1/users/logoutuser", nil)
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")
	sessionManager.Put(ctx, "AuthenticateUserId", uuid.New())

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleLogoutUser))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status differs; got %q | want %q", rec.Code, http.StatusOK)
	}

	var resBody map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to parse response body:%s\n", err.Error())
	}

	cookie := rec.Header().Get("Set-cookie")
	newReq := httptest.NewRequest("POST", "/dummy", nil)
	newReq.Header.Set("Cookie", cookie)

	newCtx, _ := sessionManager.Load(newReq.Context(), "")
	authUser := sessionManager.Get(newCtx, "AuthenticateUserId")

	if authUser != nil {
		t.Errorf("expected session value to be removed, but got: %v", authUser)
	}

}
