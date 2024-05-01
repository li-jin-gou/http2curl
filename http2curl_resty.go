package http2curl

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"sort"
	"strings"
)

func GetCurlCommandResty(req *resty.Request) (*CurlCommand, error) {
	if req.RawRequest.URL == nil {
		return nil, ErrorURINull
	}

	command := CurlCommand{}
	command.append("curl")
	schema := req.RawRequest.URL.Scheme

	requestURL := req.URL
	if schema == "" {
		schema = "http"
		if req.RawRequest.TLS != nil {
			schema = "https"
		}
		requestURL = schema + "://" + req.RawRequest.Host + req.RawRequest.URL.Path
	}
	if schema == "https" {
		command.append("-k")
	}

	command.append("-X", bashEscape(req.Method))

	if req.Body != nil {
		var buff bytes.Buffer
		_, err := buff.ReadFrom(bytes.NewReader(req.Body.([]byte)))
		if err != nil {
			return nil, fmt.Errorf("GetCurlCommandResty:  buffer read from body error: %w", err)
		}

		req.Body = io.NopCloser(bytes.NewReader(buff.Bytes()))
		if len(buff.String()) > 0 {
			bodyEscaped := bashEscape(buff.String())
			command.append("-d", bodyEscaped)
		}
	}
	var keys []string
	fmt.Println(req.Header)
	fmt.Println(req.RawRequest.Header)

	for k := range req.RawRequest.Header {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		command.append("-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], " "))))
	}

	command.append(bashEscape(requestURL))
	command.append("--compressed")
	return &command, nil
}
