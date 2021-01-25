package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetProfileSummaryV3(t *testing.T) {
	req := client.NewGetProfileSummaryV3Request()
	req.RequestBody().ProfileRequestor.ProfileUniqueID = "PF003785"
	req.RequestBody().ProfileRequestor.AuthenticationMethod = "PD"
	req.RequestBody().ProfileRequestor.AuthenticationCode = "Surname"
	req.RequestBody().ProfileRequestor.AuthenticationValue = "Ondruch"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
