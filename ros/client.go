package ros

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/reedchan7/aliyungo/common"
	"github.com/reedchan7/aliyungo/util"
)

const (
	// ROSDefaultEndpoint is the default API endpoint of ROS services
	ROSDefaultEndpoint = "https://ros.aliyuncs.com"
	ROSAPIVersion      = "2015-09-01"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	endpoint        string
	Version         string
	debug           bool
	userAgent       string
	httpClient      *http.Client
}

type Response struct {
	RequestId string `json:"request_id"`
}

// NewClient creates a new instance of ROS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		endpoint:        ROSDefaultEndpoint,
		Version:         ROSAPIVersion,
		httpClient:      &http.Client{},
	}
}

func NewROSClientWithSecurityToken(accessKeyId string, accessKeySecret string, securityToken string) *Client {
	return &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		endpoint:        ROSDefaultEndpoint,
		Version:         ROSAPIVersion,
		httpClient:      &http.Client{},
	}
}

func NewClientForAussumeRole(accessKeyId, accessKeySecret, securityToken string) *Client {
	return &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		endpoint:        ROSDefaultEndpoint,
		Version:         ROSAPIVersion,
		httpClient:      &http.Client{},
	}
}

// SetDebug sets debug mode to log the request/response message
func (client *Client) SetDebug(debug bool) {
	client.debug = debug
}

// SetUserAgent sets user agent to log the request/response message
func (client *Client) SetUserAgent(userAgent string) {
	client.userAgent = userAgent
}

func (client *Client) SetSecurityToken(securityToken string) {
	client.SecurityToken = securityToken
}

// SetTransport sets transport to the http client
func (client *Client) SetTransport(transport http.RoundTripper) {
	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}
	client.httpClient.Transport = transport
}

type Request struct {
	Method          string
	URL             string
	Version         string
	Region          common.Region
	Signature       string
	SignatureMethod string
	SignatureNonce  string
	Timestamp       util.ISO6801Time
	Body            []byte
}

// Invoke sends the raw HTTP request for ROS services
func (client *Client) Invoke(region common.Region, method string, path string, query url.Values, args interface{}, response interface{}) error {

	var reqBody []byte
	var err error
	var contentType string
	var contentMD5 string

	if args != nil {
		reqBody, err = json.Marshal(args)
		if err != nil {
			return err
		}
		contentType = "application/json"
		hasher := md5.New()
		hasher.Write(reqBody)
		contentMD5 = base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	}

	requestURL := client.endpoint + path
	if query != nil && len(query) > 0 {
		requestURL = requestURL + "?" + util.Encode(query)
	}
	var bodyReader io.Reader
	if reqBody != nil {
		bodyReader = bytes.NewReader(reqBody)
	}
	httpReq, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return common.GetClientError(err)
	}

	if region != "" {
		httpReq.Header["x-acs-region-id"] = []string{string(region)}
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}
	if contentMD5 != "" {
		httpReq.Header.Set("Content-MD5", contentMD5)
	}
	// TODO move to util and add build val flag
	httpReq.Header.Set("Date", util.GetGMTime())
	httpReq.Header.Set("Accept", "application/json")
	// httpReq.Header.Set("x-acs-version", client.Version)
	httpReq.Header["x-acs-version"] = []string{client.Version}

	httpReq.Header["x-acs-signature-version"] = []string{"1.0"}
	httpReq.Header["x-acs-signature-nonce"] = []string{util.CreateRandomString()}
	httpReq.Header["x-acs-signature-method"] = []string{"HMAC-SHA1"}

	if client.userAgent != "" {
		httpReq.Header.Set("User-Agent", client.userAgent)
	}

	if client.SecurityToken != "" {
		httpReq.Header["x-acs-security-token"] = []string{client.SecurityToken}
	}

	client.signRequest(httpReq)

	t0 := time.Now()
	httpResp, err := client.httpClient.Do(httpReq)
	t1 := time.Now()
	if err != nil {
		return common.GetClientError(err)
	}
	statusCode := httpResp.StatusCode

	if client.debug {
		fmt.Printf("Invoke %s %s %d (%v)", method, requestURL, statusCode, t1.Sub(t0))
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return common.GetClientError(err)
	}

	if client.debug {
		var prettyJSON bytes.Buffer
		_ = json.Indent(&prettyJSON, body, "", "    ")
		log.Println(prettyJSON.String())
	}

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := common.ErrorResponse{}
		_ = json.Unmarshal(body, &errorResponse)
		cErr := &common.Error{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
		return cErr
	}

	if response != nil && len(body) > 0 {
		err = json.Unmarshal(body, response)
		// fmt.Printf("%++v", response)
		if err != nil {
			return common.GetClientError(err)
		}
	}

	return nil
}
