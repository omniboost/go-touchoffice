package touchoffice

import (
	"net/http"
	"net/url"
)

func (c *Client) NewPLUList2Request() PLUList2Request {
	return PLUList2Request{
		client:      c,
		queryParams: c.NewPLUList2QueryParams(),
		pathParams:  c.NewPLUList2PathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewPLUList2RequestBody(),
	}
}

type PLUList2Request struct {
	client      *Client
	queryParams *PLUList2QueryParams
	pathParams  *PLUList2PathParams
	method      string
	headers     http.Header
	requestBody PLUList2RequestBody
}

func (c *Client) NewPLUList2QueryParams() *PLUList2QueryParams {
	return &PLUList2QueryParams{
		TAK: c.TerminalAccessKey(),
	}
}

type PLUList2QueryParams struct {
	//  the terminal access key
	TAK string
	// the site id (default 0 [head office])
	Site int `schema:"site"`
	//  Filter by department (optional)
	Department int `schema:"department"`
	// Filter by PLU group (optional)
	PLUGroup int `schema:"plugroup"`
	// Filter by Shelf Edge label { pending or all } (optional)
	SEL string `schema:"sel"`
	// Filter by plu { single PLU number or pipe seperated list } (optional)
	PLU string `schema:"plu"`
}

func (p PLUList2QueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PLUList2Request) QueryParams() *PLUList2QueryParams {
	return r.queryParams
}

func (c *Client) NewPLUList2PathParams() *PLUList2PathParams {
	return &PLUList2PathParams{}
}

type PLUList2PathParams struct {
}

func (p *PLUList2PathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PLUList2Request) PathParams() *PLUList2PathParams {
	return r.pathParams
}

func (r *PLUList2Request) SetMethod(method string) {
	r.method = method
}

func (r *PLUList2Request) Method() string {
	return r.method
}

func (s *Client) NewPLUList2RequestBody() PLUList2RequestBody {
	return PLUList2RequestBody{}
}

type PLUList2RequestBody struct {
	RequestedFields []string `json:"requestedFields"`
}

func (r *PLUList2Request) RequestBody() *PLUList2RequestBody {
	return &r.requestBody
}

func (r *PLUList2Request) SetRequestBody(body PLUList2RequestBody) {
	r.requestBody = body
}

func (r *PLUList2Request) NewResponseBody() *PLUList2ResponseBody {
	return &PLUList2ResponseBody{}
}

type PLUList2ResponseBody struct {
	PLUList2
}

func (r *PLUList2Request) URL() (url.URL, error) {
	return r.client.GetEndpointURL("products/PLUList2", r.PathParams())
}

func (r *PLUList2Request) Do() (PLUList2ResponseBody, error) {
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
