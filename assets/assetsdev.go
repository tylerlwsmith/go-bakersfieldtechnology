//go:build noembed

package assets

import (
	"io/fs"
	"log"
	"os"
)

var TmplFiles fs.FS

func init() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	PublicFiles = os.DirFS(dir + "/public")
}
