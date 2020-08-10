package got

import (
	"net/url"
	"path/filepath"
)

func GetFilename(URL string) string {

	u, err := url.Parse(URL)

	if err != nil {
		return "got_output"
	}

	basename := filepath.Base(u.Path)

	if basename == "/" {
		return "index"
	}

	return basename
}