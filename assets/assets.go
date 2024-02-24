//go:build !noembed

package assets

import "embed"

//go:embed dist
var AssetFiles embed.FS

//go:embed public
var PublicFiles embed.FS
