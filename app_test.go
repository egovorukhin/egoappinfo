package egoappinfo

import (
	"fmt"
	"testing"
)

func TestAppInfo(t *testing.T) {
	SetName("EgoAppInfo")
	SetVersion(1, 2, 3)
	app := GetApplication()
	fmt.Printf("name: %s\n", app.Name)
	fmt.Printf("version: %s\n", app.Version.String())
	fmt.Printf("hostname: %s\n", app.Hostname)
	fmt.Printf("system: %s\n", app.System)
	fmt.Printf("filename: %s\n", app.Executable.File)
	fmt.Printf("filedir: %s\n", app.Executable.Dir)
}
