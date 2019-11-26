package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandleMessage(t *testing.T) {
	// Setup
	e := echo.New()
	ctrl := Controller{}

	cases := []struct {
		in, out string
	}{
		{"alice", "alice"},
		{"bob", "bob"},
		{"carol", "carol"},
		{"dave", "dave"},
	}

	for i, cs := range cases {
		// Request
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:message")
		c.SetParamNames("message")
		c.SetParamValues(cs.in)

		// Assertions
		err := ctrl.HandleMessage(c)
		if err != nil {
			t.Fatalf("#%d: HandleGreet failed: %s", i, err)
		}
		if rec.Code != http.StatusOK {
			t.Fatalf("#%d: response code is not 200", i)
		}
		got, want := rec.Body.String(), cs.out
		if got != want {
			t.Fatalf("#%d: request: /%s, got: %s, want: %s", i, cs.in, got, want)
		}
	}
}
