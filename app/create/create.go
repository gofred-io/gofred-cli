package create

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/gofred-io/gofred-cli/embed"
	"github.com/gofred-io/gofred-cli/flags"
	"github.com/gofred-io/gofred-cli/update"
	"github.com/gofred-io/gofred-cli/utils"
	"github.com/spf13/cobra"
)

const (
	defaultPkgName = "gofred-app"
)

var (
	appPath string
	pkgName string
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Run:   createApp,
	}
)

func Init(parentCmd *cobra.Command) {
	createCmd.Flags().StringVarP(&appPath, "path", "p", "", "Path to create the application")
	createCmd.Flags().StringVarP(&pkgName, "package", "n", defaultPkgName, "Package name for the application")
	parentCmd.AddCommand(createCmd)
}

func createApp(cmd *cobra.Command, args []string) {
	var (
		err error
	)

	if appPath == "" {
		fmt.Println("Path is required")
		cmd.Usage()
		return
	}

	update.CheckForUpdates(false)
	fmt.Printf("Creating application in %s\n", appPath)

	err = createVsCodeWorkspace()
	if err != nil {
		fmt.Printf("Failed to create .vscode workspace: %v\n", err)
		return
	}

	err = createMainFile()
	if err != nil {
		fmt.Printf("Failed to create main.go file: %v\n", err)
		return
	}

	err = createWebFolder()
	if err != nil {
		fmt.Printf("Failed to create web folder: %v\n", err)
		return
	}

	fmt.Println("Application created successfully")
	fmt.Println("Run `cd " + appPath + " && gofred app run` to start the application")
}

func createVsCodeWorkspace() error {
	var (
		content []byte
		err     error
	)

	content = []byte(`{
    "gopls": {
        "build.env": {
          "GOOS": "js",
          "GOARCH": "wasm",
        }
    }
}`)

	folderPath := filepath.Join(appPath, ".vscode")
	settingsPath := filepath.Join(folderPath, "settings.json")

	// Create the .vscode folder if it doesn't exist
	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}

	// Create a settings.json file if it doesn't exist
	if _, err = os.Stat(settingsPath); os.IsNotExist(err) {
		err = os.WriteFile(settingsPath, content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func createMainFile() error {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Id("app").Op(":=").Qual("github.com/gofred-io/gofred/foundation/text", "New").Call(Lit("Hello, world")),
		Qual("github.com/gofred-io/gofred/application", "Run").Call(Id("app")),
	)

	err := f.Save(filepath.Join(appPath, "main.go"))
	if err != nil {
		return err
	}

	// Run go mod init
	err = utils.ModInit(pkgName, appPath)
	if err != nil {
		return err
	}

	// Run go mod tidy
	err = utils.ModTidy(appPath)
	if err != nil {
		return err
	}

	return nil
}

func createWebFolder() error {
	var (
		err error
	)

	// Create the web folder if it doesn't exist
	folderPath := filepath.Join(appPath, "web")
	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}

	if flags.IsOffline() {
		return createWebFolderOffline(folderPath)
	}

	return createWebFolderOnline(folderPath)
}

func createWebFolderOffline(folderPath string) error {
	var (
		err     error
		content []byte
	)

	// Copy the web folder from the embed
	embedWeb := embed.Web()
	files, err := embedWeb.ReadDir("web")
	if err != nil {
		return err
	}

	for _, file := range files {
		content, err = embedWeb.ReadFile(filepath.Join("web", file.Name()))
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(folderPath, file.Name()), content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func createWebFolderOnline(folderPath string) error {
	var (
		defaultCDNURL = os.Getenv("CDN_URL")
	)

	// Download index.json from the CDN
	response, err := http.Get(fmt.Sprintf("%s/web/index.json", defaultCDNURL))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Parse the index.json file
	var indexFile IndexFile
	err = json.NewDecoder(response.Body).Decode(&indexFile)
	if err != nil {
		return err
	}

	// Download the files
	for _, file := range indexFile.Files {
		response, err := http.Get(fmt.Sprintf("%s/web/%s", defaultCDNURL, file))
		if err != nil {
			return err
		}
		defer response.Body.Close()

		content, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(folderPath, file), content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
