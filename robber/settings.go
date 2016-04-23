package robber

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ToolsSetting struct for map yaml settings.
type ToolsSetting struct {
	Tools []Tool
}

// Set contain settings related to
// the monitoring means.
var Set ToolsSetting

// InitSet load settings from file to Set.
func InitSet(path string) {
	if path == "" {
		path = "robber.yaml"
	}

	t, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(t, &Set); err != nil {
		panic(err)
	}

	for _, v := range Set.Tools {
		if _, err := os.Stat(v.Path); os.IsNotExist(err) {
			log.Fatalf("Bin file: %v not exist.", v.Path)
		}
	}
}
