package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetRoomsRequest() GetRoomsRequest {
	r := GetRoomsRequest{
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

type GetRoomsRequest struct {
	client        *Client
	queryParams   *GetRoomsQueryParams
	pathParams    *GetRoomsPathParams
	method        string
	headers       http.Header
	requestBody   GetRoomsRequestBody
	requestHeader GetRoomsRequestHeader
}

func (r GetRoomsRequest) NewQueryParams() *GetRoomsQueryParams {
	return &GetRoomsQueryParams{}
}

type GetRoomsQueryParams struct {
}

func (p GetRoomsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetRoomsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetRoomsRequest) NewPathParams() *GetRoomsPathParams {
	return &GetRoomsPathParams{}
}

type GetRoomsPathParams struct {
}

func (p *GetRoomsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetRoomsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetRoomsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetRoomsRequest) Method() string {
	return r.method
}

func (r GetRoomsRequest) NewRequestHeader() GetRoomsRequestHeader {
	return GetRoomsRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetRoomsRequest) RequestHeader() *GetRoomsRequestHeader {
	return &r.requestHeader
}

func (r *GetRoomsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetRoomsRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetRoomsRequest) NewRequestBody() GetRoomsRequestBody {
	return GetRoomsRequestBody{}
}

type GetRoomsRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetRooms"`
	SessionID string
	RoomDate  DateTime
}

func (r *GetRoomsRequest) RequestBody() *GetRoomsRequestBody {
	return &r.requestBody
}

func (r *GetRoomsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetRoomsRequest) SetRequestBody(body GetRoomsRequestBody) {
	r.requestBody = body
}

func (r *GetRoomsRequest) NewResponseBody() *GetRoomsResponseBody {
	return &GetRoomsResponseBody{}
}

type GetRoomsResponseBody struct {
	XMLName              xml.Name       `xml:"pmsint_GetRoomsResponse"`
	PmsintGetRoomsResult ExceptionBlock `xml:"pmsint_GetRoomsResult"`
	Rooms                Rooms          `xml:"GetRooms>Rooms>Room"`
}

func (rb GetRoomsResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetRoomsResult
}

func (r *GetRoomsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetRoomsRequest) Do() (GetRoomsResponseBody, error) {
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

type Rooms []Room

type Room struct {
	RoomID string `xml:"RoomID"`
}
