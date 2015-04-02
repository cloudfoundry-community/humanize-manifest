package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/gogobosh/local"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func currentBoshManifest() string {
	configPath, err := local.DefaultBoshConfigPath()
	fatalIf(err)
	config, err := local.LoadBoshConfig(configPath)
	fatalIf(err)
	return config.CurrentDeploymentManifest()
}

func main() {
	manifest := os.Args[0]
	if manifest != "" {
		manifest = currentBoshManifest()
	}
	fmt.Println(manifest)
}
