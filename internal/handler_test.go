package internal_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	h "../internal"
	l "./link"
)

func TestPosthandler(t *testing.T) {
	t.Run("create classic link - success", func(t *testing.T) {
		req := request(t, strings.NewReader("{\n\t\"type\": \"classic\",\n\t\"data\": {\n\t\t\"title\": \"Test\",\n\t\t\"url\": \"https://example.com\"\n\t}\n}"))
		rr := httptest.NewRecorder()
		r := mux.NewRouter()
		handler := h.NewHandler(*l.NewService())
		r.HandleFunc("/v1/users/{id}", handler.PostHandler).Methods("POST")
		r.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
		assert.NotNil(t, rr.Body)
	})

	t.Run("create classic link - bad request - invalid link type", func(t *testing.T) {
		req := request(t, strings.NewReader("{\n\t\"type\": \"blah\",\n\t\"data\": {\n\t\t\"title\": \"Test\",\n\t\t\"url\": \"https://example.com\"\n\t}\n}"))
		rr := httptest.NewRecorder()
		r := mux.NewRouter()
		handler := h.NewHandler(*l.NewService())
		r.HandleFunc("/v1/users/{id}", handler.PostHandler).Methods("POST")
		r.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, rr.Body.String(), "\"invalid link type\"")
	})
}

// TODO Add tests for GET handler, even though it's a dummy implementation

func request(t *testing.T, reader io.Reader) *http.Request {
	req, err := http.NewRequest("POST", "/v1/users/1234", reader)
	if err != nil {
		t.Fatalf("problem creating request: %+v", err)
	}
	return req
}
