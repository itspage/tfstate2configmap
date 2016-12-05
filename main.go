package main

import (
	"errors"
	"flag"

	"github.com/itspage/tfstate2configmap/configmap"
	"github.com/itspage/tfstate2configmap/state"
)

func main() {
	encoding := flag.String("o", "json", "Output format. One of: json|yaml")
	flag.Parse()

	if len(flag.Args()) < 1 {
		panic(errors.New("Path of .tfstate must be provided"))
	}

	s, err := state.ReadState(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	if err := configmap.WriteConfigMap(s, *encoding); err != nil {
		panic(err)
	}
}
