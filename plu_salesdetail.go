package touchoffice

import (
	"net/http"
	"net/url"
	"time"
)

func (c *Client) NewPLUSalesdetailRequest() PLUSalesdetailRequest {
	return PLUSalesdetailRequest{
		client:      c,
		queryParams: c.NewPLUSalesdetailQueryParams(),
		pathParams:  c.NewPLUSalesdetailPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewPLUSalesdetailRequestBody(),
	}
}

type PLUSalesdetailRequest struct {
	client      *Client
	queryParams *PLUSalesdetailQueryParams
	pathParams  *PLUSalesdetailPathParams
	method      string
	headers     http.Header
	requestBody PLUSalesdetailRequestBody
}

func (c *Client) NewPLUSalesdetailQueryParams() *PLUSalesdetailQueryParams {
	return &PLUSalesdetailQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type PLUSalesdetailQueryParams struct {
	TAK       string
	Site      int  `schema:"site"`
	DateStart Date `schema:"date_start,omitempty"`
	TimeStart Time `schema:"time_start,omitempty"`
	DateEnd   Date `schema:"date_end,omitempty"`
	TimeEnd   Time `schema:"time_end,omitempty"`
	Sale      int  `schema:"sale,omitempty"`
}

func (p *PLUSalesdetailQueryParams) SetStart(t time.Time) {
	p.DateStart = Date{t}
	p.TimeStart = Time{t}
}

func (p *PLUSalesdetailQueryParams) SetEnd(t time.Time) {
	p.DateEnd = Date{t}
	p.TimeEnd = Time{t}
}

func (p PLUSalesdetailQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PLUSalesdetailRequest) QueryParams() *PLUSalesdetailQueryParams {
	return r.queryParams
}

func (c *Client) NewPLUSalesdetailPathParams() *PLUSalesdetailPathParams {
	return &PLUSalesdetailPathParams{}
}

type PLUSalesdetailPathParams struct {
}

func (p *PLUSalesdetailPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PLUSalesdetailRequest) PathParams() *PLUSalesdetailPathParams {
	return r.pathParams
}

func (r *PLUSalesdetailRequest) SetMethod(method string) {
	r.method = method
}

func (r *PLUSalesdetailRequest) Method() string {
	return r.method
}

func (s *Client) NewPLUSalesdetailRequestBody() PLUSalesdetailRequestBody {
	return PLUSalesdetailRequestBody{}
}

type PLUSalesdetailRequestBody struct {
}

func (r *PLUSalesdetailRequest) RequestBody() *PLUSalesdetailRequestBody {
	return &r.requestBody
}

func (r *PLUSalesdetailRequest) SetRequestBody(body PLUSalesdetailRequestBody) {
	r.requestBody = body
}

func (r *PLUSalesdetailRequest) NewResponseBody() *PLUSalesdetailResponseBody {
	return &PLUSalesdetailResponseBody{}
}

type PLUSalesdetailResponseBody struct {
	SalesList
}

func (r *PLUSalesdetailRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("sales/PLUsalesdetail", r.PathParams())
}

func (r *PLUSalesdetailRequest) Do() (PLUSalesdetailResponseBody, error) {
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
