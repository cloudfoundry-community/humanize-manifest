package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Name         string    `yaml:"name"`
	DirectorUUID string    `yaml:"director_uuid"`
	Releases     []Release `yaml:"releases"`
}

type Release struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func fatalIf(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func main() {
	manifestPath := os.Args[1]
	if manifestPath == "" {
		os.Exit(2)
	}

	contents, err := ioutil.ReadFile(manifestPath)
	fatalIf(err)
	manifest := &Manifest{}
	yaml.Unmarshal(contents, manifest)

	str, err := yaml.Marshal(*manifest)
	fatalIf(err)
	fmt.Println(string(str))
}
