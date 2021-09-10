package speedtest

const (
	OoklaProviderCode   = "ookla"
	FastComProviderCode = "fast"
)

var providerFinder finder = NewFinder()

// GetDownloadSpeed method checks download speed by specified provider at providerCode.
// Possible values you can see at consts
func GetDownloadSpeed(providerCode string) (float64, error) {
	client, err := providerFinder.GetClientByProviderCode(providerCode)
	if err != nil {
		return 0, err
	}

	return client.GetDownloadSpeed()
}

// GetUploadSpeed method checks upload speed by specified provider at providerCode.
// Possible values you can see at consts
func GetUploadSpeed(providerCode string) (float64, error) {
	client, err := providerFinder.GetClientByProviderCode(providerCode)
	if err != nil {
		return 0, err
	}

	return client.GetUploadSpeed()
}
