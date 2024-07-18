package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetDocumentDetail(t *testing.T) {
	req := client.NewGetDocumentDetailRequest()
	req.RequestBody().DocumentRef = "FB083562"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
