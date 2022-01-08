package main

import (
	"changeme/pkg/sys"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Department struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	sys := &sys.Sys{}
	user := sys.All().User
	return fmt.Sprintf("Welcome %s. What department do you work for as we will help setup your machine.", user.Name)
}

func (a *App) Sysvars() string {
	sys := &sys.Sys{}
	b, err := json.Marshal(sys.All())
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("error: could not convert to json")
	}
	return string(b)
}

func departments() []Department {

	var depts = []Department{
		Department{
			Value: "dev",
			Label: "Dev",
		},
		Department{
			Value: "devops",
			Label: "Dev Ops",
		},
		Department{
			Value: "marketing",
			Label: "Marketing",
		},
		Department{
			Value: "strategy",
			Label: "Strategy",
		},
	}

	return depts
}

func (a *App) GetDepartments() string {
	depts, err := json.Marshal(departments())
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("error: could not convert to json")
	}
	fmt.Println(string(depts))
	return string(depts)
}

func (a *App) Install(department string) string {
	sys := &sys.Sys{}
	switch department {
	case "marketing":
	case "strategy":
	case "dev":
		var vboxUrl = ""
		if runtime.GOOS == "windows" {
			vboxUrl = "https://download.virtualbox.org/virtualbox/6.1.30/VirtualBox-6.1.30-148432-Win.exe"
		}
		if runtime.GOOS == "darwin" {
			vboxUrl = "https://download.virtualbox.org/virtualbox/6.1.30/VirtualBox-6.1.30-148432-OSX.dmg"
		}
		if vboxUrl != "" {
			fileName := path.Base(vboxUrl)
			sys.GetRemoteFile(fileName, vboxUrl)
			sys.Exec(fileName)
		}
	case "devops":
	default:
		fmt.Println("no valid option specified")
		return "Fail"
	}
	return "Success"
}
