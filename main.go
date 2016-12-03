package main

import (
	"errors"
	"os"

	"github.com/itspage/tfstate2configmap/configmap"
	"github.com/itspage/tfstate2configmap/state"
)

func main() {
	if len(os.Args) <= 1 {
		panic(errors.New("Path of .tfstate must be provided"))
	}

	s, err := state.ReadState(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := configmap.WriteConfigMap(s); err != nil {
		panic(err)
	}

}
