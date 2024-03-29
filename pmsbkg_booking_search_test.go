package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestBookingSearch(t *testing.T) {
	req := client.NewBookingSearchRequest()
	req.RequestBody().Filters.BookRef = "HWR000051"
	// req.RequestBody().Filters.RoomPickID = "2"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
