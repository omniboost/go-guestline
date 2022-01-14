package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetRoomStatusRequest() GetRoomStatusRequest {
	r := GetRoomStatusRequest{
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

type GetRoomStatusRequest struct {
	client        *Client
	queryParams   *GetRoomStatusQueryParams
	pathParams    *GetRoomStatusPathParams
	method        string
	headers       http.Header
	requestBody   GetRoomStatusRequestBody
	requestHeader GetRoomStatusRequestHeader
}

func (r GetRoomStatusRequest) NewQueryParams() *GetRoomStatusQueryParams {
	return &GetRoomStatusQueryParams{}
}

type GetRoomStatusQueryParams struct {
}

func (p GetRoomStatusQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetRoomStatusRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetRoomStatusRequest) NewPathParams() *GetRoomStatusPathParams {
	return &GetRoomStatusPathParams{}
}

type GetRoomStatusPathParams struct {
}

func (p *GetRoomStatusPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetRoomStatusRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetRoomStatusRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetRoomStatusRequest) Method() string {
	return r.method
}

func (r GetRoomStatusRequest) NewRequestHeader() GetRoomStatusRequestHeader {
	return GetRoomStatusRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetRoomStatusRequest) RequestHeader() *GetRoomStatusRequestHeader {
	return &r.requestHeader
}

func (r *GetRoomStatusRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetRoomStatusRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetRoomStatusRequest) NewRequestBody() GetRoomStatusRequestBody {
	return GetRoomStatusRequestBody{}
}

type GetRoomStatusRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetRoomStatus"`
	SessionID string
	RoomID    string
}

func (r *GetRoomStatusRequest) RequestBody() *GetRoomStatusRequestBody {
	return &r.requestBody
}

func (r *GetRoomStatusRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetRoomStatusRequest) SetRequestBody(body GetRoomStatusRequestBody) {
	r.requestBody = body
}

func (r *GetRoomStatusRequest) NewResponseBody() *GetRoomStatusResponseBody {
	return &GetRoomStatusResponseBody{}
}

type GetRoomStatusResponseBody struct {
	XMLName                   xml.Name       `xml:GetRoomStatusResponse`
	PmsintGetRoomStatusResult ExceptionBlock `xml:"pmsint_GetRoomStatusResult"`
	RoomStatus                RoomStatus     `xml:"GetRoomStatus>Departures>cpmsint_GetRoomStatus_DepartureItem"`
}

func (rb GetRoomStatusResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetRoomStatusResult
}

func (r *GetRoomStatusRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetRoomStatusRequest) Do() (GetRoomStatusResponseBody, error) {
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

type RoomStatus struct {
}
