package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	guestline "github.com/omniboost/go-guestline"
)

func TestGetArrivals(t *testing.T) {
	req := client.NewGetArrivalsRequest()
	req.RequestBody().ArrivalDate = guestline.DateTime{time.Date(2021, 6, 2, 0, 0, 0, 0, time.Local)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
