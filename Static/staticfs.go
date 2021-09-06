package Static

import (
	"embed"
)

//go:embed resource/*
var resource embed.FS

func Resource() embed.FS {
	return resource
}
