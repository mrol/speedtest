package speedtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getClientByProviderCode_Success(t *testing.T) {
	f := NewFinder()
	cl, err := f.GetClientByProviderCode(OoklaProviderCode)

	assert.Nil(t, err)
	assert.Equal(t, providers[OoklaProviderCode], cl)
}

func Test_getClientByProviderCode_InvalidProvider(t *testing.T) {
	f := NewFinder()
	cl, err := f.GetClientByProviderCode("invalid")

	assert.Nil(t, cl)
	assert.ErrorIs(t, err, ErrInvalidProvider)
}
