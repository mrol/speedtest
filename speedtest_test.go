package speedtest

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_clients "github.com/mrol/speedtest/clients/mocks"
	mock_providers "github.com/mrol/speedtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetDownloadSpeed_ErrInvalidProvider(t *testing.T) {
	ctrl := gomock.NewController(t)

	code := "test"

	provider := mock_providers.NewMockFinder(ctrl)
	provider.EXPECT().GetClientByProviderCode(code).Return(nil, ErrInvalidProvider)

	providerFinder = provider

	_, err := GetDownloadSpeed(code)

	assert.ErrorIs(t, err, ErrInvalidProvider)
}

func TestGetDownloadSpeed_ErrClient(t *testing.T) {
	ctrl := gomock.NewController(t)

	code := "test"
	networkErr := errors.New("network is broken")

	mockClient := mock_clients.NewMockClient(ctrl)
	mockClient.EXPECT().GetDownloadSpeed().Return(float64(0), networkErr)

	provider := mock_providers.NewMockFinder(ctrl)
	provider.EXPECT().GetClientByProviderCode(code).Return(mockClient, nil)

	providerFinder = provider

	_, err := GetDownloadSpeed(code)

	assert.ErrorIs(t, err, networkErr)
}

func TestGetDownloadSpeed_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	code := "test"

	speed := float64(35)

	mockClient := mock_clients.NewMockClient(ctrl)
	mockClient.EXPECT().GetDownloadSpeed().Return(speed, nil)

	provider := mock_providers.NewMockFinder(ctrl)
	provider.EXPECT().GetClientByProviderCode(code).Return(mockClient, nil)

	providerFinder = provider

	res, err := GetDownloadSpeed(code)

	assert.Nil(t, err)
	assert.Equal(t, speed, res)
}

func TestGetUploadSpeed_ErrInvalidProvider(t *testing.T) {
	ctrl := gomock.NewController(t)

	code := "test"

	provider := mock_providers.NewMockFinder(ctrl)
	provider.EXPECT().GetClientByProviderCode(code).Return(nil, ErrInvalidProvider)

	providerFinder = provider

	_, err := GetUploadSpeed(code)

	assert.ErrorIs(t, err, ErrInvalidProvider)
}

func TestGetUploadSpeed_ErrClient(t *testing.T) {
	ctrl := gomock.NewController(t)

	code := "test"
	networkErr := errors.New("network is broken")

	mockClient := mock_clients.NewMockClient(ctrl)
	mockClient.EXPECT().GetUploadSpeed().Return(float64(0), networkErr)

	provider := mock_providers.NewMockFinder(ctrl)
	provider.EXPECT().GetClientByProviderCode(code).Return(mockClient, nil)

	providerFinder = provider

	_, err := GetUploadSpeed(code)

	assert.ErrorIs(t, err, networkErr)
}

func TestGetUploadSpeed_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	code := "test"

	speed := float64(35)

	mockClient := mock_clients.NewMockClient(ctrl)
	mockClient.EXPECT().GetUploadSpeed().Return(speed, nil)

	provider := mock_providers.NewMockFinder(ctrl)
	provider.EXPECT().GetClientByProviderCode(code).Return(mockClient, nil)

	providerFinder = provider

	res, err := GetUploadSpeed(code)

	assert.Nil(t, err)
	assert.Equal(t, speed, res)
}

func BenchmarkGetDownloadSpeed_Ookla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := GetDownloadSpeed(OoklaProviderCode)
		if err != nil {
			b.Fatalf("%s", err)
		}
		b.ReportMetric(res, "Mbps")
	}
}

func BenchmarkGetDownloadSpeed_FastCom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := GetDownloadSpeed(FastComProviderCode)
		if err != nil {
			b.Fatalf("%s", err)
		}
		b.ReportMetric(res, "Mbps")
	}
}

func BenchmarkGetUploadSpeed_Ookla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := GetUploadSpeed(OoklaProviderCode)
		if err != nil {
			b.Fatalf("%s", err)
		}
		b.ReportMetric(res, "Mbps")
	}
}
