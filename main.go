package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Version the version number for this program
var Version = ""

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
		fmt.Fprintf(os.Stderr, "ERROR: no arguments. Usage: humanize-manifest [-d] <filename>\n")
		os.Exit(2)
	}
	manifestArgIdx := 1
	if os.Args[1] == "-v" {
		if Version == "" || Version == "dev" {
			fmt.Printf("humanize-manifest (development)\n")
		} else {
			fmt.Printf("humanize-manifest v%s\n", Version)
		}
		os.Exit(0)
	} else if os.Args[1] == "-d" {
		fmt.Fprintf(os.Stderr, "DEBUG: activating debug mode\n")
		debug = true
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "ERROR: too few arguments.\n")
			os.Exit(2)
		}
		manifestArgIdx = 2
	}
	manifestPath := os.Args[manifestArgIdx]
	if manifestPath == "" {
		fmt.Fprintf(os.Stderr, "ERROR: empty path for file.\n")
		os.Exit(2)
	}
	_, err := os.Stat(manifestPath)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "ERROR: file does not exist: '%s'\n", manifestPath)
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
