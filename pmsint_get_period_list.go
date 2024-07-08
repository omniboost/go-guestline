package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetPeriodListRequest() GetPeriodListRequest {
	r := GetPeriodListRequest{
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

type GetPeriodListRequest struct {
	client        *Client
	queryParams   *GetPeriodListQueryParams
	pathParams    *GetPeriodListPathParams
	method        string
	headers       http.Header
	requestBody   GetPeriodListRequestBody
	requestHeader GetPeriodListRequestHeader
}

func (r GetPeriodListRequest) NewQueryParams() *GetPeriodListQueryParams {
	return &GetPeriodListQueryParams{}
}

type GetPeriodListQueryParams struct {
}

func (p GetPeriodListQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPeriodListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetPeriodListRequest) NewPathParams() *GetPeriodListPathParams {
	return &GetPeriodListPathParams{}
}

type GetPeriodListPathParams struct {
}

func (p *GetPeriodListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPeriodListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetPeriodListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPeriodListRequest) Method() string {
	return r.method
}

func (r GetPeriodListRequest) NewRequestHeader() GetPeriodListRequestHeader {
	return GetPeriodListRequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetPeriodListRequest) RequestHeader() *GetPeriodListRequestHeader {
	return &r.requestHeader
}

func (r *GetPeriodListRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetPeriodListRequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetPeriodListRequest) NewRequestBody() GetPeriodListRequestBody {
	return GetPeriodListRequestBody{}
}

type GetPeriodListRequestBody struct {
	XMLName       xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetPeriodList"`
	SessionID     string
	RoomDate      *DateTime  `xml:"RoomDate,omitempty"`
	IPeriodID     int        `xml:"IPeriodID,omitempty"`
	EnmPeriodType PeriodType `xml:"enmPeriodType,omitempty"`
}

func (r *GetPeriodListRequest) RequestBody() *GetPeriodListRequestBody {
	return &r.requestBody
}

func (r *GetPeriodListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetPeriodListRequest) SetRequestBody(body GetPeriodListRequestBody) {
	r.requestBody = body
}

func (r *GetPeriodListRequest) NewResponseBody() *GetPeriodListResponseBody {
	return &GetPeriodListResponseBody{}
}

type GetPeriodListResponseBody struct {
	XMLName                   xml.Name       `xml:"pmsint_GetPeriodListResponse"`
	PmsintGetPeriodListResult ExceptionBlock `xml:"pmsint_GetPeriodListResult"`
	Period                    Periods        `xml:"GetPeriodList>Periods>cpmsint_GetPeriodListItem"`
	RoomTypes                 RoomTypes      `xml:"GetPeriodList>GetPeriodList>Room"`
}

func (rb GetPeriodListResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetPeriodListResult
}

func (r *GetPeriodListRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetPeriodListRequest) Do() (GetPeriodListResponseBody, error) {
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
