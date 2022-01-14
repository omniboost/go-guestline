package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewRoomFolioItemsAndBalanceRequest() RoomFolioItemsAndBalanceRequest {
	r := RoomFolioItemsAndBalanceRequest{
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

type RoomFolioItemsAndBalanceRequest struct {
	client        *Client
	queryParams   *RoomFolioItemsAndBalanceQueryParams
	pathParams    *RoomFolioItemsAndBalancePathParams
	method        string
	headers       http.Header
	requestBody   RoomFolioItemsAndBalanceRequestBody
	requestHeader RoomFolioItemsAndBalanceRequestHeader
}

func (r RoomFolioItemsAndBalanceRequest) NewQueryParams() *RoomFolioItemsAndBalanceQueryParams {
	return &RoomFolioItemsAndBalanceQueryParams{}
}

type RoomFolioItemsAndBalanceQueryParams struct {
}

func (p RoomFolioItemsAndBalanceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *RoomFolioItemsAndBalanceRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r RoomFolioItemsAndBalanceRequest) NewPathParams() *RoomFolioItemsAndBalancePathParams {
	return &RoomFolioItemsAndBalancePathParams{}
}

type RoomFolioItemsAndBalancePathParams struct {
}

func (p *RoomFolioItemsAndBalancePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *RoomFolioItemsAndBalanceRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *RoomFolioItemsAndBalanceRequest) SetMethod(method string) {
	r.method = method
}

func (r *RoomFolioItemsAndBalanceRequest) Method() string {
	return r.method
}

func (r RoomFolioItemsAndBalanceRequest) NewRequestHeader() RoomFolioItemsAndBalanceRequestHeader {
	return RoomFolioItemsAndBalanceRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *RoomFolioItemsAndBalanceRequest) RequestHeader() *RoomFolioItemsAndBalanceRequestHeader {
	return &r.requestHeader
}

func (r *RoomFolioItemsAndBalanceRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type RoomFolioItemsAndBalanceRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r RoomFolioItemsAndBalanceRequest) NewRequestBody() RoomFolioItemsAndBalanceRequestBody {
	return RoomFolioItemsAndBalanceRequestBody{}
}

type RoomFolioItemsAndBalanceRequestBody struct {
	XMLName    xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsbkg_RoomFolioItemsAndBalance"`
	SessionID  string
	RoomID     string `xml:"RoomID,omitempty"`
	FolioID    int    `xml:"FolioID,omitempty"`
	ItemFormat int    `xml:"ItemFormat,omitempty"`
}

func (r *RoomFolioItemsAndBalanceRequest) RequestBody() *RoomFolioItemsAndBalanceRequestBody {
	return &r.requestBody
}

func (r *RoomFolioItemsAndBalanceRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *RoomFolioItemsAndBalanceRequest) SetRequestBody(body RoomFolioItemsAndBalanceRequestBody) {
	r.requestBody = body
}

func (r *RoomFolioItemsAndBalanceRequest) NewResponseBody() *RoomFolioItemsAndBalanceResponseBody {
	return &RoomFolioItemsAndBalanceResponseBody{}
}

type RoomFolioItemsAndBalanceResponseBody struct {
	XMLName                              xml.Name                        `xml:RoomFolioItemsAndBalanceResponse`
	PmsbkgRoomFolioItemsAndBalanceResult ExceptionBlock                  `xml:"pmsbkg_RoomFolioItemsAndBalanceResult"`
	SearchResults                        RoomFolioItemsAndBalanceResults `xml:"SearchResults"`
}

func (rb RoomFolioItemsAndBalanceResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsbkgRoomFolioItemsAndBalanceResult
}

func (r *RoomFolioItemsAndBalanceRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *RoomFolioItemsAndBalanceRequest) Do() (RoomFolioItemsAndBalanceResponseBody, error) {
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

type RoomFolioItemsAndBalanceResults struct {
	// Reservations RoomFolioItemsAndBalanceReservations `xml:"Reservations>Reservation"`
	FolioItems FolioItems
}

type FolioItems []FolioItem

type FolioItem struct {
}
