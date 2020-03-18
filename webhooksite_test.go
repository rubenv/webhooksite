package webhooksite

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhooks(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	c := New()
	assert.NotNil(c)

	token, err := c.CreateToken()
	assert.NoError(err)
	assert.NotNil(token)

	reqs, err := c.GetRequests(token.UUID)
	assert.NoError(err)
	assert.NotNil(reqs)
	assert.Len(reqs.Data, 0)

	// Do a webhook call
	resp, err := http.Get(c.url(fmt.Sprintf("/%s", token.UUID)))
	assert.NoError(err)
	assert.Equal(resp.StatusCode, http.StatusOK)
	defer resp.Body.Close()

	// Should have it now
	reqs, err = c.GetRequests(token.UUID)
	assert.NoError(err)
	assert.NotNil(reqs)
	assert.Len(reqs.Data, 1)
}
