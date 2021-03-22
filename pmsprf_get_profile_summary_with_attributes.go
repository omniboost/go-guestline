package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetProfileSummaryWithAttributesRequest() GetProfileSummaryWithAttributesRequest {
	r := GetProfileSummaryWithAttributesRequest{
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

type GetProfileSummaryWithAttributesRequest struct {
	client        *Client
	queryParams   *GetProfileSummaryWithAttributesQueryParams
	pathParams    *GetProfileSummaryWithAttributesPathParams
	method        string
	headers       http.Header
	requestBody   GetProfileSummaryWithAttributesRequestBody
	requestHeader GetProfileSummaryWithAttributesRequestHeader
}

func (r GetProfileSummaryWithAttributesRequest) NewQueryParams() *GetProfileSummaryWithAttributesQueryParams {
	return &GetProfileSummaryWithAttributesQueryParams{}
}

type GetProfileSummaryWithAttributesQueryParams struct {
}

func (p GetProfileSummaryWithAttributesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetProfileSummaryWithAttributesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetProfileSummaryWithAttributesRequest) NewPathParams() *GetProfileSummaryWithAttributesPathParams {
	return &GetProfileSummaryWithAttributesPathParams{}
}

type GetProfileSummaryWithAttributesPathParams struct {
}

func (p *GetProfileSummaryWithAttributesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetProfileSummaryWithAttributesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetProfileSummaryWithAttributesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetProfileSummaryWithAttributesRequest) Method() string {
	return r.method
}

func (r GetProfileSummaryWithAttributesRequest) NewRequestHeader() GetProfileSummaryWithAttributesRequestHeader {
	return GetProfileSummaryWithAttributesRequestHeader{}
}

func (r *GetProfileSummaryWithAttributesRequest) RequestHeader() *GetProfileSummaryWithAttributesRequestHeader {
	return &r.requestHeader
}

func (r *GetProfileSummaryWithAttributesRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetProfileSummaryWithAttributesRequestHeader struct{}

func (r GetProfileSummaryWithAttributesRequest) NewRequestBody() GetProfileSummaryWithAttributesRequestBody {
	return GetProfileSummaryWithAttributesRequestBody{
		ProfileRequestor: ProfileRequestor{
			AuthenticationMethod: "PD",
		},
	}
}

type GetProfileSummaryWithAttributesRequestBody struct {
	XMLName          xml.Name         `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsprf_GetProfileSummaryWithAttributes"`
	SessionID        string           `xml:"SessionId"`
	ProfileRequestor ProfileRequestor `xml:"ProfileRequestor,omitempty"`
}

func (r *GetProfileSummaryWithAttributesRequest) RequestBody() *GetProfileSummaryWithAttributesRequestBody {
	return &r.requestBody
}

func (r *GetProfileSummaryWithAttributesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetProfileSummaryWithAttributesRequest) SetRequestBody(body GetProfileSummaryWithAttributesRequestBody) {
	r.requestBody = body
}

func (r *GetProfileSummaryWithAttributesRequest) NewResponseBody() *GetProfileSummaryWithAttributesResponseBody {
	return &GetProfileSummaryWithAttributesResponseBody{}
}

type GetProfileSummaryWithAttributesResponseBody struct {
	XMLName                                     xml.Name                `xml:GetProfileSummaryWithAttributesResponse`
	PmsintGetProfileSummaryWithAttributesResult ExceptionBlock          `xml:"pmsprf_GetProfileSummaryWithAttributesResult"`
	ProfileSummary                              ProfileSummary          `xml:"Profile>Profile"`
	CustomAttributes                            ProfileCustomAttributes `xml:"Profile>CustomAttributes>ProfileCustomAttributes"`
}

func (rb GetProfileSummaryWithAttributesResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetProfileSummaryWithAttributesResult
}

func (r *GetProfileSummaryWithAttributesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx?op=pmsprf_GetProfileSummaryWithAttributes", r.PathParams())
	return &u
}

func (r *GetProfileSummaryWithAttributesRequest) Do() (GetProfileSummaryWithAttributesResponseBody, error) {
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

type ProfileCustomAttributes []ProfileCustomAttribute

type ProfileCustomAttribute struct {
	AttributeCode               string `xml:"AttributeCode"`
	Value                       string `xml:"Value"`
	Param1                      string `xml:"Param1"`
	Param2                      string `xml:"Param2"`
	ProfileAttributeCode        string `xml:"ProfileAttributeCode"`
	ProfileAttributeDescription string `xml:"ProfileAttributeDescription"`
	ProfileAttributeValue       string `xml:"ProfileAttributeValue"`
}
