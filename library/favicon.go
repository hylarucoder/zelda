package library

import (
	_ "github.com/thinkerou/favicon"
)

fav_middleware := favicon.New("./favicon.ico")
