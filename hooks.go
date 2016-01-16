package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
    "fmt"
    "sort"
)

const CONFIGURATION_FILENAME = ".ido"

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

    var hooks []Hook
	errr := json.Unmarshal(fileContent, &hooks)
    
    if errr != nil {
        fmt.Println(errr)
        return
    }
    
    fmt.Println(hooks)
    
    i := sort.Search(len(hooks), func(i int) bool {
        return hooks[i].Name == "hrysd"
    })
    
    fmt.Println(hooks[i])
}
func configurationPath() string {
	return filepath.Join(os.Getenv("HOME"), CONFIGURATION_FILENAME)
}
