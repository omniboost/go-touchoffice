package touchoffice_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestSitesList(t *testing.T) {
	client := client()
	req := client.NewSitesListRequest()

	resp, err := req.Do()
	if err != nil {
		log.Println(err == nil)
		log.Println(err)
		os.Exit(99)
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
