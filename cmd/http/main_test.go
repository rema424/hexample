package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	// Setup
	e := newEcho()
	dbx := newDBx()
	dbxx := newDBxx(dbx)
	handler := newHandler(e, dbx, dbxx)
	s := httptest.NewServer(handler)
	defer s.Close()

	// Request and Assertions
	res, err := http.Get(s.URL + "/apple")
	if err != nil {
		t.Fatalf("http.Get failed: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("res.StatusCode: got: %d, want: %d", res.StatusCode, http.StatusOK)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("ioutil.ReadAll failed: %s", err)
	}
	got := string(body)
	want := "apple"
	if got != want {
		t.Fatalf("request: /apple, got %s, want %s", got, want)
	}
}
