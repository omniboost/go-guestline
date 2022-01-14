package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewRoomTypeListRequest() RoomTypeListRequest {
	r := RoomTypeListRequest{
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

type RoomTypeListRequest struct {
	client        *Client
	queryParams   *RoomTypeListQueryParams
	pathParams    *RoomTypeListPathParams
	method        string
	headers       http.Header
	requestBody   RoomTypeListRequestBody
	requestHeader RoomTypeListRequestHeader
}

func (r RoomTypeListRequest) NewQueryParams() *RoomTypeListQueryParams {
	return &RoomTypeListQueryParams{}
}

type RoomTypeListQueryParams struct {
}

func (p RoomTypeListQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *RoomTypeListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r RoomTypeListRequest) NewPathParams() *RoomTypeListPathParams {
	return &RoomTypeListPathParams{}
}

type RoomTypeListPathParams struct {
}

func (p *RoomTypeListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *RoomTypeListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *RoomTypeListRequest) SetMethod(method string) {
	r.method = method
}

func (r *RoomTypeListRequest) Method() string {
	return r.method
}

func (r RoomTypeListRequest) NewRequestHeader() RoomTypeListRequestHeader {
	return RoomTypeListRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *RoomTypeListRequest) RequestHeader() *RoomTypeListRequestHeader {
	return &r.requestHeader
}

func (r *RoomTypeListRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type RoomTypeListRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r RoomTypeListRequest) NewRequestBody() RoomTypeListRequestBody {
	return RoomTypeListRequestBody{}
}

type RoomTypeListRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmscfg_RoomTypeList"`
	SessionID string
	RoomDate  DateTime
}

func (r *RoomTypeListRequest) RequestBody() *RoomTypeListRequestBody {
	return &r.requestBody
}

func (r *RoomTypeListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *RoomTypeListRequest) SetRequestBody(body RoomTypeListRequestBody) {
	r.requestBody = body
}

func (r *RoomTypeListRequest) NewResponseBody() *RoomTypeListResponseBody {
	return &RoomTypeListResponseBody{}
}

type RoomTypeListResponseBody struct {
	XMLName                  xml.Name       `xml:"pmsint_RoomTypeListResponse"`
	PmsintRoomTypeListResult ExceptionBlock `xml:"pmsint_RoomTypeListResult"`
	RoomTypes                RoomTypes      `xml:"RoomTypeList>RoomTypeList>Room"`
}

func (rb RoomTypeListResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintRoomTypeListResult
}

func (r *RoomTypeListRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *RoomTypeListRequest) Do() (RoomTypeListResponseBody, error) {
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

type RoomTypes []RoomType

type RoomType struct {
}
