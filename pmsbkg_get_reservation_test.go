package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetReservations(t *testing.T) {
	req := client.NewGetReservationRequest()
	req.RequestBody().BookRef = "BK005903"
	req.RequestBody().RoomPickID = 1
	resp, err := req.Do()

	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
