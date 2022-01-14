package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestRoomTypeList(t *testing.T) {
	req := client.NewRoomTypeListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
