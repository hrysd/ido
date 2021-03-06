package configuration

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const FILENAME = ".ido"

func Read() []byte {
	content, err := ioutil.ReadFile(configurationFilePath())

	if err != nil {
		panic(err) // XXX
	}

	return content
}

func configurationFilePath() string {
	return filepath.Join(os.Getenv("HOME"), FILENAME)
}
