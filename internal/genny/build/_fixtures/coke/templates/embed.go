package templates

import (
	"embed"
	"io/fs"
)

//go:embed * */*
var files embed.FS

func FS() fs.FS {
	return files
}
