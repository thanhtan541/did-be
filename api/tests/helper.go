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

func SpawnApp() TestApp {
	// Setup: initialize DB, server, etc.
	cfg, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	cfg.Application.Port = 0 //Set to open random port
	application, err := startup.Build(cfg)
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
