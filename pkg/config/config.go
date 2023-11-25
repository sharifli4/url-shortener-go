package config

import "os"

var (
	TinyDomain = os.Getenv("TINY_DOMAIN")
)
