package main

import (
	"context"
	"fmt"
	"pcap-utils-go/gomods/sub_app/batch_img_conv"
	"pcap-utils-go/gomods/utils/fs_prompts"
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

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Get_dir_path(title string) string {
	dir_path := fs_prompts.Get_directory_path(a.ctx, title)
	return dir_path
}

func (a *App) Batch_img_conv(inp_dir string, out_dir string) ([]string, error) {
	undecoded_files, err := batch_img_conv.BatchConvert(inp_dir, out_dir, a.ctx)

	if err != nil {
		return nil, err
	}

	return undecoded_files, err
}
