package main

import (
	"testing"

	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing, test passed
	default:
		t.Errorf("TestRoutes() failed: type is not *chi.Mux, got %T instead", v)
	}
}
