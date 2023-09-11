package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type settings struct {
	includeRegex string
	excludeRegex string
	localPath    string
	siteId       string
}

type project struct {
	ModelVersion string `xml:"modelVersion"`
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	Packaging    string `xml:"packaging"`
	Name         string `xml:"name"`
}

func main() {

	settings := settings{localPath: "C:\\Users\\Anders\\.m2\\repository"}
	arguments := os.Args[1:]
	num := len(arguments)

	for i := 0; i < num; i++ {

		switch arguments[i] {
		case "--include":
			{

				if i+1 < num {
					settings.includeRegex = arguments[i+1]
					i++
					continue
				} else {
					panic("Inclusion flag not followed by inclusion regex argument")
				}

			}

		case "--exclude":
			{

				if i+1 < num {
					settings.excludeRegex = arguments[i+1]
					i++
					continue
				} else {
					panic("Exclusion flag not followed by exclusion regex argument")
				}

			}

		case "--local-repo":
			{

				if i+1 < num {
					settings.localPath = arguments[i+1]
					i++
					continue
				} else {
					panic("Exclusion flag not followed by exclusion regex argument")
				}

			}

		default:
			{
				if i+1 == num {
					settings.siteId = arguments[i]
				} else {
					panic(fmt.Sprintf("Unknown argument: %v", arguments[i]))
				}
			}

		}
	}

	if settings.siteId == "" {
		panic("Missing siteId argument - must be last argument - must be defined in your maven ~/.m2/settings.xml file )")
	}

	// All set
	fmt.Println(settings)
	home, _ := os.UserHomeDir()
	repo := strings.Join([]string{home, ".m2", "repository"}, string(os.PathSeparator))
	var poms []string

	if runtime.GOOS == "windows" {
		fmt.Println("[ REPO ] :", repo)
		winPath := filepath.FromSlash(repo)
		tree(winPath, &poms)
	} else {
		tree(settings.localPath, &poms)
	}

	// fmt.Println(poms)

	for _, pom := range poms {
		text, error := os.ReadFile(pom)

		if error != nil {
			fmt.Println(error)
		}

		// fmt.Printf("%s", text)

		var proj project = project{Packaging: "jar"}

		_ = xml.Unmarshal(text, &proj)

		fmt.Println("[ PARSED ] :", proj)
	}

}

func tree(path string, poms *[]string) {

	entries, error := os.ReadDir(path)

	if error != nil {
		panic(error)
	}

	for _, entry := range entries {
		path := strings.Join([]string{path, entry.Name()}, string(os.PathSeparator))

		if entry.IsDir() {
			tree(path, poms)
		} else {
			lowerName := strings.ToLower(entry.Name())
			if strings.HasSuffix(lowerName, ".pom") {
				*poms = append(*poms, path)
			}
		}

	}

}
