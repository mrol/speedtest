//go:generate mockgen -source=providers.go -destination=mocks/providers_mock.go
package speedtest

import (
	"errors"

	"github.com/mrol/speedtest/clients"
	"github.com/mrol/speedtest/clients/fast"
	"github.com/mrol/speedtest/clients/ookla"
)

// client finder, used mostly for mocks
type finder interface {
	GetClientByProviderCode(providerCode string) (clients.Client, error)
}

type finderImpl struct {
}

func NewFinder() *finderImpl {
	return &finderImpl{}
}

// Map of service providers and its clients
var providers = map[string]clients.Client{
	OoklaProviderCode:   ookla.NewClient(),
	FastComProviderCode: fast.NewClient(),
}

var ErrInvalidProvider = errors.New("invalid provider")

//GetClientByProviderCode find provider client by code, or return ErrInvalidProvider if provider not found
func (f *finderImpl) GetClientByProviderCode(providerCode string) (clients.Client, error) {
	client, ok := providers[providerCode]
	if !ok {
		return nil, ErrInvalidProvider
	}

	return client, nil
}
