package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetReservationRequest() GetReservationRequest {
	r := GetReservationRequest{
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

type GetReservationRequest struct {
	client        *Client
	queryParams   *GetReservationQueryParams
	pathParams    *GetReservationPathParams
	method        string
	headers       http.Header
	requestBody   GetReservationRequestBody
	requestHeader GetReservationRequestHeader
}

func (r GetReservationRequest) NewQueryParams() *GetReservationQueryParams {
	return &GetReservationQueryParams{}
}

type GetReservationQueryParams struct {
}

func (p GetReservationQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReservationRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetReservationRequest) NewPathParams() *GetReservationPathParams {
	return &GetReservationPathParams{}
}

type GetReservationPathParams struct {
}

func (p *GetReservationPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetReservationRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationRequest) Method() string {
	return r.method
}

func (r GetReservationRequest) NewRequestHeader() GetReservationRequestHeader {
	return GetReservationRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetReservationRequest) RequestHeader() *GetReservationRequestHeader {
	return &r.requestHeader
}

func (r *GetReservationRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetReservationRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetReservationRequest) NewRequestBody() GetReservationRequestBody {
	return GetReservationRequestBody{}
}

type GetReservationRequestBody struct {
	XMLName    xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsbkg_GetReservation"`
	SessionID  string
	BookRef    string `xml:"bookRef,omitempty"`
	RoomPickID int    `xml:"roomPickId,omitempty"`
}

func (r *GetReservationRequest) RequestBody() *GetReservationRequestBody {
	return &r.requestBody
}

func (r *GetReservationRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetReservationRequest) SetRequestBody(body GetReservationRequestBody) {
	r.requestBody = body
}

func (r *GetReservationRequest) NewResponseBody() *GetReservationResponseBody {
	return &GetReservationResponseBody{}
}

type GetReservationResponseBody struct {
	XMLName xml.Name                   `xml:"pmsbkg_GetReservationResponse"`
	Result  PmsbkgGetReservationResult `xml:"pmsbkg_GetReservationResult"`
}

func (rb GetReservationResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.Result.ExceptionBlock
}

func (r *GetReservationRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetReservationRequest) Do() (GetReservationResponseBody, error) {
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

type PmsbkgGetReservationResult struct {
	ExceptionBlock
	Reservations Reservations `xml:"Reservation"`
}

type Reservations []Reservation

type Reservation struct {
	RoomId             string `xml:"RoomId"`
	BookRef            string `xml:"BookRef"`
	RoomPickId         int    `xml:"RoomPickId"`
	BookingType        string `xml:"BookingType"`
	BookingStatus      string `xml:"BookingStatus"`
	RoomTypeCode       string `xml:"RoomTypeCode"`
	Arrival            Date   `xml:"Arrival"`
	Departure          Date   `xml:"Departure"`
	Creation           Date   `xml:"Creation"`
	GroupAllotmentRef  string `xml:"GroupAllotmentRef"`
	MultiRoomReference string `xml:"MultiRoomReference"`
	Contact            struct {
		Name       string `xml:"Name"`
		Salutation string `xml:"Salutation"`
		Forename   string `xml:"Forename"`
		Surname    string `xml:"Surname"`
		ProfileRef string `xml:"ProfileRef"`
		LoyalityId string `xml:"LoyalityId"`
	} `xml:"Contact"`
	Company               string  `xml:"Company"`
	Agent                 string  `xml:"Agent"`
	Source                string  `xml:"Source"`
	Ledger                string  `xml:"Ledger"`
	PostsAllowed          string  `xml:"PostsAllowed"`
	BookRefRoomPickId     string  `xml:"BookRefRoomPickId"`
	AlarmCall             string  `xml:"AlarmCall"`
	MovieAccess           string  `xml:"MovieAccess"`
	AllowRoomMoves        bool    `xml:"AllowRoomMoves"`
	MarketSegment         string  `xml:"MarketSegment"`
	TotalCostNett         float64 `xml:"TotalCostNett"`
	TotalCostGross        float64 `xml:"TotalCostGross"`
	LastEdited            Date    `xml:"LastEdited"`
	GDSRef                string  `xml:"GDSRef"`
	CRSRef1               string  `xml:"CRSRef1"`
	CRSRef2               string  `xml:"CRSRef2"`
	SystemSource          string  `xml:"SystemSource"`
	DistributionChannelId int     `xml:"DistributionChannelId"`
	PreCheckIn            bool    `xml:"PreCheckIn"`
	RezlynxCRS            string  `xml:"RezlynxCRS"`
}
