package touchoffice

import (
	"net/http"
	"net/url"
)

func (c *Client) NewClerkListRequest() ClerkListRequest {
	return ClerkListRequest{
		client:      c,
		queryParams: c.NewClerkListQueryParams(),
		pathParams:  c.NewClerkListPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewClerkListRequestBody(),
	}
}

type ClerkListRequest struct {
	client      *Client
	queryParams *ClerkListQueryParams
	pathParams  *ClerkListPathParams
	method      string
	headers     http.Header
	requestBody ClerkListRequestBody
}

func (c *Client) NewClerkListQueryParams() *ClerkListQueryParams {
	return &ClerkListQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type ClerkListQueryParams struct {
	TAK string
}

func (p ClerkListQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ClerkListRequest) QueryParams() *ClerkListQueryParams {
	return r.queryParams
}

func (c *Client) NewClerkListPathParams() *ClerkListPathParams {
	return &ClerkListPathParams{}
}

type ClerkListPathParams struct {
}

func (p *ClerkListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ClerkListRequest) PathParams() *ClerkListPathParams {
	return r.pathParams
}

func (r *ClerkListRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClerkListRequest) Method() string {
	return r.method
}

func (s *Client) NewClerkListRequestBody() ClerkListRequestBody {
	return ClerkListRequestBody{}
}

type ClerkListRequestBody struct {
}

func (r *ClerkListRequest) RequestBody() *ClerkListRequestBody {
	return &r.requestBody
}

func (r *ClerkListRequest) SetRequestBody(body ClerkListRequestBody) {
	r.requestBody = body
}

func (r *ClerkListRequest) NewResponseBody() *ClerkListResponseBody {
	return &ClerkListResponseBody{}
}

type ClerkListResponseBody struct {
}

func (r *ClerkListRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("clerks/clerkList", r.PathParams())
}

func (r *ClerkListRequest) Do() (ClerkListResponseBody, error) {
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
