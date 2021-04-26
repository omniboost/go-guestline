package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetProfileEventsRequest() GetProfileEventsRequest {
	r := GetProfileEventsRequest{
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

type GetProfileEventsRequest struct {
	client        *Client
	queryParams   *GetProfileEventsQueryParams
	pathParams    *GetProfileEventsPathParams
	method        string
	headers       http.Header
	requestBody   GetProfileEventsRequestBody
	requestHeader GetProfileEventsRequestHeader
}

func (r GetProfileEventsRequest) NewQueryParams() *GetProfileEventsQueryParams {
	return &GetProfileEventsQueryParams{}
}

type GetProfileEventsQueryParams struct {
}

func (p GetProfileEventsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetProfileEventsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetProfileEventsRequest) NewPathParams() *GetProfileEventsPathParams {
	return &GetProfileEventsPathParams{}
}

type GetProfileEventsPathParams struct {
}

func (p *GetProfileEventsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetProfileEventsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetProfileEventsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetProfileEventsRequest) Method() string {
	return r.method
}

func (r GetProfileEventsRequest) NewRequestHeader() GetProfileEventsRequestHeader {
	return GetProfileEventsRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetProfileEventsRequest) RequestHeader() *GetProfileEventsRequestHeader {
	return &r.requestHeader
}

func (r *GetProfileEventsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetProfileEventsRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetProfileEventsRequest) NewRequestBody() GetProfileEventsRequestBody {
	return GetProfileEventsRequestBody{}
}

type GetProfileEventsRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsprf_GetProfileEvents"`
	SessionID string   `xml:"SessionId"`
	TransID   int      `xml:"TransId"`
}

func (r *GetProfileEventsRequest) RequestBody() *GetProfileEventsRequestBody {
	return &r.requestBody
}

func (r *GetProfileEventsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetProfileEventsRequest) SetRequestBody(body GetProfileEventsRequestBody) {
	r.requestBody = body
}

func (r *GetProfileEventsRequest) NewResponseBody() *GetProfileEventsResponseBody {
	return &GetProfileEventsResponseBody{}
}

type GetProfileEventsResponseBody struct {
	XMLName                      xml.Name       `xml:GetProfileEventsResponse`
	PmsprfGetProfileEventsResult ExceptionBlock `xml:"pmsprf_GetProfileEventsResult"`
	Events                       ProfileEvents  `xml:"EventResults>Events>GuestProfileEventItem"`
}

func (rb GetProfileEventsResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsprfGetProfileEventsResult
}

func (r *GetProfileEventsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetProfileEventsRequest) Do() (GetProfileEventsResponseBody, error) {
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

type ProfileEvents []ProfileEvent

type ProfileEvent struct {
	TransID         string    `xml:"TransID"`
	EventType       string    `xml:"EventType"`
	EventTimeStamp  time.Time `xml:"EventTimeStamp"`
	AllocatedRoomID string    `xml:"AllocatedRoomID"`
	GuestProfileRef string    `xml:"GuestProfileRef"`
	FolioID         int       `xml:"FolioID"`
	Forename        string    `xml:"Forename"`
	Surname         string    `xml:"Surname"`
	Title           string    `xml:"Title"`
	GuestStatus     string    `xml:"GuestStatus"`
	ArrivalDate     string    `xml:"ArrivalDate"`
	DepartureDate   string    `xml:"DepartureDate"`
	ReservationType string    `xml:"ReservationType"`
	FolioIndex      string    `xml:"FolioIndex"`
	BookRefRoomRef  string    `xml:"BookRefRoomRef"`
	PMSPackageCode  string    `xml:"PMSPackageCode"`
	SPAPackageCode  string    `xml:"SPAPackageCode"`
}
