package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetDocumentSummaryListRequest() GetDocumentSummaryListRequest {
	r := GetDocumentSummaryListRequest{
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

type GetDocumentSummaryListRequest struct {
	client        *Client
	queryParams   *GetDocumentSummaryListQueryParams
	pathParams    *GetDocumentSummaryListPathParams
	method        string
	headers       http.Header
	requestBody   GetDocumentSummaryListRequestBody
	requestHeader GetDocumentSummaryListRequestHeader
}

func (r GetDocumentSummaryListRequest) NewQueryParams() *GetDocumentSummaryListQueryParams {
	return &GetDocumentSummaryListQueryParams{}
}

type GetDocumentSummaryListQueryParams struct {
}

func (p GetDocumentSummaryListQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDocumentSummaryListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetDocumentSummaryListRequest) NewPathParams() *GetDocumentSummaryListPathParams {
	return &GetDocumentSummaryListPathParams{}
}

type GetDocumentSummaryListPathParams struct {
}

func (p *GetDocumentSummaryListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDocumentSummaryListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetDocumentSummaryListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDocumentSummaryListRequest) Method() string {
	return r.method
}

func (r GetDocumentSummaryListRequest) NewRequestHeader() GetDocumentSummaryListRequestHeader {
	return GetDocumentSummaryListRequestHeader{}
}

func (r *GetDocumentSummaryListRequest) RequestHeader() *GetDocumentSummaryListRequestHeader {
	return &r.requestHeader
}

func (r *GetDocumentSummaryListRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetDocumentSummaryListRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetDocumentSummaryListRequest) NewRequestBody() GetDocumentSummaryListRequestBody {
	return GetDocumentSummaryListRequestBody{}
}

type GetDocumentSummaryListRequestBody struct {
	XMLName       xml.Name     `xml:"rlx:pmsdoc_GetDocumentSummaryList"`
	SessionID     string       `xml:"rlx:SessionID"`
	FromTimestamp DateTime     `xml:"rlx:FromTimestamp"`
	ToTimestamp   DateTime     `xml:"rlx:ToTimestamp"`
	DocumentTypes DocumentType `xml:"rlx:FinancialDocumentType"`
}

func (r *GetDocumentSummaryListRequest) RequestBody() *GetDocumentSummaryListRequestBody {
	return &r.requestBody
}

func (r *GetDocumentSummaryListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetDocumentSummaryListRequest) SetRequestBody(body GetDocumentSummaryListRequestBody) {
	r.requestBody = body
}

func (r *GetDocumentSummaryListRequest) NewResponseBody() *GetDocumentSummaryListResponseBody {
	return &GetDocumentSummaryListResponseBody{}
}

type GetDocumentSummaryListResponseBody struct {
	XMLName                            xml.Name                       `xml:"pmsdoc_GetDocumentSummaryListResponse"`
	PmsintGetDocumentSummaryListResult ExceptionBlock                 `xml:"pmsdoc_GetDocumentSummaryListResult"`
	DocumentSummaryList                []FinancialDocumentSummaryItem `xml:"DocumentSummaryList>FinancialDocumentSummaryItem"`
}

func (rb GetDocumentSummaryListResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetDocumentSummaryListResult
}

func (r *GetDocumentSummaryListRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetDocumentSummaryListRequest) Do() (GetDocumentSummaryListResponseBody, error) {
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
