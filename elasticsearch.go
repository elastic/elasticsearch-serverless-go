// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package elasticsearch

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/elastic/elasticsearch-serverless-go/internal/version"
	"github.com/elastic/elasticsearch-serverless-go/typedapi"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	tpversion "github.com/elastic/elastic-transport-go/v8/elastictransport/version"
)

const (
	// Version returns the package version as a string.
	Version        = version.Client
	unknownProduct = "the client noticed that the server is not Elasticsearch and we do not support this unknown product"

	// HeaderClientMeta Key for the HTTP Header related to telemetry data sent with
	// each request to Elasticsearch.
	HeaderClientMeta = "x-elastic-client-meta"
)

var (
	userAgent           string
	reGoVersion         = regexp.MustCompile(`go(\d+\.\d+\..+)`)
	reMetaVersion       = regexp.MustCompile("([0-9.]+)(.*)")
	reServerlessVersion = regexp.MustCompile(`([0-9.]+)\+([0-9]+)(.*)`)
)

func init() {
	userAgent = initUserAgent()
}

// Config represents the client configuration.
type Config struct {
	Address  string // A list of Elasticsearch nodes to use.
	Username string // Username for HTTP Basic Authentication.
	Password string // Password for HTTP Basic Authentication.

	CloudID      string // Endpoint for the Elastic Service (https://elastic.co/cloud).
	APIKey       string // Base64-encoded token for authorization; if set, overrides username/password and service token.
	ServiceToken string // Service token for authorization; if set, overrides username/password.

	Header http.Header // Global HTTP request header.

	RetryOnStatus []int                           // List of status codes for retry. Default: 502, 503, 504.
	DisableRetry  bool                            // Default: false.
	MaxRetries    int                             // Default: 3.
	RetryOnError  func(*http.Request, error) bool // Optional function allowing to indicate which error should be retried. Default: nil.

	CompressRequestBody      bool // Default: false.
	CompressRequestBodyLevel int  // Default: gzip.DefaultCompression.

	DiscoverNodesOnStart  bool          // Discover nodes when initializing the client. Default: false.
	DiscoverNodesInterval time.Duration // Discover nodes periodically. Default: disabled.

	EnableDebugLogger bool // Enable the debug logging.

	DisableMetaHeader bool // Disable the additional "X-Elastic-Client-Meta" HTTP header.

	RetryBackoff func(attempt int) time.Duration // Optional backoff duration. Default: nil.

	Transport http.RoundTripper       // The HTTP transport object.
	Logger    elastictransport.Logger // The logger object.
}

// BaseClient represents the Elasticsearch client.
type BaseClient struct {
	Transport  elastictransport.Interface
	metaHeader string

	disableMetaHeader   bool
	productCheckMu      sync.RWMutex
	productCheckSuccess bool
}

// Client represents the Functional Options API.
type Client struct {
	BaseClient
}

// TypedClient represents the Typed API.
type TypedClient struct {
	BaseClient
	*typedapi.API
}

// NewClient creates a new client with configuration from cfg.
//
// It will use http://localhost:9200 as the default address.
//
// It will use the ELASTICSEARCH_URL environment variable, if set,
// to configure the addresses; use a comma to separate multiple URLs.
//
// If cfg.Address or cfg.CloudID is set, the ELASTICSEARCH_URL
// environment variable is ignored.
//
// It's an error to set both cfg.Address and cfg.CloudID.
func NewClient(cfg Config) (*TypedClient, error) {
	tp, err := newTransport(cfg)
	if err != nil {
		return nil, err
	}

	client := &TypedClient{
		BaseClient: BaseClient{
			Transport:         tp,
			disableMetaHeader: cfg.DisableMetaHeader,
			metaHeader:        initMetaHeader(tp),
		},
	}
	client.API = typedapi.New(client)

	return client, nil
}

func newTransport(cfg Config) (*elastictransport.Client, error) {
	var addr string

	if cfg.Address == "" && cfg.CloudID == "" {
		addr = addrsFromEnvironment()
	} else {
		if cfg.Address != "" && cfg.CloudID != "" {
			return nil, errors.New("cannot create client: both Addresses and CloudID are set")
		}

		if cfg.CloudID != "" {
			cloudAddr, err := addrFromCloudID(cfg.CloudID)
			if err != nil {
				return nil, fmt.Errorf("cannot create client: cannot parse CloudID: %s", err)
			}
			addr = cloudAddr
		}

		if addr == "" {
			addr = cfg.Address
		}
	}

	if addr == "" {
		return nil, errors.New("cannot create client: address is empty")
	}

	endpoint, err := addrToURL(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %s", err)
	}

	// TODO(karmi): Refactor
	if endpoint.User != nil {
		cfg.Username = endpoint.User.Username()
		pw, _ := endpoint.User.Password()
		cfg.Password = pw
	}

	tp, err := elastictransport.New(elastictransport.Config{
		UserAgent: userAgent,

		URLs:         []*url.URL{endpoint},
		Username:     cfg.Username,
		Password:     cfg.Password,
		APIKey:       cfg.APIKey,
		ServiceToken: cfg.ServiceToken,

		Header: cfg.Header,

		RetryOnStatus: cfg.RetryOnStatus,
		DisableRetry:  cfg.DisableRetry,
		RetryOnError:  cfg.RetryOnError,
		MaxRetries:    cfg.MaxRetries,
		RetryBackoff:  cfg.RetryBackoff,

		CompressRequestBody:      cfg.CompressRequestBody,
		CompressRequestBodyLevel: cfg.CompressRequestBodyLevel,

		EnableDebugLogger: cfg.EnableDebugLogger,

		DiscoverNodesInterval: cfg.DiscoverNodesInterval,

		Transport: cfg.Transport,
		Logger:    cfg.Logger,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating transport: %s", err)
	}

	return tp, nil
}

// Perform delegates to Transport to execute a request and return a response.
func (c *BaseClient) Perform(req *http.Request) (*http.Response, error) {
	if !c.disableMetaHeader {
		existingMetaHeader := req.Header.Get(HeaderClientMeta)
		if existingMetaHeader != "" {
			req.Header.Set(HeaderClientMeta, strings.Join([]string{c.metaHeader, existingMetaHeader}, ","))
		} else {
			req.Header.Add(HeaderClientMeta, c.metaHeader)
		}
	} else {
		req.Header.Del(HeaderClientMeta)
	}

	// Retrieve the original request.
	res, err := c.Transport.Perform(req)

	// ResponseCheck, we run the header check on the first answer from ES.
	if err == nil && (res.StatusCode >= 200 && res.StatusCode < 300) {
		checkHeader := func() error { return genuineCheckHeader(res.Header) }
		if err := c.doProductCheck(checkHeader); err != nil {
			res.Body.Close()
			return nil, err
		}
	}

	return res, err
}

// doProductCheck calls f if there as not been a prior successful call to doProductCheck,
// returning nil otherwise.
func (c *BaseClient) doProductCheck(f func() error) error {
	c.productCheckMu.RLock()
	productCheckSuccess := c.productCheckSuccess
	c.productCheckMu.RUnlock()

	if productCheckSuccess {
		return nil
	}

	c.productCheckMu.Lock()
	defer c.productCheckMu.Unlock()

	if c.productCheckSuccess {
		return nil
	}

	if err := f(); err != nil {
		return err
	}

	c.productCheckSuccess = true

	return nil
}

// genuineCheckHeader validates the presence of the X-Elastic-Product header
func genuineCheckHeader(header http.Header) error {
	if header.Get("X-Elastic-Product") != "Elasticsearch" {
		return errors.New(unknownProduct)
	}
	return nil
}

// Metrics returns the client metrics.
func (c *BaseClient) Metrics() (elastictransport.Metrics, error) {
	if mt, ok := c.Transport.(elastictransport.Measurable); ok {
		return mt.Metrics()
	}
	return elastictransport.Metrics{}, errors.New("transport is missing method Metrics()")
}

// addrsFromEnvironment returns a list of addresses by splitting
// the ELASTICSEARCH_URL environment variable with comma, or an empty list.
func addrsFromEnvironment() string {
	var addr string

	if envURL, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok && envURL != "" {
		addr = strings.TrimSpace(envURL)
	}

	return addr
}

// addrsToURLs creates a list of url.URL structures from url list.
func addrToURL(addr string) (*url.URL, error) {
	u, err := url.Parse(strings.TrimRight(addr, "/"))
	if err != nil {
		return nil, fmt.Errorf("cannot parse url: %v", err)
	}
	return u, nil
}

// addrFromCloudID extracts the Elasticsearch URL from CloudID.
// See: https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html
func addrFromCloudID(input string) (string, error) {
	var scheme = "https://"

	values := strings.Split(input, ":")
	if len(values) != 2 {
		return "", fmt.Errorf("unexpected format: %q", input)
	}
	data, err := base64.StdEncoding.DecodeString(values[1])
	if err != nil {
		return "", err
	}
	parts := strings.Split(string(data), "$")

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid encoded value: %s", parts)
	}

	return fmt.Sprintf("%s%s.%s", scheme, parts[1], parts[0]), nil
}

func initUserAgent() string {
	var b strings.Builder

	b.WriteString("elasticsearch-serverless-go")
	b.WriteRune('/')
	b.WriteString(Version)
	b.WriteRune(' ')
	b.WriteRune('(')
	b.WriteString(runtime.GOOS)
	b.WriteRune(' ')
	b.WriteString(runtime.GOARCH)
	b.WriteString("; ")
	b.WriteString("Go ")
	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		b.WriteString(v)
	} else {
		b.WriteString(runtime.Version())
	}
	b.WriteRune(')')

	return b.String()
}

func initMetaHeader(transport interface{}) string {
	var b strings.Builder
	var strippedGoVersion string
	var strippedEsVersion string
	var strippedTransportVersion string

	strippedEsVersion = buildServerLessStrippedVersion(Version)
	strippedGoVersion = buildStrippedVersion(runtime.Version())

	if _, ok := transport.(*elastictransport.Client); ok {
		strippedTransportVersion = buildStrippedVersion(tpversion.Transport)
	} else {
		strippedTransportVersion = strippedEsVersion
	}

	var duos = [][]string{
		{
			"esv",
			strippedEsVersion,
		},
		{
			"go",
			strippedGoVersion,
		},
		{
			"t",
			strippedTransportVersion,
		},
		{
			"hc",
			strippedGoVersion,
		},
	}

	var arr []string
	for _, duo := range duos {
		arr = append(arr, strings.Join(duo, "="))
	}
	b.WriteString(strings.Join(arr, ","))

	return b.String()
}

func buildStrippedVersion(version string) string {
	v := reMetaVersion.FindStringSubmatch(version)

	if len(v) == 3 && !strings.Contains(version, "devel") {
		switch {
		case v[2] != "":
			return v[1] + "p"
		default:
			return v[1]
		}
	}

	return "0.0p"
}

func buildServerLessStrippedVersion(version string) string {
	v := reServerlessVersion.FindStringSubmatch(version)

	switch len(v) {
	case 3, 4:
		return v[1] + "+" + v[2]
	default:
		return "0.0p"
	}
}
