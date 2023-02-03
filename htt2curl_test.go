package http2curl

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetCurlCommand(t *testing.T) {
	t.Parallel()

	t.Run("test1", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "https://example.com/index?a=1&b=2", nil)
		req.Header.Set("a", "2")
		c, err := GetCurlCommand(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'GET' -H 'A: 2' 'https://example.com/index?a=1&b=2' --compressed", c.String())
	})

	t.Run("test2", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "https://example.com/index", bytes.NewBufferString(`{"a":"b"}`))
		req.Header.Set("Content-Type", "application/json")

		command, err := GetCurlCommand(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'POST' -d '{\"a\":\"b\"}' -H 'Content-Type: application/json' 'https://example.com/index' --compressed", command.String())
	})
}
