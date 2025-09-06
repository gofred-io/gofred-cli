package main

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

func checkForUpdates() {
	var (
		defaultCDNURL = os.Getenv("CDN_URL")
	)

	if flags.IsOffline() {
		return
	}

	fmt.Println("Checking for updates...")

	// Download index.json from the CDN
	response, err := http.Get(fmt.Sprintf("%s/cli/index.json", defaultCDNURL))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer response.Body.Close()

	var indexFile IndexFile
	err = json.NewDecoder(response.Body).Decode(&indexFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if indexFile.Version == flags.Version {
		return
	}

	// Download the gofred binary from the CDN
	fmt.Println("Downloading gofred binary, this may take a while...")

	binaryURL := fmt.Sprintf(
		"%s/cli/%s/gofred-%s-%s",
		defaultCDNURL,
		indexFile.Version,
		runtime.GOOS,
		runtime.GOARCH,
	)

	if runtime.GOOS == "windows" {
		binaryURL += ".exe"
	}

	binaryResponse, err := http.Get(binaryURL)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer binaryResponse.Body.Close()

	err = selfupdate.Apply(binaryResponse.Body, selfupdate.Options{})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Updates applied, please run `gofred` again")
	os.Exit(0)
}
