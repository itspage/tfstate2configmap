package configmap

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-yaml/yaml"

	"github.com/itspage/tfstate2configmap/state"
)

var template = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: "tfstate"
data:
  %v
`

func WriteConfigMap(s *state.State, encoding string) error {
	conf := make(map[string]string)

	for _, m := range s.Modules {
		for k, o := range m.Outputs {
			conf[k] = o.Value
		}
	}

	var b []byte
	var err error

	switch encoding {
	case "yaml":
		b, err = yaml.Marshal(conf)
	case "json":
		b, err = json.Marshal(conf)
	default:
		return errors.New("unknown encoding")
	}

	if err != nil {
		return err
	}

	fmt.Printf(template, string(b))
	return nil
}
