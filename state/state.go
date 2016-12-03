package state

import (
	"encoding/json"
	"io/ioutil"
)

type Output struct {
	Sensitive bool   `json:"sensitive"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type Module struct {
	Outputs map[string]Output
}

type State struct {
	Modules []Module `json:"modules"`
}

func ReadState(path string) (*State, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s State
	return &s, json.Unmarshal(f, &s)
}
