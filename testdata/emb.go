package test

import "embed"

//go:embed * buzz
var File embed.FS
