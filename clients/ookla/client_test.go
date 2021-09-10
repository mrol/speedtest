package ookla

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Not so good test because of external network dependencies
func TestClient_GetDownloadSpeed(t *testing.T) {
	c := NewClient()
	res, err := c.GetDownloadSpeed()

	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

// Not so good test because of external network dependencies
func TestClient_GetUploadSpeed(t *testing.T) {
	c := NewClient()
	res, err := c.GetUploadSpeed()

	require.NoError(t, err)
	assert.NotEmpty(t, res)
}
