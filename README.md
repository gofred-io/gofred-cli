# Gofred CLI

A powerful command-line tool for creating and running Go WebAssembly applications with hot reload capabilities.

## Features

- 🚀 **Quick App Creation**: Generate new Go WebAssembly applications with a single command
- 🔥 **Hot Reload**: Automatic recompilation and browser refresh on file changes
- 🌐 **Web Server**: Built-in development server with WebSocket support
- 📦 **Cross-Platform**: Support for Windows, macOS, and Linux (AMD64 and ARM64)
- 🔧 **VS Code Integration**: Automatic workspace configuration for optimal development experience
- 📱 **Responsive Development**: Live reload across local and network devices

## Installation

### Quick Install (Recommended)

Run the installation script:

```bash
curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/refs/heads/master/install.sh | bash
```

This script will:
- Detect your operating system and architecture
- Download the appropriate binary
- Install it to `~/.local/bin` (or `~/AppData/Local/bin` on Windows)
- Add the directory to your PATH if needed

### Manual Installation

1. Download the appropriate binary for your platform from the [releases page](https://github.com/gofred-io/gofred-cli/releases)
2. Extract the binary to a directory in your PATH
3. Make it executable (on Unix systems): `chmod +x gofred`

### Verify Installation

```bash
gofred version
```

## Usage

### Create a New Application

Create a new Go WebAssembly application:

```bash
gofred app create my-app --package my-app
```

**Options:**
- `path`: Path where to create the application (required, positional argument)
- `--package, -p`: Package name for the application (default: "gofred-app")

**Example:**
```bash
gofred app create hello-world --package hello-world
```

This will create:
- A `main.go` file with a basic "Hello, world" application
- A `web/` folder with necessary HTML, CSS, and JavaScript files
- A `.vscode/settings.json` file configured for Go WebAssembly development
- A `go.mod` file initialized with the specified package name

### Run Your Application

Navigate to your application directory and start the development server:

```bash
cd my-app
gofred app run
```

This will:
- Compile your Go code to WebAssembly
- Start a development server (usually on `http://localhost:PORT`)
- Open your browser automatically
- Watch for file changes and automatically recompile and reload

### Available Commands

```bash
gofred --help                    # Show help information
gofred version                   # Show version information
gofred update                    # Update to the latest version
gofred app create <path> [flags] # Create a new application
gofred app run                   # Run the application (must be in app directory)
```

### Global Flags

- `--offline, -o`: Run in offline mode (uses embedded web assets instead of downloading from CDN)

## Development Workflow

1. **Create your app**: `gofred app create my-app --package my-app`
2. **Navigate to the directory**: `cd my-app`
3. **Start development**: `gofred app run`
4. **Edit your code**: Modify `main.go` or other `.go` files
5. **See changes instantly**: The app automatically recompiles and reloads in your browser

## Project Structure

When you create a new application, the following structure is generated:

```
my-app/
├── main.go                 # Your main application file
├── go.mod                  # Go module file
├── web/                    # Web assets
│   ├── index.html         # Main HTML file
│   ├── index.css          # Styles
│   ├── index.js           # JavaScript runtime
│   ├── wasm_exec.js       # WebAssembly execution helper
│   └── env.js             # Environment configuration (auto-generated)
├── .vscode/               # VS Code configuration
│   └── settings.json      # Go WebAssembly settings
└── web/main.wasm          # Compiled WebAssembly binary (auto-generated)
```

## Example Application

The generated `main.go` file contains a simple example:

```go
package main

import (
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/application"
)

func main() {
    app := text.New("Hello, world")
    application.Run(app)
}
```

## Requirements

- Go 1.25.1 or later
- A modern web browser that supports WebAssembly
- For installation: `curl`, `tar`, and `jq` (automatically installed by the install script)

## Troubleshooting

### Installation Issues

- **Permission denied**: Make sure the installation directory is writable
- **Command not found**: Ensure the binary directory is in your PATH
- **Download failed**: Check your internet connection and try again

### Development Issues

- **Compilation errors**: Check your Go code for syntax errors
- **Browser not opening**: Manually navigate to the URL shown in the terminal
- **Hot reload not working**: Ensure you're running the command from within your app directory

### Offline Mode

If you're having network issues, you can run in offline mode:

```bash
gofred app create my-app --package my-app --offline
gofred app run --offline
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please:
1. Check the troubleshooting section above
2. Search existing [GitHub Issues](https://github.com/gofred-io/gofred-cli/issues)
3. Create a new issue with detailed information about your problem

---

**Happy coding with Gofred! 🚀**
