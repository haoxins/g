package g

import "path/filepath"

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
