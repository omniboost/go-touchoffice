package touchoffice

import (
	"net/http"
	"net/url"
)

func (c *Client) NewPLUGroupsListRequest() PLUGroupsListRequest {
	return PLUGroupsListRequest{
		client:      c,
		queryParams: c.NewPLUGroupsListQueryParams(),
		pathParams:  c.NewPLUGroupsListPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewPLUGroupsListRequestBody(),
	}
}

type PLUGroupsListRequest struct {
	client      *Client
	queryParams *PLUGroupsListQueryParams
	pathParams  *PLUGroupsListPathParams
	method      string
	headers     http.Header
	requestBody PLUGroupsListRequestBody
}

func (c *Client) NewPLUGroupsListQueryParams() *PLUGroupsListQueryParams {
	return &PLUGroupsListQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type PLUGroupsListQueryParams struct {
	TAK string
}

func (p PLUGroupsListQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PLUGroupsListRequest) QueryParams() *PLUGroupsListQueryParams {
	return r.queryParams
}

func (c *Client) NewPLUGroupsListPathParams() *PLUGroupsListPathParams {
	return &PLUGroupsListPathParams{}
}

type PLUGroupsListPathParams struct {
}

func (p *PLUGroupsListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PLUGroupsListRequest) PathParams() *PLUGroupsListPathParams {
	return r.pathParams
}

func (r *PLUGroupsListRequest) SetMethod(method string) {
	r.method = method
}

func (r *PLUGroupsListRequest) Method() string {
	return r.method
}

func (s *Client) NewPLUGroupsListRequestBody() PLUGroupsListRequestBody {
	return PLUGroupsListRequestBody{}
}

type PLUGroupsListRequestBody struct {
}

func (r *PLUGroupsListRequest) RequestBody() *PLUGroupsListRequestBody {
	return &r.requestBody
}

func (r *PLUGroupsListRequest) SetRequestBody(body PLUGroupsListRequestBody) {
	r.requestBody = body
}

func (r *PLUGroupsListRequest) NewResponseBody() *PLUGroupsListResponseBody {
	return &PLUGroupsListResponseBody{}
}

type PLUGroupsListResponseBody struct {
	PLUGroupsList
}

func (r *PLUGroupsListRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("groups/pluGroupsList", r.PathParams())
}

func (r *PLUGroupsListRequest) Do() (PLUGroupsListResponseBody, error) {
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
