package touchoffice

import (
	"net/http"
	"net/url"
)

func (c *Client) NewDepartmentsListRequest() DepartmentsListRequest {
	return DepartmentsListRequest{
		client:      c,
		queryParams: c.NewDepartmentsListQueryParams(),
		pathParams:  c.NewDepartmentsListPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewDepartmentsListRequestBody(),
	}
}

type DepartmentsListRequest struct {
	client      *Client
	queryParams *DepartmentsListQueryParams
	pathParams  *DepartmentsListPathParams
	method      string
	headers     http.Header
	requestBody DepartmentsListRequestBody
}

func (c *Client) NewDepartmentsListQueryParams() *DepartmentsListQueryParams {
	return &DepartmentsListQueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type DepartmentsListQueryParams struct {
	TAK string
}

func (p DepartmentsListQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DepartmentsListRequest) QueryParams() *DepartmentsListQueryParams {
	return r.queryParams
}

func (c *Client) NewDepartmentsListPathParams() *DepartmentsListPathParams {
	return &DepartmentsListPathParams{}
}

type DepartmentsListPathParams struct {
}

func (p *DepartmentsListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DepartmentsListRequest) PathParams() *DepartmentsListPathParams {
	return r.pathParams
}

func (r *DepartmentsListRequest) SetMethod(method string) {
	r.method = method
}

func (r *DepartmentsListRequest) Method() string {
	return r.method
}

func (s *Client) NewDepartmentsListRequestBody() DepartmentsListRequestBody {
	return DepartmentsListRequestBody{}
}

type DepartmentsListRequestBody struct {
}

func (r *DepartmentsListRequest) RequestBody() *DepartmentsListRequestBody {
	return &r.requestBody
}

func (r *DepartmentsListRequest) SetRequestBody(body DepartmentsListRequestBody) {
	r.requestBody = body
}

func (r *DepartmentsListRequest) NewResponseBody() *DepartmentsListResponseBody {
	return &DepartmentsListResponseBody{}
}

type DepartmentsListResponseBody struct {
	DepartmentsList
}

func (r *DepartmentsListRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("department/departmentsList", r.PathParams())
}

func (r *DepartmentsListRequest) Do() (DepartmentsListResponseBody, error) {
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
