package touchoffice

import (
	"net/http"
	"net/url"
)

func (c *Client) NewSitesListRequest() SitesListRequest {
	return SitesListRequest{
		client:      c,
		queryParams: c.NewSitesListQueryParams(),
		pathParams:  c.NewSitesListPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewSitesListRequestBody(),
	}
}

type SitesListRequest struct {
	client      *Client
	queryParams *SitesListQueryParams
	pathParams  *SitesListPathParams
	method      string
	headers     http.Header
	requestBody SitesListRequestBody
}

func (c *Client) NewSitesListQueryParams() *SitesListQueryParams {
	return &SitesListQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type SitesListQueryParams struct {
	TAK string
}

func (p SitesListQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SitesListRequest) QueryParams() *SitesListQueryParams {
	return r.queryParams
}

func (c *Client) NewSitesListPathParams() *SitesListPathParams {
	return &SitesListPathParams{}
}

type SitesListPathParams struct {
}

func (p *SitesListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SitesListRequest) PathParams() *SitesListPathParams {
	return r.pathParams
}

func (r *SitesListRequest) SetMethod(method string) {
	r.method = method
}

func (r *SitesListRequest) Method() string {
	return r.method
}

func (s *Client) NewSitesListRequestBody() SitesListRequestBody {
	return SitesListRequestBody{}
}

type SitesListRequestBody struct {
}

func (r *SitesListRequest) RequestBody() *SitesListRequestBody {
	return &r.requestBody
}

func (r *SitesListRequest) SetRequestBody(body SitesListRequestBody) {
	r.requestBody = body
}

func (r *SitesListRequest) NewResponseBody() *SitesListResponseBody {
	return &SitesListResponseBody{}
}

type SitesListResponseBody struct {
	SitesList
}

func (r *SitesListRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("sites/sitesList", r.PathParams())
}

func (r *SitesListRequest) Do() (SitesListResponseBody, error) {
	u, err := r.URL()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), u, nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
