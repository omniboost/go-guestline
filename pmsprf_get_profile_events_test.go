package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetprofileEvents(t *testing.T) {
	req := client.NewGetProfileEventsRequest()
	req.RequestBody().TransID = 26680
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
