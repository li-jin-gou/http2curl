package http2curl

import (
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/protocol"
)

// GetCurlCommandHertz returns a CurlCommand corresponding to an http.Request
func GetCurlCommandHertz(req *protocol.Request) (*CurlCommand, error) {
	if req.URI() == nil {
		return nil, ErrorURINull
	}

	command := CurlCommand{}

	command.append("curl")

	requestURL := req.URI().String()

	if string(req.URI().Scheme()) == "https" {
		command.append("-k")
	}

	command.append("-X", bashEscape(b2s(req.Method())))

	body := req.Body()
	if len(body) != 0 {
		bodyEscaped := bashEscape(b2s(body))
		command.append("-d", bodyEscaped)

	}

	headers := make(map[string][]string)

	req.Header.VisitAll(func(key, value []byte) {
		_, ok := headers[b2s(key)]
		if !ok {
			headers[b2s(key)] = []string{b2s(value)}
		} else {
			headers[b2s(key)] = append(headers[b2s(key)], b2s(value))
		}
	})

	for key, values := range headers {
		command.append("-H", bashEscape(fmt.Sprintf("%s: %s", key, strings.Join(values, " "))))
	}

	command.append(bashEscape(requestURL))

	command.append("--compressed")

	return &command, nil
}