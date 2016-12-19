package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskForName(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Peter\n")

	r := bytes.NewBuffer(b)
	AskForName(r)

	a.Equal(len(ActorNames), 1)
	a.Equal(ActorNames[0], "Peter")
}

func TestAskForNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Peter\nLois\nn\n")

	r := bytes.NewBuffer(b)
	AskForNames(r)

	a.Equal(len(ActorNames), 2)
	a.Equal(ActorNames[0], "Peter")
	a.Equal(ActorNames[1], "Lois")
}

func TestAskForMoreThan2Names(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Peter\nLois\ny\nStewie\ny\nMaggie\nn\n")

	r := bytes.NewBuffer(b)
	AskForNames(r)

	a.Equal(len(ActorNames), 4)
	a.Equal(ActorNames[0], "Peter")
	a.Equal(ActorNames[1], "Lois")
	a.Equal(ActorNames[2], "Stewie")
	a.Equal(ActorNames[3], "Maggie")
}

func setup() {
	ActorNames = []string{}
}
