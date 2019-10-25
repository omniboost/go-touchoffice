package touchoffice_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPLUList2(t *testing.T) {
	client := client()
	req := client.NewPLUList2Request()

	req.QueryParams().Site = 1

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
