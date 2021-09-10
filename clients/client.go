//go:generate mockgen -source=client.go -destination=mocks/client_mock.go
package clients

//Client describes contract for the different speed measure providers
type Client interface {
	//GetDownloadSpeed returns download speed in Mbps or error
	GetDownloadSpeed() (float64, error)
	//GetUploadSpeed returns download speed in Mbps or error
	GetUploadSpeed() (float64, error)
}
