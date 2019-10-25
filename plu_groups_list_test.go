package touchoffice_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPLUGroupsList(t *testing.T) {
	client := client()
	req := client.NewPLUGroupsListRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
