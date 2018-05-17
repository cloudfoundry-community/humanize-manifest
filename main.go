package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

var debug = false

func main() {
	manifestPath := parseArguments()

	originalYaml, err := ioutil.ReadFile(manifestPath)
	fatalIf(err)

	orderedYaml := reorderKnownSchema(originalYaml)

	orderedCompleteYaml := addNodesThatAreUnknownToSchema(originalYaml, orderedYaml)

	fmt.Println("---")
	fmt.Print(string(orderedCompleteYaml))
}

func parseArguments() string {
	if len(os.Args) == 0 {
		os.Exit(2)
	}
	manifestArgIdx := 1
	if os.Args[1] == "-d" {
		fmt.Fprintf(os.Stderr, "DEBUG: activating debug mode\n")
		debug = true
		if len(os.Args) < 3 {
			os.Exit(2)
		}
		manifestArgIdx = 2
	}
	manifestPath := os.Args[manifestArgIdx]
	if manifestPath == "" {
		os.Exit(2)
	}
	return manifestPath
}

func reorderKnownSchema(originalYaml []byte) []byte {
	manifest := &Manifest{}
	err := yaml.Unmarshal(originalYaml, manifest)
	fatalIf(err)

	yamlOutput, err := yaml.Marshal(*manifest)
	fatalIf(err)
	// if debug {
	// 	fmt.Fprintf(os.Stderr, "---\n")
	// 	fmt.Fprintf(os.Stderr, string(yamlOutput))
	// }
	return yamlOutput
}

func addNodesThatAreUnknownToSchema(originalYaml []byte, orderedYaml []byte) []byte {
	completeTree := &yaml.MapSlice{}
	err := yaml.Unmarshal(originalYaml, &completeTree)
	fatalIf(err)

	orderedTree := &yaml.MapSlice{}
	err = yaml.Unmarshal(orderedYaml, &orderedTree)
	fatalIf(err)

	storeDest := func(modified yaml.MapSlice) {
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: storing modified MapSlice at path '/' (root MapSlice)\n")
		}
		orderedTree = &modified
	}
	err = appendMissingNodes(*completeTree, *orderedTree, storeDest, "/")
	fatalIf(err)

	orderedCompleteYaml, err := yaml.Marshal(orderedTree)
	fatalIf(err)

	return orderedCompleteYaml
}
