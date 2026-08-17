package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestProbeEmptyBlobNameReturns400 asserts that GET /blobs/ (no blob name) is
// rejected with HTTP 400, not silently treated as a lookup of an empty-named
// blob (which would yield 404).
func TestProbeEmptyBlobNameReturns400(t *testing.T) {
	_, h := newTestServer()
	req := httptest.NewRequest(http.MethodGet, "/blobs/", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("GET /blobs/ status = %d want 400", rec.Code)
	}
}
