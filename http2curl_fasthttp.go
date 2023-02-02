package http2curl

import (
	"fmt"
	"sort"
	"strings"

	"github.com/valyala/fasthttp"
)

// GetCurlCommandFastHttp returns a CurlCommand corresponding to an http.Request
func GetCurlCommandFastHttp(req *fasthttp.Request) (*CurlCommand, error) {
	if req.URI() == nil {
		return nil, ErrorURINull
	}

	command := CurlCommand{}

	command.append("curl")

	requestURL := req.URI().String()

	if string(req.URI().Scheme()) == "https" {
		command.append("-k")
	}

	command.append("-X", bashEscape(b2s(req.Header.Method())))

	body := req.Body()
	if len(body) != 0 {
		bodyEscaped := bashEscape(b2s(body))
		command.append("-d", bodyEscaped)

	}

	headers := make(map[string][]string)
	keys := make([]string, 0)

	req.Header.VisitAll(func(key, value []byte) {
		_, ok := headers[b2s(key)]
		if !ok {
			keys = append(keys, b2s(key))
			headers[b2s(key)] = []string{b2s(value)}
		} else {
			headers[b2s(key)] = append(headers[b2s(key)], b2s(value))
		}
	})

	sort.Strings(keys)

	for _, key := range keys {
		command.append("-H", bashEscape(fmt.Sprintf("%s: %s", key, strings.Join(headers[key], " "))))
	}

	command.append(bashEscape(requestURL))

	command.append("--compressed")

	return &command, nil
}
