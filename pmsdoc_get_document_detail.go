package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetDocumentDetailRequest() GetDocumentDetailRequest {
	r := GetDocumentDetailRequest{
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

type GetDocumentDetailRequest struct {
	client        *Client
	queryParams   *GetDocumentDetailQueryParams
	pathParams    *GetDocumentDetailPathParams
	method        string
	headers       http.Header
	requestBody   GetDocumentDetailRequestBody
	requestHeader GetDocumentDetailRequestHeader
}

func (r GetDocumentDetailRequest) NewQueryParams() *GetDocumentDetailQueryParams {
	return &GetDocumentDetailQueryParams{}
}

type GetDocumentDetailQueryParams struct {
}

func (p GetDocumentDetailQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDocumentDetailRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetDocumentDetailRequest) NewPathParams() *GetDocumentDetailPathParams {
	return &GetDocumentDetailPathParams{}
}

type GetDocumentDetailPathParams struct {
}

func (p *GetDocumentDetailPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDocumentDetailRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetDocumentDetailRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDocumentDetailRequest) Method() string {
	return r.method
}

func (r GetDocumentDetailRequest) NewRequestHeader() GetDocumentDetailRequestHeader {
	return GetDocumentDetailRequestHeader{}
}

func (r *GetDocumentDetailRequest) RequestHeader() *GetDocumentDetailRequestHeader {
	return &r.requestHeader
}

func (r *GetDocumentDetailRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetDocumentDetailRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetDocumentDetailRequest) NewRequestBody() GetDocumentDetailRequestBody {
	return GetDocumentDetailRequestBody{}
}

type GetDocumentDetailRequestBody struct {
	XMLName     xml.Name `xml:"rlx:pmsdoc_GetDocumentDetail"`
	SessionID   string   `xml:"rlx:SessionID"`
	DocumentRef string   `xml:"rlx:DocumentRef"`
}

func (r *GetDocumentDetailRequest) RequestBody() *GetDocumentDetailRequestBody {
	return &r.requestBody
}

func (r *GetDocumentDetailRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetDocumentDetailRequest) SetRequestBody(body GetDocumentDetailRequestBody) {
	r.requestBody = body
}

func (r *GetDocumentDetailRequest) NewResponseBody() *GetDocumentDetailResponseBody {
	return &GetDocumentDetailResponseBody{}
}

type GetDocumentDetailResponseBody struct {
	XMLName                       xml.Name       `xml:"pmsdoc_GetDocumentDetailResponse"`
	PmsintGetDocumentDetailResult ExceptionBlock `xml:"pmsdoc_GetDocumentDetailResult"`
	DocumentDetail                DocumentDetail `xml:"DocumentDetail"`
}

func (rb GetDocumentDetailResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetDocumentDetailResult
}

func (r *GetDocumentDetailRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetDocumentDetailRequest) Do() (GetDocumentDetailResponseBody, error) {
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
