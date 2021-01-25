package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewGetProfileSummaryV3Request() GetProfileSummaryV3Request {
	r := GetProfileSummaryV3Request{
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

type GetProfileSummaryV3Request struct {
	client        *Client
	queryParams   *GetProfileSummaryV3QueryParams
	pathParams    *GetProfileSummaryV3PathParams
	method        string
	headers       http.Header
	requestBody   GetProfileSummaryV3RequestBody
	requestHeader GetProfileSummaryV3RequestHeader
}

func (r GetProfileSummaryV3Request) NewQueryParams() *GetProfileSummaryV3QueryParams {
	return &GetProfileSummaryV3QueryParams{}
}

type GetProfileSummaryV3QueryParams struct {
}

func (p GetProfileSummaryV3QueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetProfileSummaryV3Request) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetProfileSummaryV3Request) NewPathParams() *GetProfileSummaryV3PathParams {
	return &GetProfileSummaryV3PathParams{}
}

type GetProfileSummaryV3PathParams struct {
}

func (p *GetProfileSummaryV3PathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetProfileSummaryV3Request) PathParams() PathParams {
	return r.pathParams
}

func (r *GetProfileSummaryV3Request) SetMethod(method string) {
	r.method = method
}

func (r *GetProfileSummaryV3Request) Method() string {
	return r.method
}

func (r GetProfileSummaryV3Request) NewRequestHeader() GetProfileSummaryV3RequestHeader {
	return GetProfileSummaryV3RequestHeader{
		AuthenticationContext: AuthenticationContext{
			SiteID:       r.client.SiteID(),
			InterfaceID:  r.client.InterfaceID(),
			OperatorCode: r.client.OperatorCode(),
			Password:     r.client.Password(),
		},
	}
}

func (r *GetProfileSummaryV3Request) RequestHeader() *GetProfileSummaryV3RequestHeader {
	return &r.requestHeader
}

func (r *GetProfileSummaryV3Request) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetProfileSummaryV3RequestHeader struct {
	AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetProfileSummaryV3Request) NewRequestBody() GetProfileSummaryV3RequestBody {
	return GetProfileSummaryV3RequestBody{
		ProfileRequestor: ProfileRequestor{
			AuthenticationMethod: "PD",
		},
	}
}

type GetProfileSummaryV3RequestBody struct {
	XMLName          xml.Name         `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsprf_GetProfileSummaryV3"`
	ProfileRequestor ProfileRequestor `xml:"profileRequestor,omitempty"`
}

func (r *GetProfileSummaryV3Request) RequestBody() *GetProfileSummaryV3RequestBody {
	return &r.requestBody
}

func (r *GetProfileSummaryV3Request) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetProfileSummaryV3Request) SetRequestBody(body GetProfileSummaryV3RequestBody) {
	r.requestBody = body
}

func (r *GetProfileSummaryV3Request) NewResponseBody() *GetProfileSummaryV3ResponseBody {
	return &GetProfileSummaryV3ResponseBody{}
}

type GetProfileSummaryV3ResponseBody struct {
	XMLName xml.Name `xml:GetProfileSummaryV3Response`
	// PmsintGetProfileSummaryV3Result ExceptionBlock `xml:"pmsprf_GetProfileSummaryV3Result"`
	ProfileSummary ProfileSummary `xml:"pmsprf_GetProfileSummaryV3Result>ProfileSummary"`
}

func (r *GetProfileSummaryV3Request) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetProfileSummaryV3Request) Do() (GetProfileSummaryV3ResponseBody, error) {
	var err error

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

type AuthenticationContext struct {
	XMLName xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 AuthenticationContext"`

	SiteID       string `xml:"SiteId"`
	InterfaceID  string `xml:"InterfaceId"`
	OperatorCode string `xml:"OperatorCode"`
	Password     string `xml:"Password"`
}

type ProfileRequestor struct {
	// Unique profile ID, must be either a profile reference, unique email address or loyalty ID
	ProfileUniqueID string `xml:"ProfileUniqueId,omitempty"`
	// The password that is associated with this profile; if you know the
	// password then you donâ€™t need to send any of the other values below,
	// otherwise leave this blank and send through all the second factor
	// authentication values below
	ProfileUniqueIdAuthenticatorValue string `xml:"ProfileUniqueIdAuthenticatorValue,omitemtpy"`
	// "PD" stands for "Person Details" and should always be set to this
	AuthenticationMethod string `xml:"AuthenticationMethod"`
	// The method of second line authentication to be used i.e. "Forename", "Surname" or "PostCode"
	AuthenticationCode string `xml:"AuthenticationCode"`
	// The authentication value, e.g. if the method is "Surname" and the profile
	// is under the name of "Jones" then the value will be "Jones", if the
	// method is "PostCode" then the value will be "SY26LG"
	AuthenticationValue string `xml:"AuthenticationValue"`
	// Optional filter to filter results according to the reservation status
	// "Checked Out", "Resident" or "Future"
	Filters string `xml:"Filters,omitempty"`
	// Used to indicate whether searching for a reservation contact, reservation guest or both
	// "Contact", "Guest" or "Both"
	ProfileSearchMode string `xml:"ProfileSearchMode,omitempty"`
}

type ProfileSummary struct {
	ProfileRef                    string `xml:"ProfileRef"`
	Salutation                    string `xml:"Salutation"`
	Forename                      string `xml:"Forename"`
	Surname                       string `xml:"Surname"`
	EMailAddress                  string `xml:"EMailAddress"`
	Gender                        string `xml:"Gender"`
	Street                        string `xml:"Street"`
	Area                          string `xml:"Area"`
	Town                          string `xml:"Town"`
	County                        string `xml:"County"`
	PostCode                      string `xml:"PostCode"`
	Country                       string `xml:"Country"`
	TelephoneNo                   string `xml:"TelephoneNo"`
	FaxNo                         string `xml:"FaxNo"`
	MobileNo                      string `xml:"MobileNo"`
	Nationality                   string `xml:"Nationality"`
	Language                      string `xml:"Language"`
	CompanyRef                    string `xml:"CompanyRef"`
	Vip                           string `xml:"Vip"`
	DateOfBirth                   string `xml:"DateOfBirth"`
	StudentId                     string `xml:"StudentId"`
	TypeOfPerson                  string `xml:"TypeOfPerson"`
	PassportNumber                string `xml:"PassportNumber"`
	PreferredRoomType             string `xml:"PreferredRoomType"`
	ExcludeFromMailings           string `xml:"ExcludeFromMailings"`
	Smoker                        string `xml:"Smoker"`
	BlackListed                   string `xml:"BlackListed"`
	LoyaltyId                     string `xml:"LoyaltyId"`
	ProfileType                   string `xml:"ProfileType"`
	ExcludeFromThirdPartyMailings string `xml:"ExcludeFromThirdPartyMailings"`
	Suffix                        string `xml:"Suffix"`
	MiddleName                    string `xml:"MiddleName"`
	AddressLine1                  string `xml:"AddressLine1"`
	AddressLine2                  string `xml:"AddressLine2"`
	State                         string `xml:"State"`
	PublicNotes                   string `xml:"PublicNotes"`
	PrivateNotes                  string `xml:"PrivateNotes"`
	CustomNotes1                  string `xml:"CustomNotes1"`
	CustomNotes2                  string `xml:"CustomNotes2"`
	CustomNotes3                  string `xml:"CustomNotes3"`
}
