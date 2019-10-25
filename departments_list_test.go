package touchoffice_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDepartmentsList(t *testing.T) {
	client := client()
	req := client.NewDepartmentsListRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
