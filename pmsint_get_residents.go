package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetResidentsRequest() GetResidentsRequest {
	r := GetResidentsRequest{
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

type GetResidentsRequest struct {
	client        *Client
	queryParams   *GetResidentsQueryParams
	pathParams    *GetResidentsPathParams
	method        string
	headers       http.Header
	requestBody   GetResidentsRequestBody
	requestHeader GetResidentsRequestHeader
}

func (r GetResidentsRequest) NewQueryParams() *GetResidentsQueryParams {
	return &GetResidentsQueryParams{}
}

type GetResidentsQueryParams struct {
}

func (p GetResidentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetResidentsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetResidentsRequest) NewPathParams() *GetResidentsPathParams {
	return &GetResidentsPathParams{}
}

type GetResidentsPathParams struct {
}

func (p *GetResidentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetResidentsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetResidentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetResidentsRequest) Method() string {
	return r.method
}

func (r GetResidentsRequest) NewRequestHeader() GetResidentsRequestHeader {
	return GetResidentsRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetResidentsRequest) RequestHeader() *GetResidentsRequestHeader {
	return &r.requestHeader
}

func (r *GetResidentsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetResidentsRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetResidentsRequest) NewRequestBody() GetResidentsRequestBody {
	return GetResidentsRequestBody{}
}

type GetResidentsRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetResidents"`
	SessionID string
	FromDate  Time `xml:"FromDate,omitempty"`
	ToDate    Time `xml:"ToDate,omitempty"`
}

func (r *GetResidentsRequest) RequestBody() *GetResidentsRequestBody {
	return &r.requestBody
}

func (r *GetResidentsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetResidentsRequest) SetRequestBody(body GetResidentsRequestBody) {
	r.requestBody = body
}

func (r *GetResidentsRequest) NewResponseBody() *GetResidentsResponseBody {
	return &GetResidentsResponseBody{}
}

type GetResidentsResponseBody struct {
	XMLName                  xml.Name       `xml:GetResidentsResponse`
	PmsbkgGetResidentsResult ExceptionBlock `xml:"PmsintGetResidentsResult"`
	Residents                Residents      `xml:"GetResidents>Residents>cpmsint_GetResidents_ResidentItem"`
}

func (rb GetResidentsResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsbkgGetResidentsResult
}

func (r *GetResidentsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetResidentsRequest) Do() (GetResidentsResponseBody, error) {
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

type Residents []Resident

type Resident struct {
	DepositPaid    string `xml:"DepositPaid"`
	DepositDue     string `xml:"DepositDue"`
	Infants        int    `xml:"Infants"`
	Children       int    `xml:"Children"`
	Adults         int    `xml:"Adults"`
	RoomType       string `xml:"RoomType"`
	Package        string `xml:"Package"`
	Company        string `xml:"Company"`
	Notes          string `xml:"Notes"`
	BookRef        string `xml:"BookRef"`
	BookRefRoomRef string `xml:"BookRefRoomRef"`
	ETA            string `xml:"ETA"`
	Arrival        Time   `xml:"Arrival"`
	Departure      Time   `xml:"Departure"`
	Salutation     string `xml:"Salutation"`
	Surname        string `xml:"Surname"`
	Forename       string `xml:"Forename"`
	RoomID         string `xml:"RoomID"`
	MovieAccess    string `xml:"MovieAccess"`
	PrivateNotes   string `xml:"PrivateNotes"`
	PublicNotes    string `xml:"PublicNotes"`
	CustomNotes1   string `xml:"CustomNotes1"`
	CustomNotes2   string `xml:"CustomNotes2"`
	CustomNotes3   string `xml:"CustomNotes3"`
	ExternalNotes  string `xml:"ExternalNotes"`
	Master         string `xml:"Master"`
}
