package guestline_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/omniboost/go-guestline"
)

func TestGetPeriodList(t *testing.T) {
	req := client.NewGetPeriodListRequest()
	req.RequestBody().IPeriodID = 1
	req.RequestBody().EnmPeriodType = guestline.PeriodTypeEndOfDay
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
