package main

import (
	"fmt"
	"minedle/internal"
)

func main() {
	apiURL := fmt.Sprintf("%s/%s", internal.GithubAPIBase, internal.TargetPath)

	fmt.Println("Downloading:", apiURL)
	if err := internal.DownloadDir(apiURL, internal.OutputDir); err != nil {
		panic(err)
	}

	fmt.Println("Download complete.")
}
