package handlers

import (
	"encoding/json"
	"net/http"
	"Backend/pkg/models"
)

var photos = []models.Photo{
	{ID: "1", Name: "Photo 1", URL: "http://example.com/1.jpg"}
	{}
}