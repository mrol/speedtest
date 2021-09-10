package ookla

import (
	"github.com/mrol/speedtest/common"
	"github.com/showwin/speedtest-go/speedtest"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetDownloadSpeed() (float64, error) {
	servers, err := findServers()
	if err != nil {
		return 0, err
	}

	measures := make([]float64, 0, len(servers))

	for _, s := range servers {
		err = s.DownloadTest(false)
		if err != nil {
			return 0, err
		}

		measures = append(measures, s.DLSpeed)
	}

	return common.AvgFloat64(measures), nil
}

func (c *Client) GetUploadSpeed() (float64, error) {
	servers, err := findServers()
	if err != nil {
		return 0, err
	}

	measures := make([]float64, 0, len(servers))

	for _, s := range servers {
		err = s.UploadTest(false)
		if err != nil {
			return 0, err
		}

		measures = append(measures, s.ULSpeed)
	}

	return common.AvgFloat64(measures), nil
}

func findServers() (speedtest.Servers, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, err
	}

	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return nil, err
	}

	targets, err := serverList.FindServer([]int{})
	return targets, err
}
