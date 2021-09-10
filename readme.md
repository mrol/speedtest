Package provides methods to measure download and upload speed by different providers.

## How to use
```go
import (
	"fmt"
	"os"

	"github.com/mrol/speedtest"
)

func main() {
	speed, err := speedtest.GetDownloadSpeed(speedtest.OoklaProviderCode)
	if err != nil {
		fmt.Printf("Error while measuring speed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Your download speed is %.2f Mbps", speed)
}
```

## Enabled providers
```go
const (
	OoklaProviderCode   = "ookla"
	FastComProviderCode = "fast"
)
```
Note, that upload test is not provided for fast.com by this package.