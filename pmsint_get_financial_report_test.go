package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetFinancialReport(t *testing.T) {
	req := client.NewGetFinancialReportRequest()
	req.RequestBody().PeriodID = 806
	req.RequestBody().KepyoReport = false
	req.RequestBody().SelectionCriteria.REPORTTITLE = "OMNIBOOST"
	req.RequestBody().UseValidXmlFormat = true
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

