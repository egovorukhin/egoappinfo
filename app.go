package egoappinfo

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const unknown = "unknown"

type Application struct {
	Name       string
	Version    Version
	Hostname   string
	System     string
	Executable Executable
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

type Executable struct {
	File string
	Dir  string
}

var application = New()

func New() *Application {

	hostname, _ := os.Hostname()
	eDir := unknown
	eFile, err := os.Executable()
	if err != nil {
		eFile = unknown
	} else {
		eDir = filepath.Dir(eFile)
	}

	ext := filepath.Ext(eFile)
	fn := filepath.Base(eFile)
	name := strings.Title(fn[:len(fn)-len(ext)])

	system := runtime.GOOS
	switch system {
	case "windows":
		system = "Windows"
	case "darwin":
		system = "MAC operating system"
	case "linux":
		system = "Linux"
	}

	return &Application{
		Name: name,
		Version: Version{
			Major: 0,
			Minor: 0,
			Patch: 1,
		},
		Hostname: hostname,
		System:   system,
		Executable: Executable{
			File: eFile,
			Dir:  eDir,
		},
	}
}

func SetName(name string) {
	application.Name = name
}

func SetVersion(major, minor, patch int) {
	application.Version = Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func GetApplication() *Application {
	return application
}

func GetUserAgent() string {
	return application.Name + "/" + application.Version.String()
}

func GetApplicationName() string {
	return application.Name
}

func GetVersion() Version {
	return application.Version
}

func GetExecutable() Executable {
	return application.Executable
}
