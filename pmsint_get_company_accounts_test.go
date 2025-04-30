package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetCompanyAccounts(t *testing.T) {
	req := client.NewGetCompanyAccountsRequest()
	// req.RequestBody().AccountCode = "BDBEV"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

