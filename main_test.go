package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)

	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200, but got %d", w.Code)
	}
}
