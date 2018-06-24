package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
)

type Version struct {
	major, minor, patch uint
	Version             string `json:"version"`
}

func NewVersion(maj, min, patch uint) *Version {
	return &Version{maj, min, patch, fmt.Sprintf("%d.%d.%d", maj, min, patch)}
}

func (v Version) PrintVersion() {
	fmt.Println(v.Version)
}

func (v Version) PrintJsonVersion() {
	version, _ := json.Marshal(v)
	fmt.Println(string(version))
}

func (v Version) GetVersion() string {
	return v.Version
}

func (v Version) GetJsonVersion() string {
	version, _ := json.Marshal(v)
	return string(version)
}

func ParseVersionComponents(version string) (uint, uint, uint) {
	re := regexp.MustCompile(`\.*`)
	var a, b, c uint
	fmt.Sscanf(re.ReplaceAllString(version, " "), "%d %d %d", &a, &b, &c)
	return a, b, c
}

func ParseVersion(version string) *Version {
	re := regexp.MustCompile(`\.*`)
	var a, b, c uint
	fmt.Sscanf(re.ReplaceAllString(version, " "), "%d %d %d", &a, &b, &c)
	return NewVersion(a, b, c)
}

func ReadVersion() Version {
	var major, minor, patch uint
	var version string
	flag.UintVar(&major, "major", 0, "Set version major (default 0)")
	flag.UintVar(&minor, "minor", 0, "Set version minor (default 0)")
	flag.UintVar(&patch, "patch", 0, "Set version patch (default 0)")
	flag.StringVar(&version, "version", "", "Set version, expects full version string (e.g. 1.2.3)")
	flag.Parse()

	if (major == 0) && (minor == 0) && (patch == 0) && (version != "") {
		return *ParseVersion(version)
	} else {
		return *NewVersion(major, minor, patch)
	}
}


var (
    hardcoded string
)

func main() {
	// Make a new version object and print simple and json version (use constructor to get the full benefits)
	v := NewVersion(1, 2, 3)
	v.PrintVersion()
	v.PrintJsonVersion()
	
	// Read version from the command line using the helper function
	v2 := ReadVersion()
	v2.PrintJsonVersion()

	// Use 'go build -ldflags "-X main.hardcoded=1.0.1" main.go' to pass the compile-time hardcoded value for v3
	v3 :=  *ParseVersion(hardcoded)
	v3.PrintJsonVersion()
}
