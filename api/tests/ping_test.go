package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	app := SpawnApp()
	resp := app.getPing()

	assert.Equal(t, 200, resp.StatusCode, "Request to Ping not sucessful")
}
