package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetReservationProductsRequest() GetReservationProductsRequest {
	r := GetReservationProductsRequest{
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

type GetReservationProductsRequest struct {
	client        *Client
	queryParams   *GetReservationProductsQueryParams
	pathParams    *GetReservationProductsPathParams
	method        string
	headers       http.Header
	requestBody   GetReservationProductsRequestBody
	requestHeader GetReservationProductsRequestHeader
}

func (r GetReservationProductsRequest) NewQueryParams() *GetReservationProductsQueryParams {
	return &GetReservationProductsQueryParams{}
}

type GetReservationProductsQueryParams struct {
}

func (p GetReservationProductsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReservationProductsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetReservationProductsRequest) NewPathParams() *GetReservationProductsPathParams {
	return &GetReservationProductsPathParams{}
}

type GetReservationProductsPathParams struct {
}

func (p *GetReservationProductsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationProductsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetReservationProductsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationProductsRequest) Method() string {
	return r.method
}

func (r GetReservationProductsRequest) NewRequestHeader() GetReservationProductsRequestHeader {
	return GetReservationProductsRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetReservationProductsRequest) RequestHeader() *GetReservationProductsRequestHeader {
	return &r.requestHeader
}

func (r *GetReservationProductsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetReservationProductsRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetReservationProductsRequest) NewRequestBody() GetReservationProductsRequestBody {
	return GetReservationProductsRequestBody{}
}

type GetReservationProductsRequestBody struct {
	XMLName    xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsbkg_GetReservationProducts"`
	SessionID  string
	BookRef    string `xml:"bookRef,omitempty"`
	RoomPickID int    `xml:"roomPickId,omitempty"`
}

func (r *GetReservationProductsRequest) RequestBody() *GetReservationProductsRequestBody {
	return &r.requestBody
}

func (r *GetReservationProductsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetReservationProductsRequest) SetRequestBody(body GetReservationProductsRequestBody) {
	r.requestBody = body
}

func (r *GetReservationProductsRequest) NewResponseBody() *GetReservationProductsResponseBody {
	return &GetReservationProductsResponseBody{}
}

type GetReservationProductsResponseBody struct {
	XMLName                            xml.Name                           `xml:"pmsbkg_GetReservationProductsResponse"`
	PmsbkgGetReservationProductsResult PmsbkgGetReservationProductsResult `xml:"pmsbkg_GetReservationProductsResult"`
}

func (rb GetReservationProductsResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsbkgGetReservationProductsResult.ExceptionBlock
}

func (r *GetReservationProductsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetReservationProductsRequest) Do() (GetReservationProductsResponseBody, error) {
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

type PmsbkgGetReservationProductsResult struct {
	ExceptionBlock
	ReservationProducts ReservationProducts `xml:"ReservationProducts>ReservationProduct"`
}

type ReservationProducts []ReservationProduct

type ReservationProduct struct {
	ID          string `xml:"Id"`
	ProductCode string `xml:"ProductCode"`
	Description string `xml:"Description"`
	Quantity    string `xml:"Quantity"`
	Value       string `xml:"Value"`
	ChargeMode  string `xml:"ChargeMode"`
	PostingMode struct {
		Text   string `xml:",chardata"`
		Custom struct {
			Text         string `xml:",chardata"`
			PostingItems struct {
				Text        string `xml:",chardata"`
				PostingItem []struct {
					Text     string `xml:",chardata"`
					ForDate  string `xml:"ForDate"`
					Quantity string `xml:"Quantity"`
					Value    string `xml:"Value"`
				} `xml:"PostingItem"`
			} `xml:"PostingItems"`
		} `xml:"Custom"`
	} `xml:"PostingMode"`
}
