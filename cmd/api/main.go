package main

import (
	"fmt"
	"minedle/internal"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiURL := fmt.Sprintf("%s/%s", internal.GithubAPIBase, internal.TargetPath)

	fmt.Println("Downloading:", apiURL)
	if err := internal.DownloadDir(apiURL, internal.OutputDir); err != nil {
		panic(err)
	}

	data, err := internal.ParseMinecraftData("1.21.1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

	fmt.Println("Download complete.")
}
