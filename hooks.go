package main

import (
	"encoding/json"

	"github.com/hrysd/ido/internal/utils"
)

type Hook struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Hooks []Hook

func (self Hooks) DetectByName(hookName string) Hook {
	var hook Hook

	for _, h := range self {
		if h.Name == hookName {
			hook = h
		}
	}

	if hook.Name == "" {
		panic("Can not detect hook")
	}

	return hook
}

func DetectHook(hookName string) Hook {
	hooks := loadHooks()
	hook := hooks.DetectByName(hookName)

	return hook
}

func loadHooks() Hooks {
	var hooks Hooks
	err := json.Unmarshal(configuration.Read(), &hooks)

	if err != nil {
		panic(err) // XXX
	}

	return hooks
}
