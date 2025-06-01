package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Decide how API is structure
type PingResponse struct {
	Message string `json:"message"`
}

func TestPing(t *testing.T) {
	app := SpawnApp()
	resp := app.getPing()

	assert.Equal(t, 200, resp.StatusCode, "Request to Ping not sucessful")

	defer resp.Body.Close()

	var data PingResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("❌ Failed to json decode body: %v", err)
	}

	assert.Equal(t, "pong", data.Message, "Body has incorrect format")
}
