//go:build !noembed

package assets

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

//go:embed dist
var AssetFiles embed.FS

//go:embed public
var PublicFiles embed.FS

type manifestEntry struct {
	File    string `json:"file"`
	IsEntry bool   `json:"isEntry"`
	Src     string `json:"src"`
}

type assetManifest map[string]manifestEntry

var manifest assetManifest

func (m assetManifest) resolve(src string) (resolved string, err error) {
	if entry, ok := m[src]; ok {
		return entry.Src, nil
	} else {
		return "", fmt.Errorf("asset %v not found", src)
	}
}

func init() {
	f, err := AssetFiles.Open("manifest.json")
	if err != nil {
		log.Fatal("unable to open asset manifest.")
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("unable to open asset manifest.")
	}

	err = json.Unmarshal(b, &manifest)
	if err != nil {
		log.Fatal("unable to unmarshal manifest")
	}
}

func Get(src string) (resolved string, err error) {
	resolved, err = manifest.resolve(src)
	return resolved, err
}
