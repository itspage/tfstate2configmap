package configmap

import (
	"encoding/json"
	"fmt"

	"github.com/itspage/tfstate2configmap/state"
)

var template = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: "tfstate"
data:
  tfstate: |
    %v
`

func WriteConfigMap(s *state.State) error {
	conf := make(map[string]string)

	for _, m := range s.Modules {
		for k, o := range m.Outputs {
			conf[k] = o.Value
		}
	}

	b, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	fmt.Printf(template, string(b))
	return nil
}
