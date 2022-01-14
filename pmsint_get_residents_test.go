package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	guestline "github.com/omniboost/go-guestline"
)

func TestGetResidents(t *testing.T) {
	req := client.NewGetResidentsRequest()
	req.RequestBody().FromDate = guestline.Time{time.Date(2021, 10, 1, 0, 0, 0, 0, time.Local)}
	req.RequestBody().ToDate = guestline.Time{time.Date(2021, 11, 1, 0, 0, 0, 0, time.Local)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
