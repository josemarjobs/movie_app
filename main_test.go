package main

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	setup()
	a := assert.New(t)
	r := bytes.NewBuffer([]byte("Peter\nLois\nn\n"))
	w := &bytes.Buffer{}

	Run(r, w)
	res := w.String()

	a.Contains(res, "You selected the following 2 actors:")
	a.Contains(res, "Peter")
	a.Contains(res, "Lois")
}
