package main

import (
	"changeme/pkg/sys"
	"context"
	"encoding/json"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
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
