package main

import (
	"context"
	"fmt"
	"pcap-utils-go/gomods/sub_app/batch_img_conv"
	"pcap-utils-go/gomods/sub_app/name_verification"
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

// Prompts the user to choose a directory and returns the absolute path to it
func (a *App) Get_dir_path(title string) string {
	dir_path := fs_prompts.Get_directory_path(a.ctx, title)
	return dir_path
}

// Prompts the user to choose a file and returns the absolute path to it
func (a *App) Get_file_path(title string) string {
	file_path := fs_prompts.Get_file_path(a.ctx, title)
	return file_path
}

// Takes all the images from inp_dir and re-encodes them to jpeg with the .jpg file extension
func (a *App) Batch_img_conv(inp_dir string, out_dir string) ([]string, error) {
	undecoded_files, err := batch_img_conv.BatchConvert(inp_dir, out_dir, a.ctx)

	if err != nil {
		return nil, err
	}

	return undecoded_files, err
}

// Reads lines in a text file with the format [number] [text (supposedly a name)]
// and returns the names read
func (a *App) Read_verified_names(filepath string, preview_count int) ([]string, error) {
	verified_names, err := name_verification.ReadVerifiedNames(filepath, preview_count)

	if err != nil {
		return nil, err
	}

	return verified_names, nil
}

// Reads a .csv file and returns an object that stores the first, middle, and last name separately
func (a *App) Read_gforms_names(filepath string, preview_count int) ([]name_verification.CrimName, error) {
	gforms_names, err := name_verification.ReadGformsNames(filepath, preview_count)

	if err != nil {
		return nil, err
	}

	return gforms_names, nil
}

func (a *App) Verify_names(ver_names_fp string, gforms_names_fp string, output_fp string) ([]name_verification.ForManualReview, error) {
	for_man_review, err := name_verification.CompareNames(ver_names_fp, gforms_names_fp)

	if err != nil {
		return nil, err
	}

	// TODO: Write the manual matches for review to a file

	return for_man_review, nil
}
