package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetFinancialReportRequest() GetFinancialReportRequest {
	r := GetFinancialReportRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	r.requestHeader = r.NewRequestHeader()
	return r
}

type GetFinancialReportRequest struct {
	client        *Client
	queryParams   *GetFinancialReportQueryParams
	pathParams    *GetFinancialReportPathParams
	method        string
	headers       http.Header
	requestBody   GetFinancialReportRequestBody
	requestHeader GetFinancialReportRequestHeader
}

func (r GetFinancialReportRequest) NewQueryParams() *GetFinancialReportQueryParams {
	return &GetFinancialReportQueryParams{}
}

type GetFinancialReportQueryParams struct {
}

func (p GetFinancialReportQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetFinancialReportRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetFinancialReportRequest) NewPathParams() *GetFinancialReportPathParams {
	return &GetFinancialReportPathParams{}
}

type GetFinancialReportPathParams struct {
}

func (p *GetFinancialReportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetFinancialReportRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetFinancialReportRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFinancialReportRequest) Method() string {
	return r.method
}

func (r GetFinancialReportRequest) NewRequestHeader() GetFinancialReportRequestHeader {
	return GetFinancialReportRequestHeader{}
}

func (r *GetFinancialReportRequest) RequestHeader() *GetFinancialReportRequestHeader {
	return &r.requestHeader
}

func (r *GetFinancialReportRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetFinancialReportRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetFinancialReportRequest) NewRequestBody() GetFinancialReportRequestBody {
	return GetFinancialReportRequestBody{
		SelectionCriteria: SelectionCriteria{},
	}
}

type GetFinancialReportRequestBody struct {
	XMLName           xml.Name          `xml:"rlx:pmsint_GetFinancialReport"`
	SessionID         string            `xml:"rlx:SessionID"`
	PeriodID          int               `xml:"rlx:PeriodID"`
	SelectionCriteria SelectionCriteria `xml:"rlx:SelectionCriteria"`
	KepyoReport       bool              `xml:"rlx:KepyoReport"`
	UseValidXmlFormat bool              `xml:"rlx:UseValidXmlFormat"`
}

func (r *GetFinancialReportRequest) RequestBody() *GetFinancialReportRequestBody {
	return &r.requestBody
}

func (r *GetFinancialReportRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetFinancialReportRequest) SetRequestBody(body GetFinancialReportRequestBody) {
	r.requestBody = body
}

func (r *GetFinancialReportRequest) NewResponseBody() *GetFinancialReportResponseBody {
	return &GetFinancialReportResponseBody{}
}

type GetFinancialReportResponseBody struct {
	XMLName                        xml.Name            `xml:"pmsint_GetFinancialReportResponse"`
	PmsintGetFinancialReportResult ExceptionBlock      `xml:"pmsint_GetFinancialReportResult"`
	Data                           FinancialReportData `xml:"GetFinancialReport>Data"`
}

func (rb GetFinancialReportResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetFinancialReportResult
}

func (r *GetFinancialReportRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetFinancialReportRequest) Do() (GetFinancialReportResponseBody, error) {
	var err error

	// fetch a new token if it isn't set already
	r.requestBody.SessionID, err = r.client.SessionID()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
