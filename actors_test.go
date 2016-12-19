package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFecthActorWithResults(t *testing.T) {
	a := assert.New(t)

	actor, err := FetchActor("Brad Pitt")
	a.NoError(err)
	a.Equal("Brad Pitt", actor.Name)
}

func TestFetchActorWithNoResults(t *testing.T) {
	a := assert.New(t)

	_, err := FetchActor("Brad Pitt")
	a.Error(err)
	a.Equal("There are no search results for: Brad Pitt!", err.Error())
}
