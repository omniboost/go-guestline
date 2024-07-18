package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/omniboost/go-guestline"
)

func TestGetDocumentSummaryList(t *testing.T) {
	req := client.NewGetDocumentSummaryListRequest()
	req.RequestBody().FromTimestamp = guestline.DateTime{time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().ToTimestamp = guestline.DateTime{time.Date(2024, 5, 30, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().DocumentTypes = guestline.DocumentTypeInvoicesOnly
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
