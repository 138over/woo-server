package main
  
import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandlers(t *testing.T) {
    t.Run("ping", func(t *testing.T) {
        request, err := http.NewRequest("GET", "/ping", nil)
        if err != nil {
            t.Fatal(err)
        }

        recorder := httptest.NewRecorder()
        handler := http.HandlerFunc(ping)

        handler.ServeHTTP(recorder, request)
        if status := recorder.Code; status != http.StatusOK {
            t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
        }

        contentType := recorder.Header().Get("Content-Type")
        if contentType != "application/json" {
            t.Errorf("content type does not match: got %v want %v", contentType, "application/json")
        }

        expected := `{"alive":"true"}`
        if recorder.Body.String() != expected {
            t.Errorf("unexpected body: got %v want %v", recorder.Body.String(), expected)
        }

    })
}

