package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSignupUser(t *testing.T) {
	api := Api{}

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

	if !reflect.DeepEqual(payload, resBody) {
		t.Errorf("response body differs from payload;got: %q | want %q", resBody, payload)
	}

}
