package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test Suite for connector Controller
func TestController(t *testing.T) {

	// Setup code goes here

	t.Run("check HealthCheck endpoint", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/healthcheck", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthCheck)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.string(resp)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		resp, _ := ioutil.ReadAll(rr.Body)
		t.Log(strings.TrimSuffix(string(resp), "\n") == `{"title":"Server is up!"}`)

	})
	// TearDown code goes here

}
