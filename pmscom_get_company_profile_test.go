package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetCompanyProfile(t *testing.T) {
	req := client.NewGetCompanyProfileRequest()
	req.RequestBody().CompanyRef = "GROUP_2021"
	// req.RequestBody().Filters.RoomPickID = "2"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

