package update

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/fynelabs/selfupdate"
	"github.com/gofred-io/gofred-cli/flags"
)

type IndexFile struct {
	Version string `json:"version"`
}

func CheckForUpdates(verbose bool) {
	var (
		defaultCDNURL = flags.DefaultCDNURL()
		binaryBaseURL = flags.BinaryURL()
	)

	if flags.IsOffline() {
		return
	}

	// Download index.json from the CDN
	response, err := http.Get(fmt.Sprintf("%s/cli/index.json", defaultCDNURL))
	if err != nil {
		log.Fatalf("Error downloading index.json: %v", err)
	}
	defer response.Body.Close()

	var indexFile IndexFile
	err = json.NewDecoder(response.Body).Decode(&indexFile)
	if err != nil {
		log.Fatalf("Error decoding index.json: %v", err)
	}

	if indexFile.Version == flags.Version {
		if verbose {
			fmt.Println("You already have the latest version")
		}
		return
	}

	// Download the gofred binary from the CDN
	fmt.Printf("Found new version, updating to version %s...\n", indexFile.Version)
	fmt.Println("Downloading gofred binary, this may take a while...")

	binaryURL := fmt.Sprintf(
		"%s/v%s/gofred-%s-%s",
		binaryBaseURL,
		indexFile.Version,
		runtime.GOOS,
		runtime.GOARCH,
	)

	if runtime.GOOS == "windows" {
		binaryURL += ".exe"
	}

	binaryResponse, err := http.Get(binaryURL)
	if err != nil {
		log.Fatalf("Error downloading gofred binary: %v", err)
	}
	defer binaryResponse.Body.Close()

	err = selfupdate.Apply(binaryResponse.Body, selfupdate.Options{})
	if err != nil {
		log.Fatalf("Error applying update: %v", err)
	}

	fmt.Println("Updates applied, please run `gofred` again to start the application")
	os.Exit(0)
}
