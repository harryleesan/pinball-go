package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		t.Errorf("Cannot GET https://google.com %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("GET https://google.com returned: %v instead of 200", resp.StatusCode)
	}
}
