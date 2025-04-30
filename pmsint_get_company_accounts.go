package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetCompanyAccountsRequest() GetCompanyAccountsRequest {
	r := GetCompanyAccountsRequest{
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

type GetCompanyAccountsRequest struct {
	client        *Client
	queryParams   *GetCompanyAccountsQueryParams
	pathParams    *GetCompanyAccountsPathParams
	method        string
	headers       http.Header
	requestBody   GetCompanyAccountsRequestBody
	requestHeader GetCompanyAccountsRequestHeader
}

func (r GetCompanyAccountsRequest) NewQueryParams() *GetCompanyAccountsQueryParams {
	return &GetCompanyAccountsQueryParams{}
}

type GetCompanyAccountsQueryParams struct {
}

func (p GetCompanyAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompanyAccountsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetCompanyAccountsRequest) NewPathParams() *GetCompanyAccountsPathParams {
	return &GetCompanyAccountsPathParams{}
}

type GetCompanyAccountsPathParams struct {
}

func (p *GetCompanyAccountsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompanyAccountsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetCompanyAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanyAccountsRequest) Method() string {
	return r.method
}

func (r GetCompanyAccountsRequest) NewRequestHeader() GetCompanyAccountsRequestHeader {
	return GetCompanyAccountsRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetCompanyAccountsRequest) RequestHeader() *GetCompanyAccountsRequestHeader {
	return &r.requestHeader
}

func (r *GetCompanyAccountsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetCompanyAccountsRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetCompanyAccountsRequest) NewRequestBody() GetCompanyAccountsRequestBody {
	return GetCompanyAccountsRequestBody{}
}

type GetCompanyAccountsRequestBody struct {
	XMLName     xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetCompanyAccounts"`
	SessionID   string
	AccountCode string `xml:"AccountCode"`
	AccountName string `xml:"AccountName"`
}

func (r *GetCompanyAccountsRequest) RequestBody() *GetCompanyAccountsRequestBody {
	return &r.requestBody
}

func (r *GetCompanyAccountsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetCompanyAccountsRequest) SetRequestBody(body GetCompanyAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetCompanyAccountsRequest) NewResponseBody() *GetCompanyAccountsResponseBody {
	return &GetCompanyAccountsResponseBody{}
}

type GetCompanyAccountsResponseBody struct {
	XMLName                        xml.Name        `xml:GetCompanyAccountsResponse`
	PmsintGetCompanyAccountsResult ExceptionBlock  `xml:"pmsint_GetCompanyAccountsResult"`
	CompanyAccounts                CompanyAccounts `xml:"CompanyAccounts>CompanyAccounts>cpmsint_GetCompanyAccounts_CompanyAccountItem"`
}

func (rb GetCompanyAccountsResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetCompanyAccountsResult
}

func (r *GetCompanyAccountsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetCompanyAccountsRequest) Do() (GetCompanyAccountsResponseBody, error) {
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

type CompanyAccounts []CompanyAccount

type CompanyAccount struct {
	CompanyRef      string `xml:"CompanyRef"`
	CompanyName     string `xml:"CompanyName"`
	HoldStatus      string `xml:"HoldStatus"`
	CreditRemaining string `xml:"CreditRemaining"`
}
