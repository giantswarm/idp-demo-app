// Golang REST API unit testing program
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func setup() {
	if router == nil {
		router = setupRouter(storeMemory)
	}
}

func TestPingRoute(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong\n", w.Body.String())
}
func TestCreateAlbum(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	reqBody := strings.NewReader(`{"title":"New Album","artist":"Artist Name","price":9.99}`)
	req, _ := http.NewRequest(http.MethodPost, "/albums", reqBody)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
}

func TestDeleteAlbum(t *testing.T) {
	setup()

	// Use the ID of a test album
	albumID := "2"

	// Delete the album
	deleteReq, _ := http.NewRequest(http.MethodDelete, "/albums/"+albumID, nil)
	deleteResp := httptest.NewRecorder()
	router.ServeHTTP(deleteResp, deleteReq)

	assert.Equal(t, http.StatusFound, deleteResp.Code)
}

func TestListAlbums(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAlbumByID(t *testing.T) {
	setup()

	albumID := "1"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/albums/"+albumID, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateAlbum(t *testing.T) {
	setup()

	// Use the ID of a test album
	albumID := "1"

	// Update the album with a new artist value
	reqBody := strings.NewReader(`{"id": 1, "artist":"New Artist Name", "title":"New Album", "price":9.99}`)
	req, _ := http.NewRequest(http.MethodPut, "/albums/"+albumID, reqBody)
	req.Header.Set("Content-Type", "application/json")
	updateResp := httptest.NewRecorder()
	router.ServeHTTP(updateResp, req)

	assert.Equal(t, http.StatusFound, updateResp.Code)
}
