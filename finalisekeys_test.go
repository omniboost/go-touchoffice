package touchoffice_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/omniboost/go-touchoffice"
)

func TestFinalisekeys(t *testing.T) {
	client := client()
	req := client.NewFinalisekeysRequest()
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterday = time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	req.QueryParams().Site = 1
	req.QueryParams().DateStart = touchoffice.Date{yesterday}
	req.QueryParams().TimeStart = touchoffice.Time{yesterday}
	req.QueryParams().DateEnd = touchoffice.Date{today}
	req.QueryParams().TimeEnd = touchoffice.Time{today}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
