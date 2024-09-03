package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetRoomStatus(t *testing.T) {
	req := client.NewGetRoomStatusRequest()
	req.RequestBody().TheRequest = VoidPOSCheckRequest {}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
