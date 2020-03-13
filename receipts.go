package touchoffice

import (
	"net/http"
	"net/url"
	"time"
)

func (c *Client) NewReceiptsRequest() ReceiptsRequest {
	return ReceiptsRequest{
		client:      c,
		queryParams: c.NewReceiptsQueryParams(),
		pathParams:  c.NewReceiptsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewReceiptsRequestBody(),
	}
}

type ReceiptsRequest struct {
	client      *Client
	queryParams *ReceiptsQueryParams
	pathParams  *ReceiptsPathParams
	method      string
	headers     http.Header
	requestBody ReceiptsRequestBody
}

func (c *Client) NewReceiptsQueryParams() *ReceiptsQueryParams {
	return &ReceiptsQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type ReceiptsQueryParams struct {
	TAK       string
	Site      int  `schema:"site"`
	DateStart Date `schema:"date_start,omitempty"`
	TimeStart Time `schema:"time_start,omitempty"`
	DateEnd   Date `schema:"date_end,omitempty"`
	TimeEnd   Time `schema:"time_end,omitempty"`
	Sale      int  `schema:"sale,omitempty"`
}

func (p *ReceiptsQueryParams) SetStart(t time.Time) {
	p.DateStart = Date{t}
	p.TimeStart = Time{t}
}

func (p *ReceiptsQueryParams) SetEnd(t time.Time) {
	p.DateEnd = Date{t}
	p.TimeEnd = Time{t}
}

func (p ReceiptsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ReceiptsRequest) QueryParams() *ReceiptsQueryParams {
	return r.queryParams
}

func (c *Client) NewReceiptsPathParams() *ReceiptsPathParams {
	return &ReceiptsPathParams{}
}

type ReceiptsPathParams struct {
}

func (p *ReceiptsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ReceiptsRequest) PathParams() *ReceiptsPathParams {
	return r.pathParams
}

func (r *ReceiptsRequest) SetMethod(method string) {
	r.method = method
}

func (r *ReceiptsRequest) Method() string {
	return r.method
}

func (s *Client) NewReceiptsRequestBody() ReceiptsRequestBody {
	return ReceiptsRequestBody{}
}

type ReceiptsRequestBody struct {
}

func (r *ReceiptsRequest) RequestBody() *ReceiptsRequestBody {
	return &r.requestBody
}

func (r *ReceiptsRequest) SetRequestBody(body ReceiptsRequestBody) {
	r.requestBody = body
}

func (r *ReceiptsRequest) NewResponseBody() *ReceiptsResponseBody {
	return &ReceiptsResponseBody{}
}

type ReceiptsResponseBody struct {
	ReceiptsList
}

func (r *ReceiptsRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("receiptsAndBills/receipts", r.PathParams())
}

func (r *ReceiptsRequest) Do() (ReceiptsResponseBody, error) {
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
