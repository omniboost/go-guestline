package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetCompanyProfileRequest() GetCompanyProfileRequest {
	r := GetCompanyProfileRequest{
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

type GetCompanyProfileRequest struct {
	client        *Client
	queryParams   *GetCompanyProfileQueryParams
	pathParams    *GetCompanyProfilePathParams
	method        string
	headers       http.Header
	requestBody   GetCompanyProfileRequestBody
	requestHeader GetCompanyProfileRequestHeader
}

func (r GetCompanyProfileRequest) NewQueryParams() *GetCompanyProfileQueryParams {
	return &GetCompanyProfileQueryParams{}
}

type GetCompanyProfileQueryParams struct {
}

func (p GetCompanyProfileQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompanyProfileRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetCompanyProfileRequest) NewPathParams() *GetCompanyProfilePathParams {
	return &GetCompanyProfilePathParams{}
}

type GetCompanyProfilePathParams struct {
}

func (p *GetCompanyProfilePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompanyProfileRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetCompanyProfileRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanyProfileRequest) Method() string {
	return r.method
}

func (r GetCompanyProfileRequest) NewRequestHeader() GetCompanyProfileRequestHeader {
	return GetCompanyProfileRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetCompanyProfileRequest) RequestHeader() *GetCompanyProfileRequestHeader {
	return &r.requestHeader
}

func (r *GetCompanyProfileRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetCompanyProfileRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetCompanyProfileRequest) NewRequestBody() GetCompanyProfileRequestBody {
	return GetCompanyProfileRequestBody{}
}

type GetCompanyProfileRequestBody struct {
	XMLName    xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmscom_GetCompanyProfile"`
	SessionID  string
	CompanyRef string
}

func (r *GetCompanyProfileRequest) RequestBody() *GetCompanyProfileRequestBody {
	return &r.requestBody
}

func (r *GetCompanyProfileRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetCompanyProfileRequest) SetRequestBody(body GetCompanyProfileRequestBody) {
	r.requestBody = body
}

func (r *GetCompanyProfileRequest) NewResponseBody() *GetCompanyProfileResponseBody {
	return &GetCompanyProfileResponseBody{}
}

type GetCompanyProfileResponseBody struct {
	XMLName                       xml.Name          `xml:"pmscom_GetCompanyProfileResponse"`
	PmscomGetCompanyProfileResult ExceptionBlock    `xml:"pmscom_GetCompanyProfileResult"`
	GetCompanyProfile             GetCompanyProfile `xml:"GetCompanyProfile"`
}

func (rb GetCompanyProfileResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmscomGetCompanyProfileResult
}

func (r *GetCompanyProfileRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetCompanyProfileRequest) Do() (GetCompanyProfileResponseBody, error) {
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
