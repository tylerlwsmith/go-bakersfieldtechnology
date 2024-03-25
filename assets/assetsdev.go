//go:build noembed

package assets

import (
	"io/fs"
	"log"
	"net/http"
	"os"
)

var AssetsHandler http.Handler
var PublicFiles fs.FS
var viteOrigin string = "http://localhost:5173"

func init() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	PublicFiles = os.DirFS(dir + "/assets/public")

	AssetsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	viteOriginEnv := os.Getenv("VITE_ORIGIN")
	if viteOriginEnv != "" {
		viteOrigin = viteOriginEnv
	}

	if viteOrigin[len(viteOrigin)-1:] != "/" {
		viteOrigin = viteOrigin + "/"
	}
}

func Get(src string) (resolved string, err error) {
	return viteOrigin + "assets/" + src, err
}
