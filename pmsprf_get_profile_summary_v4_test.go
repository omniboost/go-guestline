package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetProfileSummaryV4(t *testing.T) {
	req := client.NewGetProfileSummaryV4Request()
	req.RequestBody().ProfileRequestor.ProfileUniqueID = "PF1031991"
	req.RequestBody().ProfileRequestor.AuthenticationMethod = "PD"
	req.RequestBody().ProfileRequestor.AuthenticationCode = "Surname"
	req.RequestBody().ProfileRequestor.AuthenticationValue = "Castor"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
