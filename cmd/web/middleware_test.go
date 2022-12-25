package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("TestNoSurf() failed: type is not http.Handler, got %T instead", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("TestSessionLoad() failed: type is not http.Handler, got %T instead", v)
	}
}
