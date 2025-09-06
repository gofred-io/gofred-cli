package flags

var (
	offline bool
)

func IsOffline() bool {
	return offline
}

func OfflineRef() *bool {
	return &offline
}
