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
      "name": "Peter Griffin",
      "popularity": 12,
      "profile_path": "/peter.jpg"
    }],
    "total_pages": 1,
    "total_results": 1
  }`

	FakeServer(body, func() {
		actor, err := FetchActor("Peter Griffin")
		a.NoError(err)
		a.Equal("Peter Griffin", actor.Name)
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
		_, err := FetchActor("Peter Griffin")
		a.Error(err)
		a.Equal("There are no search results for: Peter Griffin!", err.Error())
	})
}

func FakeServer(body string, f func()) {
	root := ApiRoot

	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, body)
		}))
	defer testServer.Close()

	ApiRoot = testServer.URL
	f()

	ApiRoot = root
}
