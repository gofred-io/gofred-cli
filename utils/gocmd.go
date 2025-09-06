package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ModInit(pkgName string, appPath string) error {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command("go", "mod", "init", pkgName)
	cmd.Dir = appPath
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("go mod init %s: %w: %s", pkgName, err, stderr.String())
	}
	return nil
}

func ModTidy(appPath string) error {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = appPath
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("go mod tidy: %w: %s", err, stderr.String())
	}
	return nil
}
