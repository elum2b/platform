// Package static provides the embedded SolidJS dashboard bundle.
package static

import (
	"embed"
	"io/fs"
)

// embeddedFiles contains the dashboard build output embedded into the executable.
// The frontend build output is placed in the files directory before release.
//
//go:embed all:files
var embeddedFiles embed.FS

// Files contains the dashboard build output without the files directory prefix.
var Files = embeddedFS()

func embeddedFS() fs.FS {
	files, err := fs.Sub(embeddedFiles, "files")
	if err != nil {
		panic(err)
	}

	return files
}
