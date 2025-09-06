package flags

var (
	defaultCDNURL = "https://cdn.gofred.io"
	binaryURL     = "https://github.com/gofred-io/gofred-cli/releases/download"
)

func DefaultCDNURL() string {
	return defaultCDNURL
}

func BinaryURL() string {
	return binaryURL
}
