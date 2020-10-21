package http

import (
	chatapp "github.com/cavke/go-chat-app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cavke/go-chat-app/mock"
)

func TestHandler(t *testing.T) {
	// Inject our mock into our handler.
	var us mock.UserService
	var h Handler
	h.UserService = &us

	// Mock our User() call.
	us.UserFn = func(id int) (*chatapp.User, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &chatapp.User{ID: 100, Name: "susy"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/100", nil)
	h.ServeHTTP(w, r)

	// Validate mock.
	if !us.UserInvoked {
		t.Fatal("expected User() to be invoked")
	}
}
