package ookla

import (
	"github.com/mrol/speedtest/common"
)

type Client struct {
	serverFinder serverFinder
}

func New() *Client {
	return &Client{
		serverFinder: newServerFinder(),
	}
}

func (c *Client) GetDownloadSpeed() (float64, error) {
	servers, err := c.serverFinder.FindServers()
	if err != nil {
		return 0, err
	}

	measures := make([]float64, 0, len(servers))

	for _, s := range servers {
		speed, err := s.DownloadTest()
		if err != nil {
			return 0, err
		}

		measures = append(measures, speed)
	}

	return common.AvgFloat64(measures), nil
}

func (c *Client) GetUploadSpeed() (float64, error) {
	servers, err := c.serverFinder.FindServers()
	if err != nil {
		return 0, err
	}

	measures := make([]float64, 0, len(servers))

	for _, s := range servers {
		speed, err := s.UploadTest()
		if err != nil {
			return 0, err
		}

		measures = append(measures, speed)
	}

	return common.AvgFloat64(measures), nil
}
