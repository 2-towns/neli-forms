package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

var (
	Port      = flag.Int("port", 8080, "Port for running application")
	VideosURL = flag.String("videos_url", "http://localhost:8080/neli", "Url to video content")
	ApiURL    = flag.String("api_url", "http://localhost:8082", "Url to api")
)

func init() {
	iniflags.Parse()
}
