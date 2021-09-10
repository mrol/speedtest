package ookla

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Not so good test because of external network dependencies
func TestClient_GetDownloadSpeed(t *testing.T) {
	c := New()
	res, err := c.GetDownloadSpeed()

	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

// Not so good test because of external network dependencies
func TestClient_GetUploadSpeed(t *testing.T) {
	c := New()
	res, err := c.GetUploadSpeed()

	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestClient_GetDownloadSpeed_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	speed := 32.7

	mockServer := NewMockserver(ctrl)
	mockServer.EXPECT().DownloadTest().Return(speed, nil)

	mockServerFinder := NewMockserverFinder(ctrl)
	mockServerFinder.EXPECT().FindServers().Return([]server{mockServer}, nil)

	client := &Client{serverFinder: mockServerFinder}
	result, err := client.GetDownloadSpeed()

	require.NoError(t, err)
	assert.Equal(t, speed, result)
}

func TestClient_GetDownloadSpeed_FindServersError(t *testing.T) {
	ctrl := gomock.NewController(t)

	findServersErr := errors.New("fatal error")

	mockServerFinder := NewMockserverFinder(ctrl)
	mockServerFinder.EXPECT().FindServers().Return(nil, findServersErr)

	client := &Client{serverFinder: mockServerFinder}
	result, err := client.GetDownloadSpeed()

	assert.Empty(t, result)
	assert.ErrorIs(t, err, findServersErr)
}

func TestClient_GetDownloadSpeed_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)

	serverErr := errors.New("fatal error")

	mockServer := NewMockserver(ctrl)
	mockServer.EXPECT().DownloadTest().Return(0.0, serverErr)

	mockServerFinder := NewMockserverFinder(ctrl)
	mockServerFinder.EXPECT().FindServers().Return([]server{mockServer}, nil)

	client := &Client{serverFinder: mockServerFinder}
	result, err := client.GetDownloadSpeed()

	assert.Empty(t, result)
	assert.ErrorIs(t, err, serverErr)
}

func TestClient_GetUploadSpeed_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	speed := 32.7

	mockServer := NewMockserver(ctrl)
	mockServer.EXPECT().UploadTest().Return(speed, nil)

	mockServerFinder := NewMockserverFinder(ctrl)
	mockServerFinder.EXPECT().FindServers().Return([]server{mockServer}, nil)

	client := &Client{serverFinder: mockServerFinder}
	result, err := client.GetUploadSpeed()

	require.NoError(t, err)
	assert.Equal(t, speed, result)
}

func TestClient_GetUploadSpeed_FindServersError(t *testing.T) {
	ctrl := gomock.NewController(t)

	findServersErr := errors.New("fatal error")

	mockServerFinder := NewMockserverFinder(ctrl)
	mockServerFinder.EXPECT().FindServers().Return(nil, findServersErr)

	client := &Client{serverFinder: mockServerFinder}
	result, err := client.GetUploadSpeed()

	assert.Empty(t, result)
	assert.ErrorIs(t, err, findServersErr)
}

func TestClient_GetUploadSpeed_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)

	serverErr := errors.New("fatal error")

	mockServer := NewMockserver(ctrl)
	mockServer.EXPECT().UploadTest().Return(0.0, serverErr)

	mockServerFinder := NewMockserverFinder(ctrl)
	mockServerFinder.EXPECT().FindServers().Return([]server{mockServer}, nil)

	client := &Client{serverFinder: mockServerFinder}
	result, err := client.GetUploadSpeed()

	assert.Empty(t, result)
	assert.ErrorIs(t, err, serverErr)
}
