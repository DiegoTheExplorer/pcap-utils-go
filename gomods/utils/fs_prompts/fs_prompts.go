package fs_prompts

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Hello() {
	fmt.Println("Hello")
}

func Get_directory_path(ctx context.Context, title string) string {
	dir_path, err := runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{
		Title: title,
	})

	if err != nil {
		return ""
	}

	return dir_path
}
