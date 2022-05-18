package g

import (
	"net/url"
	"path/filepath"
)

func Resolve(from string, to string) string {
	if filepath.IsAbs(to) {
		return to
	}

	if filepath.IsAbs(from) {
		return filepath.Join(from, to)
	}

	abs, err := filepath.Abs(from)

	if err != nil {
		abs = "/"
	}

	return filepath.Join(abs, to)
}

func ResolveURL(from string, to string) (string, error) {
	u, e := url.Parse(from)
	if e != nil {
		return "", e
	}

	if u.Path == "" {
		u.Path = Resolve("/", to)
	} else {
		u.Path = Resolve(u.Path, to)
	}

	return u.String(), nil
}
