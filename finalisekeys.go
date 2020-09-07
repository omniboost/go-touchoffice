package touchoffice

import (
	"net/http"
	"net/url"
	"time"
)

func (c *Client) NewFinalisekeysRequest() FinalisekeysRequest {
	return FinalisekeysRequest{
		client:      c,
		queryParams: c.NewFinalisekeysQueryParams(),
		pathParams:  c.NewFinalisekeysPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFinalisekeysRequestBody(),
	}
}

type FinalisekeysRequest struct {
	client      *Client
	queryParams *FinalisekeysQueryParams
	pathParams  *FinalisekeysPathParams
	method      string
	headers     http.Header
	requestBody FinalisekeysRequestBody
}

func (c *Client) NewFinalisekeysQueryParams() *FinalisekeysQueryParams {
	return &FinalisekeysQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type FinalisekeysQueryParams struct {
	TAK       string
	Site      int  `schema:"site"`
	DateStart Date `schema:"date_start,omitempty"`
	TimeStart Time `schema:"time_start,omitempty"`
	DateEnd   Date `schema:"date_end,omitempty"`
	TimeEnd   Time `schema:"time_end,omitempty"`
	Sale      int  `schema:"sale,omitempty"`
}

func (p *FinalisekeysQueryParams) SetStart(t time.Time) {
	p.DateStart = Date{t}
	p.TimeStart = Time{t}
}

func (p *FinalisekeysQueryParams) SetEnd(t time.Time) {
	p.DateEnd = Date{t}
	p.TimeEnd = Time{t}
}

func (p FinalisekeysQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinalisekeysRequest) QueryParams() *FinalisekeysQueryParams {
	return r.queryParams
}

func (c *Client) NewFinalisekeysPathParams() *FinalisekeysPathParams {
	return &FinalisekeysPathParams{}
}

type FinalisekeysPathParams struct {
}

func (p *FinalisekeysPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinalisekeysRequest) PathParams() *FinalisekeysPathParams {
	return r.pathParams
}

func (r *FinalisekeysRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinalisekeysRequest) Method() string {
	return r.method
}

func (s *Client) NewFinalisekeysRequestBody() FinalisekeysRequestBody {
	return FinalisekeysRequestBody{}
}

type FinalisekeysRequestBody struct {
}

func (r *FinalisekeysRequest) RequestBody() *FinalisekeysRequestBody {
	return &r.requestBody
}

func (r *FinalisekeysRequest) SetRequestBody(body FinalisekeysRequestBody) {
	r.requestBody = body
}

func (r *FinalisekeysRequest) NewResponseBody() *FinalisekeysResponseBody {
	return &FinalisekeysResponseBody{}
}

type FinalisekeysResponseBody struct {
	FinalisekeysList
}

func (r *FinalisekeysRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("sales/finaliseKeys", r.PathParams())
}

func (r *FinalisekeysRequest) Do() (FinalisekeysResponseBody, error) {
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
