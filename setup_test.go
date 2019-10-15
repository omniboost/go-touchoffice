package touchoffice_test

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-touchoffice"
)

func client() *touchoffice.Client {
	apiKey := os.Getenv("API_KEY")
	terminalAccessKey := os.Getenv("TERMINAL_ACCESS_KEY")

	client := touchoffice.NewClient(nil)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)
	client.SetAPIKey(apiKey)
	client.SetTerminalAccessKey(terminalAccessKey)
	return client
}
