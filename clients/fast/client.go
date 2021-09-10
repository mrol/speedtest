package fast

import (
	"errors"

	"github.com/mrol/speedtest/common"
	"gopkg.in/ddo/go-fast.v0"
)

var ErrMethodNotProvided = errors.New("method not provided")

type Client struct {
	fastCom *fast.Fast
}

func NewClient() *Client {
	return &Client{fastCom: fast.New()}
}

func (c *Client) GetDownloadSpeed() (float64, error) {
	err := c.fastCom.Init()
	if err != nil {
		return 0, err
	}

	// get urls
	urls, err := c.fastCom.GetUrls()
	if err != nil {
		return 0, err
	}

	kbpsChan := make(chan float64)
	resultChan := make(chan float64)

	var measures []float64

	go func() {
		for kbps := range kbpsChan {
			measures = append(measures, kbps/1000)
		}

		resultChan <- common.AvgFloat64(measures)
	}()

	err = c.fastCom.Measure(urls, kbpsChan)
	if err != nil {
		return 0, err
	}
	return <-resultChan, nil
}

func (c *Client) GetUploadSpeed() (float64, error) {
	return 0, ErrMethodNotProvided
}
