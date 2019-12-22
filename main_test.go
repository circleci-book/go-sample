package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"HelloWorld\"}", w.Body.String())
}

func TestFetchAllProducts(t *testing.T) {
	router := setupRouter()

	var err error

	db, err = setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Create(&Product{Code: "L1212", Price: 1000})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	dropDB()
}
