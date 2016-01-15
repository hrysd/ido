package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
    "fmt"
)

const CONFIGURATION_FILENAME = ".ido"

type Hooks struct {
    Hooks []Hook
}

type Hook struct {
    Name string `json:"name"`
    Token string `json:"token"`
}

func main() {
	LoadHooks()
}

func LoadHooks() {
	fileContent, err := ioutil.ReadFile(configurationPath())

	if err != nil {
		fmt.Println(err)
		return
	}

    hooks := Hooks{}
	errr := json.Unmarshal(fileContent, &hooks)
    
    if errr != nil {
        fmt.Println(errr)
        return
    }
    
    fmt.Println(hooks)
}
func configurationPath() string {
	return filepath.Join(os.Getenv("HOME"), CONFIGURATION_FILENAME)
}
