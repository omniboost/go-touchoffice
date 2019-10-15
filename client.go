package touchoffice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"sync"
	"text/template"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-touchoffice/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.touchoffice.net",
		Path:   "index.php",
		// RawQuery: "demo=true",
	}
)

// NewClient returns a new InvoiceXpress Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http: httpClient,
	}

	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	apiKey            string
	terminalAccessKey string

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetAPIKey(key string) {
	c.apiKey = key
}

func (c Client) APIKey() string {
	return c.apiKey
}

func (c *Client) SetTerminalAccessKey(key string) {
	c.terminalAccessKey = key
}

func (c Client) TerminalAccessKey() string {
	return c.terminalAccessKey
}

func (c *Client) GetEndpointURL(relative string, pathParams PathParams) (url.URL, error) {
	clientURL := c.BaseURL()
	relativeURL, err := url.Parse(relative)
	if err != nil {
		return clientURL, err
	}

	clientURL.Path = path.Join(clientURL.Path, relativeURL.Path)

	query := url.Values{}
	for k, v := range clientURL.Query() {
		query[k] = append(query[k], v...)
	}
	for k, v := range relativeURL.Query() {
		query[k] = append(query[k], v...)
	}
	clientURL.RawQuery = query.Encode()

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		return clientURL, err
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()

	for k, v := range params {
		params[k] = QueryEscape(v)
	}

	err = tmpl.Execute(buf, params)
	if err != nil {
		return clientURL, err
	}
	clientURL.Path = buf.String()

	return clientURL, nil
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())
	req.Header.Add("X-API-KEY", c.APIKey())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	errorResponse := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, &responseBody, &errorResponse)
	if err != nil {
		return httpResp, err
	}

	if len(errorResponse.Message.Errors) > 0 {
		return httpResp, errorResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(len(vv))
	errs := []error{}
	writers := make([]io.Writer, len(vv))

	for i, v := range vv {
		pr, pw := io.Pipe()
		writers[i] = pw

		go func(i int, v interface{}, pr *io.PipeReader, pw *io.PipeWriter) {
			dec := json.NewDecoder(pr)
			if c.disallowUnknownFields {
				dec.DisallowUnknownFields()
			}

			err := dec.Decode(v)
			if err != nil {
				errs = append(errs, err)
			}

			// mark routine as done
			wg.Done()

			// Drain reader
			io.Copy(ioutil.Discard, pr)

			// close reader
			// pr.CloseWithError(err)
			pr.Close()
		}(i, v, pr, pw)
	}

	// copy the data in a multiwriter
	mw := io.MultiWriter(writers...)
	_, err := io.Copy(mw, r)
	if err != nil {
		return err
	}

	wg.Wait()
	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}
	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; (c >= 200 && c <= 299) || c == 400 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Message Message
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Message)
}

func (r *ErrorResponse) Error() string {
	return r.Message.Error()
}

type Message struct {
	Message        string  `json:"message"`
	ErrorCode      string  `json:"errorCode"`
	DeveloperHint  string  `json:"developerHint"`
	LogID          string  `json:"logId"`
	HTTPStatusCode int     `json:"httpStatusCode"`
	Errors         []Error `json:"errors"`
	LogTime        LogTime `json:"logTime"`
	SchemaPath     URL     `json:"schemaPath"`
}

func (m Message) Error() string {
	err := []string{}
	for _, e := range m.Errors {
		err = append(err, e.Error())
	}

	return strings.Join(err, ", ")
}

type Error struct {
	PropertyName  string      `json:"propertyName"`
	ErrorMessage  string      `json:"errorMessage"`
	ErrorCode     string      `json:"errorCode"`
	InputValue    interface{} `json:"inputValue"`
	DeveloperHint string      `json:"developerHint"`
}

func (e *Error) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err == nil {
		e.ErrorMessage = str
		log.Println("1")
		return nil
	}

	type alias Error
	a := alias(*e)
	err = json.Unmarshal(data, &a)
	if err != nil {
		log.Println("2")
		return err
	}

	*e = Error(a)
	log.Println("3")
	return nil
}

func (r Error) Error() string {
	if r.ErrorCode == "" && r.ErrorMessage != "" {
		return r.ErrorMessage
	}

	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}
