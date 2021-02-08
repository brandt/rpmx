package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sassoftware/go-rpmutils"
)

func usage() {
	fmt.Println("Usage: rpmx <src_file> [dest_dir]")
}

func main() {
	// Check for -h / -? / --help flags
	for _, e := range os.Args[1:] {
		if e == "-h" || e == "-?" || e == "--help" {
			usage()
			os.Exit(0)
		}
	}

	var rpmPath string
	var destDir string

	switch len(os.Args) {
	case 2:
		rpmPath = os.Args[1]
	case 3:
		rpmPath = os.Args[1]
		destDir = os.Args[2]
	default:
		fmt.Println("Invalid number of arguments.")
		usage()
		os.Exit(2)
	}

	if _, err := os.Stat(rpmPath); os.IsNotExist(err) {
		fmt.Println("Error: RPM file does not exist.")
		os.Exit(1)
	}

	f, err := os.Open(rpmPath)
	if err != nil {
		fmt.Printf("Error opening RPM file: %s\n", err)
		os.Exit(1)
	}

	rpm, err := rpmutils.ReadRpm(f)
	if err != nil {
		fmt.Printf("Error reading RPM: %s\n", err)
		os.Exit(1)
	}

	// Default to directory with same name as RPM w/o the .src.rpm/.rpm extension
	if destDir == "" {
		name := strings.TrimSuffix(filepath.Base(rpmPath), ".rpm")
		name = strings.TrimSuffix(name, ".src")
		destDir = filepath.Join(filepath.Dir(rpmPath), name)
	}

	if _, err := os.Stat(destDir); err == nil {
		fmt.Printf("Error: Destination already exists: %s\n", destDir)
		os.Exit(1)
	}

	fmt.Printf("Extracting to directory: %s\n", destDir)

	if err := rpm.ExpandPayload(destDir); err != nil {
		fmt.Printf("Error expanding package: %s\n", err)
		os.Exit(1)
	}
}
