package util

import "os"

// Credit: https://stackoverflow.com/a/22467409/1768931
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
