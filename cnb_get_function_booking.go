package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetFunctionBookingRequest() FunctionBookingRequest {
	r := FunctionBookingRequest{
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

type FunctionBookingRequest struct {
	client        *Client
	queryParams   *FunctionBookingQueryParams
	pathParams    *FunctionBookingPathParams
	method        string
	headers       http.Header
	requestBody   FunctionBookingRequestBody
	requestHeader FunctionBookingRequestHeader
}

func (r FunctionBookingRequest) NewQueryParams() *FunctionBookingQueryParams {
	return &FunctionBookingQueryParams{}
}

type FunctionBookingQueryParams struct {
}

func (p FunctionBookingQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FunctionBookingRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r FunctionBookingRequest) NewPathParams() *FunctionBookingPathParams {
	return &FunctionBookingPathParams{}
}

type FunctionBookingPathParams struct {
}

func (p *FunctionBookingPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FunctionBookingRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *FunctionBookingRequest) SetMethod(method string) {
	r.method = method
}

func (r *FunctionBookingRequest) Method() string {
	return r.method
}

func (r FunctionBookingRequest) NewRequestHeader() FunctionBookingRequestHeader {
	return FunctionBookingRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *FunctionBookingRequest) RequestHeader() *FunctionBookingRequestHeader {
	return &r.requestHeader
}

func (r *FunctionBookingRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type FunctionBookingRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r FunctionBookingRequest) NewRequestBody() FunctionBookingRequestBody {
	return FunctionBookingRequestBody{}
}

type FunctionBookingRequestBody struct {
	XMLName xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 cnb_GetFunctionBooking"`
	BookRef string   `xml:"bookRef,omitempty"`
}

func (r *FunctionBookingRequest) RequestBody() *FunctionBookingRequestBody {
	return &r.requestBody
}

func (r *FunctionBookingRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *FunctionBookingRequest) SetRequestBody(body FunctionBookingRequestBody) {
	r.requestBody = body
}

func (r *FunctionBookingRequest) NewResponseBody() *FunctionBookingResponseBody {
	return &FunctionBookingResponseBody{}
}

type FunctionBookingResponseBody struct {
	XMLName xml.Name                    `xml:"cnb_GetFunctionBookingResponse"`
	Result  CnbGetFunctionBookingResult `xml:"cnb_GetFunctionBookingResult"`
}

func (r *FunctionBookingRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *FunctionBookingRequest) Do() (FunctionBookingResponseBody, error) {
	var err error

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

type CnbGetFunctionBookingResult struct {
	ExceptionBlock
	FunctionRef    string `xml:"FunctionRef"`
	EventRef       string `xml:"EventRef"`
	FunctionType   string `xml:"FunctionType"`
	FunctionName   string `xml:"FunctionName"`
	FunctionStatus string `xml:"FunctionStatus"`
	StartDate      Date   `xml:"StartDate"`
	EndDate        Date   `xml:"EndDate"`
	Delegates      int    `xml:"Delegates"`
	FunctionRoom   string `xml:"FunctionRoom"`
	FunctionLayout string `xml:"FunctionLayout"`
}
