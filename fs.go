package g

import (
	"io"
	"os"
	"path"
)

func Exists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CopyFile copies a file from source to target.
func CopyFile(source, target string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

// SyncFile copies a file from source to target.
// If target is in a not exists directory,
// SyncFile will create the directory first.
// If target is exists, SyncFile can overwrite it or not.
func SyncFile(source, target string, overwrite bool) error {
	err := createDirIfNotExists(target)
	if err != nil {
		return err
	}

	exists, err := Exists(target)
	if err != nil {
		return err
	}

	if !exists {
		return CopyFile(source, target)
	}

	if overwrite {
		os.Remove(target)
		return CopyFile(source, target)
	}

	return nil
}

func createDirIfNotExists(filename string) error {
	d := path.Dir(filename)
	if d == "" {
		return nil
	}

	exists, err := Exists(d)
	if err != nil {
		return err
	}
	if !exists {
		err := os.MkdirAll(d, os.ModePerm)
		return err
	}

	return nil
}
