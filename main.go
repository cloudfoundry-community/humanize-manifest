package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"launchpad.net/goyaml"

	"github.com/cloudfoundry-community/gogobosh/local"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

type Manifest struct {
	Name         string    `yaml:"name"`
	DirectorUUID string    `yaml:"director_uuid"`
	Releases     []Release `yaml:"releases"`
}

type Release struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
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
	manifest := &Manifest{}
	goyaml.Unmarshal(contents, manifest)

	str, err := goyaml.Marshal(*manifest)
	fatalIf(err)
	fmt.Println(string(str))
}
