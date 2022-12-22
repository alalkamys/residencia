package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// run tests before exiting
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
