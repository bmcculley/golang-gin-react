package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setupServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/api/", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestJokeRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setupServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/api/jokes", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestJokeLikeRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setupServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Post(fmt.Sprintf("%s/api/jokes/like/4", ts.URL), "application/json", nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
