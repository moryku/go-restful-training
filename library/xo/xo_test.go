package xo_test

import (
	"testing"

	"github.com/iswanulumam/go-restful-training/library/xo"
	"github.com/stretchr/testify/assert"
)

func TestXo(t *testing.T) {
	assert.Equal(t, xo.CountXO("xoxoxo"), true, "they should be equal")
	assert.Equal(t, xo.CountXO("oxooxo"), false, "they sould be equal")
	assert.Equal(t, xo.CountXO("oxo"), false, "they sould be equal")
	assert.Equal(t, xo.CountXO("xxxooo"), true, "they sould be equal")
	assert.Equal(t, xo.CountXO("xoxooxxo"), true, "they sould be equal")
}
