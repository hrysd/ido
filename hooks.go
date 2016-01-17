package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const CONFIGURATION_FILENAME = ".ido"

type Hook struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Hooks []Hook

func DetectHook(hookName string) Hook {
	hooks := LoadHooks()
    hook := hooks.Detect(hookName)
    
    return hook
}

func LoadHooks() Hooks {
    var hooks Hooks 
    err := json.Unmarshal(readHookData(), &hooks)
    
    if err != nil {
        panic(err)
    }

    return hooks
}

func (self Hooks)Detect(hookName string) Hook {
    var hook Hook
    
    for _, h := range self {
		if h.Name == hookName {
			hook = h
		}
	}
    
    return hook
}

func readHookData() []byte {
    fileContent, err := ioutil.ReadFile(configurationPath())
    
    if err != nil {
        panic(err)
    }
    
    return fileContent
}

func configurationPath() string {
	return filepath.Join(os.Getenv("HOME"), CONFIGURATION_FILENAME)
}
