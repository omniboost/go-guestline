package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetPeriodList(t *testing.T) {
	req := client.NewGetPeriodListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
