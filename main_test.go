package main

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	setup()
	a := assert.New(t)
	r := bytes.NewBuffer([]byte("Brad Pitt\nJennifer Aniston\nn\n"))
	w := &bytes.Buffer{}

	Run(r, w)
	res := w.String()

	a.Contains(res, "You selected the following 2 actors:")
	a.Contains(res, "Brad")
	a.Contains(res, "Jennifer")
	a.Contains(res, "Friends")
}
