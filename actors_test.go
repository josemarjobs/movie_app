package main

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestFecthActorWithResults(t *testing.T) {
	a := assert.New(t)
	body := `{
    "page": 1,
    "results": [{
      "id": 239,
      "name": "Brad Pitt",
      "popularity": 12,
      "profile_path": "/brad.jpg"
    }],
    "total_pages": 1,
    "total_results": 1
  }`

	FakeServer(body, func() {
		actor, err := FetchActor("Brad Pitt")
		a.NoError(err)
		a.Equal("Brad Pitt", actor.Name)
	})
}

func TestFetchActorWithNoResults(t *testing.T) {
	a := assert.New(t)
	body := `{
    "page": 1,
    "results": [],
    "total_pages": 1,
    "total_results": 0
  }`
	FakeServer(body, func() {
		_, err := FetchActor("Brad Pitt")
		a.Error(err)
		a.Equal("There are no search results for: Brad Pitt!", err.Error())
	})
}

func FakeServer(body string, f func()) {
	root := ApitRoot

	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, body)
		}))
	defer testServer.Close()

	ApitRoot = testServer.URL
	f()

	ApitRoot = root
}
