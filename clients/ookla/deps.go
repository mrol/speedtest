//go:generate mockgen -source=deps.go -destination=deps_mock.go -package=ookla
package ookla

import "github.com/showwin/speedtest-go/speedtest"

//serverFinder finds servers to check speed
type serverFinder interface {
	FindServers() ([]server, error)
}

//server provides methods to check speed
type server interface {
	DownloadTest() (float64, error)
	UploadTest() (float64, error)
}

type serverFinderImpl struct {
}

func newServerFinder() *serverFinderImpl {
	return &serverFinderImpl{}
}

func (s *serverFinderImpl) FindServers() ([]server, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, err
	}

	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return nil, err
	}

	targets, err := serverList.FindServer([]int{})

	servers := make([]server, 0, len(targets))

	for _, target := range targets {
		servers = append(servers, newServer(target))
	}

	return servers, err
}

type serverImpl struct {
	speedtestServer *speedtest.Server
}

func newServer(speedtestServer *speedtest.Server) *serverImpl {
	return &serverImpl{speedtestServer: speedtestServer}
}

func (s *serverImpl) DownloadTest() (float64, error) {
	err := s.speedtestServer.DownloadTest(false)
	return s.speedtestServer.DLSpeed, err
}

func (s *serverImpl) UploadTest() (float64, error) {
	err := s.speedtestServer.UploadTest(false)
	return s.speedtestServer.ULSpeed, err
}
