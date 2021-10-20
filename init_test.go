package main

import (
	"net/http"
	"net/http/httptest"
	db "slot/models"
	"slot/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToDb(t *testing.T) {
	err := db.ConnectToDb()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRoute(t *testing.T) {
	router := routes.RouterSetup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Events booking app!", w.Body.String())
}
