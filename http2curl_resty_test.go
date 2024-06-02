package http2curl

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestGetCurlCommandResty(t *testing.T) {
	t.Parallel()

	t.Run("GET", func(t *testing.T) {
		var req *resty.Request
		client := resty.New()
		resp, _ := client.R().
			SetQueryParams(map[string]string{
				"a": "1",
				"b": "2",
			}).SetHeader("a", "2").Get("https://example.com/index")
		req = resp.Request

		c, err := GetCurlCommandResty(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'GET' -H 'A: 2' -H 'User-Agent: go-resty/2.12.0 (https://github.com/go-resty/resty)' 'https://example.com/index?a=1&b=2' --compressed", c.String())
	})
	t.Run("POST", func(t *testing.T) {
		var req *resty.Request
		client := resty.New()
		resp, _ := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody([]byte(`{"a":"b"}`)).
			Post("https://example.com/index")
		req = resp.Request
		command, err := GetCurlCommandResty(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'POST' -d '{\"a\":\"b\"}' -H 'Accept: application/json' -H 'Content-Type: application/json' -H 'User-Agent: go-resty/2.12.0 (https://github.com/go-resty/resty)' 'https://example.com/index' --compressed", command.String())
	})
}
