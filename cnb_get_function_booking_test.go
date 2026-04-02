package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestFunctionBooking(t *testing.T) {
	req := client.NewGetFunctionBookingRequest()
	req.RequestBody().BookRef = "AMSM9034575"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
