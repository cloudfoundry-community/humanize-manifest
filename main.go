package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"launchpad.net/goyaml"

	"github.com/cloudfoundry-community/gogobosh/local"
	"github.com/cloudfoundry-community/gogobosh/models"
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
	manifestPath := os.Args[0]
	if manifestPath != "" {
		manifestPath = currentBoshManifest()
	}

	contents, err := ioutil.ReadFile(manifestPath)
	fatalIf(err)
	manifest := &models.DeploymentManifest{}
	goyaml.Unmarshal(contents, manifest)

	str, err := goyaml.Marshal(*manifest)
	fatalIf(err)
	fmt.Println(string(str))
}
