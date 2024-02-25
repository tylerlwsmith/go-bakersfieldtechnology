//go:build !noembed

package assets

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
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
		return entry.File, nil
	} else {
		return "", fmt.Errorf("assets.resolve: asset %v not found", src)
	}
}

func init() {
	f, err := AssetFiles.ReadFile("dist/manifest.json")
	if err != nil {
		log.Fatalf("assets.init %v", err)
	}

	err = json.Unmarshal(f, &manifest)
	if err != nil {
		log.Fatalf("assets.init %v", err)
	}
}

func Get(src string) (resolved string, err error) {
	resolved, err = manifest.resolve(src)

	return "/assets/" + resolved, err
}

func AssetHandler() http.Handler {
	log.Print("using embed mode")
	fsys, err := fs.Sub(AssetFiles, "dist/public")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(fsys))
}
