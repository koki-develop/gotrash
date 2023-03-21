package util

import (
	"os"
)

func CreateDir(dir string) error {
	exists, err := Exists(dir)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func Exists(p string) (bool, error) {
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
