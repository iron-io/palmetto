package palmetto

import (
	"testing"

	"github.com/iron-io/palmetto/Godeps/_workspace/src/github.com/arschles/assert"
)

func TestGetNotFound(t *testing.T) {
	c := JSONClient(ServerURL)
	resp, err := c.Get("TestGetNotFound")
	assert.Err(t, ErrNotFound, err)
	assert.True(t, resp == nil, "response was not nil when key not found")
}

func TestGetPutGet(t *testing.T) {
	c := JSONClient(ServerURL)
	resp, err := c.Get("TestGetPutGet")
	assert.Err(t, ErrNotFound, err)
	assert.True(t, resp == nil, "response was not nil when key not found")
}
