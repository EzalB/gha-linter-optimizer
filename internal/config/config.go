package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Path     string
	Repo     string
	PRNumber int
}

func Load() *Config {
	var cfg Config

	flag.StringVar(&cfg.Path, "path", ".github/workflows", "Path to workflows")
	flag.StringVar(&cfg.Repo, "repo", os.Getenv("GITHUB_REPOSITORY"), "GitHub repo (e.g., user/repo)")
	pr := os.Getenv("PR_NUMBER")
	if pr == "" {
		pr = "0"
	}
	cfg.PRNumber, _ = strconv.Atoi(pr)

	flag.Parse()
	return &cfg
}

func (c *Config) IsPRContext() bool {
	return c.PRNumber > 0 && c.Repo != ""
}
