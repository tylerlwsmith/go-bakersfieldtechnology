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
var assetFiles embed.FS

//go:embed public
var publicFiles embed.FS
var PublicFiles fs.FS

var AssetsHandler http.Handler

type manifestEntry struct {
	File    string `json:"file"`
	IsEntry bool   `json:"isEntry"`
	Src     string `json:"src"`
}

var manifest map[string]manifestEntry

func init() {
	f, err := assetFiles.ReadFile("dist/manifest.json")
	if err != nil {
		log.Fatalf("assets.init %v", err)
	}

	err = json.Unmarshal(f, &manifest)
	if err != nil {
		log.Fatalf("assets.init %v", err)
	}

	log.Print("using embed mode")
	assetFsys, err := fs.Sub(assetFiles, "dist/public")
	if err != nil {
		panic(err)
	}

	AssetsHandler = http.FileServer(http.FS(assetFsys))

	PublicFiles, err = fs.Sub(publicFiles, "public")
	if err != nil {
		panic(err)
	}
}

func Get(src string) (resolved string, err error) {
	if entry, ok := manifest[src]; ok {
		return "/assets/" + entry.File, nil
	} else {
		return "", fmt.Errorf("assets.resolve: asset %v not found", src)
	}
}
