package webhooksite

import (
	"net/http"
	"testing"
)

func notNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Fail()
	}
}

func noError(t *testing.T, err error) {
	if err != nil {
		t.Logf("Unexpected error: %s", err)
		t.Fail()
	}
}

func hasLen(t *testing.T, coll []*Request, count int) {
	if len(coll) != count {
		t.Logf("Unexpected length: %d != %d", len(coll), count)
		t.Fail()
	}
}

func TestWebhooks(t *testing.T) {
	t.Parallel()

	c := New()
	notNil(t, c)

	token, err := c.CreateToken()
	noError(t, err)
	notNil(t, token)

	reqs, err := c.GetRequests(token.UUID)
	noError(t, err)
	notNil(t, reqs)
	hasLen(t, reqs.Data, 0)

	// Do a webhook call
	resp, err := http.Get(token.URL)
	noError(t, err)
	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}
	defer resp.Body.Close()

	// Should have it now
	reqs, err = c.GetRequests(token.UUID)
	noError(t, err)
	notNil(t, reqs)
	hasLen(t, reqs.Data, 1)
}
