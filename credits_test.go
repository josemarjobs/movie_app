package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCreditsWithResults(t *testing.T) {
	a := assert.New(t)
	body := `{
    "cast": [{
      "id": 4534,
      "title": "Family Guy: Star wars"
    },{
      "id": 1232,
      "name":"Star Trek Guy"
    }]
  }`
	actor := Actor{}
	FakeServer(body, func() {
		err := FetchCredits(&actor)
		a.NoError(err)
		a.Equal(2, len(actor.Credits))
		a.Equal("Family Guy: Star wars", actor.Credits[0].NameOrTitle())
		a.Equal("Star Trek Guy", actor.Credits[1].NameOrTitle())
	})
}

func TestFetchCreditsWithNoResults(t *testing.T) {
	a := assert.New(t)
	body := `{
    "cast": []
  }`
	actor := Actor{}
	FakeServer(body, func() {
		err := FetchCredits(&actor)
		a.NoError(err)
		a.Equal(0, len(actor.Credits))
	})
}

func TestFilterCredits(t *testing.T) {
	a := assert.New(t)
	brad := Actor{Credits: []Credit{
		{ID: 1, Name: "Friends"},
		{ID: 2, Title: "World War Z"},
	}}
	jen := Actor{Credits: []Credit{
		{ID: 1, Name: "Friends"},
		{ID: 4, Title: "The Break Up"},
		{ID: 3, Title: "Along Came Polly"},
	}}
	actors := []Actor{brad, jen}
	credits := FilterCredits(actors)
	a.Equal(1, len(credits))
	a.Equal("Friends", credits[0].NameOrTitle())
}
