package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetReservationBookingLinesRequest() GetReservationBookingLinesRequest {
	r := GetReservationBookingLinesRequest{
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

type GetReservationBookingLinesRequest struct {
	client        *Client
	queryParams   *GetReservationBookingLinesQueryParams
	pathParams    *GetReservationBookingLinesPathParams
	method        string
	headers       http.Header
	requestBody   GetReservationBookingLinesRequestBody
	requestHeader GetReservationBookingLinesRequestHeader
}

func (r GetReservationBookingLinesRequest) NewQueryParams() *GetReservationBookingLinesQueryParams {
	return &GetReservationBookingLinesQueryParams{}
}

type GetReservationBookingLinesQueryParams struct {
}

func (p GetReservationBookingLinesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReservationBookingLinesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetReservationBookingLinesRequest) NewPathParams() *GetReservationBookingLinesPathParams {
	return &GetReservationBookingLinesPathParams{}
}

type GetReservationBookingLinesPathParams struct {
}

func (p *GetReservationBookingLinesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationBookingLinesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetReservationBookingLinesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationBookingLinesRequest) Method() string {
	return r.method
}

func (r GetReservationBookingLinesRequest) NewRequestHeader() GetReservationBookingLinesRequestHeader {
	return GetReservationBookingLinesRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetReservationBookingLinesRequest) RequestHeader() *GetReservationBookingLinesRequestHeader {
	return &r.requestHeader
}

func (r *GetReservationBookingLinesRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetReservationBookingLinesRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetReservationBookingLinesRequest) NewRequestBody() GetReservationBookingLinesRequestBody {
	return GetReservationBookingLinesRequestBody{}
}

type GetReservationBookingLinesRequestBody struct {
	XMLName    xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsbkg_GetReservationBookingLines"`
	SessionID  string
	BookRef    string `xml:"bookRef,omitempty"`
	RoomPickID int    `xml:"roomPickId,omitempty"`
}

func (r *GetReservationBookingLinesRequest) RequestBody() *GetReservationBookingLinesRequestBody {
	return &r.requestBody
}

func (r *GetReservationBookingLinesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetReservationBookingLinesRequest) SetRequestBody(body GetReservationBookingLinesRequestBody) {
	r.requestBody = body
}

func (r *GetReservationBookingLinesRequest) NewResponseBody() *GetReservationBookingLinesResponseBody {
	return &GetReservationBookingLinesResponseBody{}
}

type GetReservationBookingLinesResponseBody struct {
	XMLName xml.Name                               `xml:"pmsbkg_GetReservationBookingLinesResponse"`
	Result  PmsbkgGetReservationBookingLinesResult `xml:"pmsbkg_GetReservationBookingLinesResult"`
}

func (rb GetReservationBookingLinesResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.Result.ExceptionBlock
}

func (r *GetReservationBookingLinesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetReservationBookingLinesRequest) Do() (GetReservationBookingLinesResponseBody, error) {
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

type PmsbkgGetReservationBookingLinesResult struct {
	ExceptionBlock
	BookingLines BookingLines `xml:"BookingLines>BookingLine"`
}

type BookingLines []BookingLine

type BookingLine struct {
	ID            string  `xml:"Id"`
	Date          Time    `xml:"Date"`
	Adults        int     `xml:"Adults"`
	Children      int     `xml:"Children"`
	Infants       int     `xml:"Infants"`
	RatePlan      string  `xml:"RatePlan"`
	MarketSegment string  `xml:"MarketSegment"`
	RoomType      string  `xml:"RoomType"`
	RoomId        string  `xml:"RoomId"`
	Nett          float64 `xml:"Nett"`
	Gross         float64 `xml:"Gross"`
	DayLet        string  `xml:"DayLet"`
}
