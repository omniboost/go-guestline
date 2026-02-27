package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetEventRequest() GetEventRequest {
	r := GetEventRequest{
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

type GetEventRequest struct {
	client        *Client
	queryParams   *GetEventQueryParams
	pathParams    *GetEventPathParams
	method        string
	headers       http.Header
	requestBody   GetEventRequestBody
	requestHeader GetEventRequestHeader
}

func (r GetEventRequest) NewQueryParams() *GetEventQueryParams {
	return &GetEventQueryParams{}
}

type GetEventQueryParams struct {
}

func (p GetEventQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetEventRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetEventRequest) NewPathParams() *GetEventPathParams {
	return &GetEventPathParams{}
}

type GetEventPathParams struct {
}

func (p *GetEventPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetEventRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetEventRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetEventRequest) Method() string {
	return r.method
}

func (r GetEventRequest) NewRequestHeader() GetEventRequestHeader {
	return GetEventRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetEventRequest) RequestHeader() *GetEventRequestHeader {
	return &r.requestHeader
}

func (r *GetEventRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetEventRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetEventRequest) NewRequestBody() GetEventRequestBody {
	return GetEventRequestBody{}
}

type GetEventRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 cnb_GetEvent"`
	SessionID string
	EventRef  string `xml:"eventRef,omitempty"`
}

func (r *GetEventRequest) RequestBody() *GetEventRequestBody {
	return &r.requestBody
}

func (r *GetEventRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetEventRequest) SetRequestBody(body GetEventRequestBody) {
	r.requestBody = body
}

func (r *GetEventRequest) NewResponseBody() *GetEventResponseBody {
	return &GetEventResponseBody{}
}

type GetEventResponseBody struct {
	XMLName xml.Name          `xml:"cnb_GetEventResponse"`
	Result  CnbGetEventResult `xml:"cnb_GetEventResult"`
}

func (rb GetEventResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.Result.ExceptionBlock
}

func (r *GetEventRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetEventRequest) Do() (GetEventResponseBody, error) {
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

type CnbGetEventResult struct {
	ExceptionBlock
	EventRef      string `xml:"EventRef"`
	ProfileRef    string `xml:"ProfileRef"`
	CompanyRef    string `xml:"CompanyRef"`
	AgentRef      string `xml:"AgentRef"`
	StartDate     Date   `xml:"StartDate"`
	EndDate       Date   `xml:"EndDate"`
	EventStatus   int    `xml:"EventStatus"`
	Delegates     int    `xml:"Delegates"`
	EventName     string `xml:"EventName"`
	Coordinator   string `xml:"Coordinator"`
	EventType     string `xml:"EventType"`
	MarketSegment string `xml:"MarketSegment"`
	PORef         string `xml:"PORef"`
	InternalNotes string `xml:"InternalNotes"`
	EventNotes    string `xml:"EventNotes"`
	CustomNotes1  string `xml:"CustomNotes1"`
	CustomNotes2  string `xml:"CustomNotes2"`
	CustomNotes3  string `xml:"CustomNotes3"`
	TokenSource   string `xml:"TokenSource"`
	BillingNotes  string `xml:"BillingNotes"`
}
