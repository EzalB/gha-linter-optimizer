package config

import (
	"flag"
)

type Config struct {
	Path string
}

func Load() Config {
	path := flag.String("path", ".github/workflows", "Path to scan for workflow files")
	flag.Parse()

	return Config{Path: *path}
}
