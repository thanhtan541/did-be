package tests

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thanhtan541/did-be-wp/api/configuration"
	"github.com/thanhtan541/did-be-wp/api/startup"
)

type TestApp struct {
	Port   int
	Url    string
	Client *http.Client
}

func (ta *TestApp) getPing() *http.Response {
	apiUrl := fmt.Sprintf("%s/ping", ta.Url)
	res, err := ta.Client.Get(apiUrl)
	if err != nil {
		log.Fatalf("❌ Failed to call: %v", err)
	}

	return res
}

// Spawn an isolated test sandox for each testcase
// Port 0 is special number that lets OS
// pick any avaiable port
func SpawnApp() TestApp {
	// Setup: initialize DB, server, etc.
	cfg, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	telemtryName := "test-telemetry"
	cfg.Application.Port = 0 //Set to open random port
	application, err := startup.Build(cfg, telemtryName)
	if err != nil {
		log.Fatalf("❌ Failed to build application: %v", err)
	}

	log.Printf("Listening on port %d", application.Port) // Store base URL and HTTP client
	url := fmt.Sprintf("http://127.0.0.1:%d", application.Port)
	client := &http.Client{}
	ta := TestApp{
		application.Port,
		url,
		client,
	}

	return ta
}
