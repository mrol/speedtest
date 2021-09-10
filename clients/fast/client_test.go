package fast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetUploadSpeed(t *testing.T) {
	c := NewClient()

	_, err := c.GetUploadSpeed()

	assert.ErrorIs(t, err, ErrMethodNotProvided)
}

// Not so good test because of external network dependencies
func TestClient_GetDownloadSpeed(t *testing.T) {
	c := NewClient()
	res, err := c.GetDownloadSpeed()

	require.NoError(t, err)
	assert.NotEmpty(t, res)
}
