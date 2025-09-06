# Contributing to Gofred CLI

Thank you for your interest in contributing to Gofred CLI! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Submitting Changes](#submitting-changes)
- [Issue Guidelines](#issue-guidelines)
- [Pull Request Guidelines](#pull-request-guidelines)

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the maintainers.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Create a new branch for your feature or bugfix
4. Make your changes
5. Test your changes thoroughly
6. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.25.1 or later
- Git
- Make (optional, for using the Makefile)

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/gofred-io/gofred-cli.git
   cd gofred-cli
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   go build -o gofred .
   ```

   Or use the Makefile:
   ```bash
   make build
   ```

### Running Tests

```bash
go test ./...
```

### Running the CLI

After building, you can run the CLI:
```bash
./gofred --help
```

## Making Changes

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` to format your code
- Use `golint` or `golangci-lint` to check for issues
- Write clear, self-documenting code
- Add comments for public functions and complex logic

### Commit Messages

Use clear, descriptive commit messages:

- Use the imperative mood ("Add feature" not "Added feature")
- Keep the first line under 50 characters
- Add a blank line between the subject and body
- Use the body to explain what and why, not how

Example:
```
Add support for custom package names

The create command now accepts a --package flag to specify
a custom package name for the generated application.
This allows users to create apps with meaningful names
instead of the default "gofred-app".
```

### Testing

- Write tests for new functionality
- Ensure all existing tests pass
- Test on multiple platforms if possible
- Test edge cases and error conditions

## Submitting Changes

### Pull Request Process

1. **Create a feature branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following the guidelines above

3. **Test your changes** thoroughly

4. **Commit your changes** with clear commit messages

5. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Create a Pull Request** on GitHub

### Pull Request Guidelines

- Provide a clear description of what the PR does
- Reference any related issues
- Include screenshots or examples if applicable
- Ensure all CI checks pass
- Keep PRs focused and reasonably sized
- Respond to feedback promptly

### Review Process

- All PRs require review from maintainers
- Address feedback promptly
- Be open to suggestions and improvements
- Maintain a positive, collaborative attitude

## Issue Guidelines

### Before Creating an Issue

1. Search existing issues to avoid duplicates
2. Check if the issue has been fixed in the latest version
3. Gather relevant information about your environment

### Bug Reports

When reporting bugs, please include:

- **OS and version** (e.g., macOS 14.0, Ubuntu 22.04)
- **Go version** (`go version`)
- **Gofred CLI version** (`gofred version`)
- **Steps to reproduce** the issue
- **Expected behavior** vs **actual behavior**
- **Error messages** (if any)
- **Screenshots** (if applicable)

### Feature Requests

When requesting features, please include:

- **Use case** - Why is this feature needed?
- **Proposed solution** - How should it work?
- **Alternatives considered** - What other approaches were considered?
- **Additional context** - Any other relevant information

## Development Areas

We welcome contributions in these areas:

- **Core functionality** - CLI commands and features
- **Documentation** - README, code comments, examples
- **Testing** - Unit tests, integration tests
- **Build system** - CI/CD, cross-platform builds
- **User experience** - Error messages, help text, usability
- **Performance** - Optimization and efficiency improvements

## Release Process

Releases are managed by maintainers and follow semantic versioning:

- **MAJOR** - Breaking changes
- **MINOR** - New features (backward compatible)
- **PATCH** - Bug fixes (backward compatible)

## Getting Help

- **GitHub Issues** - For bugs and feature requests
- **GitHub Discussions** - For questions and general discussion
- **Code Review** - For feedback on your contributions

## Recognition

Contributors will be recognized in:
- Release notes
- CONTRIBUTORS.md file
- GitHub contributor list

Thank you for contributing to Gofred CLI! 🚀
