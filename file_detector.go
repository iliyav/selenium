package selenium

import "os"

type FileDetectorFunc func(path string) bool

func GetLocalFileDummy(path string) bool {
	return false
}

func GetLocalFile(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}
