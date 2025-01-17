package net

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnupgradable(t *testing.T) {
	cases := []http.Header{
		nil,
		{},
		{"content-type": []string{"application/json"}},
		{"Upgrade": []string{"h2c"}},
	}

	for _, c := range cases {
		assert.False(t, upgradable(c))
	}
}

func TestUpgradable(t *testing.T) {
	assert.True(t, upgradable(http.Header{
		"Upgrade": []string{"websocket"},
	}))
}
