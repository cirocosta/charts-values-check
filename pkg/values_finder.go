package pkg

import (
	"reflect"

	"github.com/pkg/errors"
	// see https://github.com/go-yaml/yaml/issues/139
	yaml "github.com/superwhiskers/yaml"
)

type ValuesFinder struct{}

func (p *ValuesFinder) Find(input []byte) (paths []string, err error) {
	parsed := map[string]interface{}{}

	err = yaml.Unmarshal(input, parsed)
	if err != nil {
		err = errors.Wrapf(err,
			"couldn't unmarshal input")
		return
	}

	for k, v := range parsed {
		paths = append(paths, discoverKeyPaths(k, v)...)
	}

	return
}

func discoverKeyPaths(k string, v interface{}) (res []string) {
	if v == nil {
		res = []string{k}
		return
	}

	if reflect.TypeOf(v).Kind() != reflect.Map {
		res = []string{k}
		return
	}

	value := v.(map[interface{}]interface{})

	if len(value) == 0 {
		res = []string{k}
		return
	}

	for innerKey, innerValue := range value {
		paths := discoverKeyPaths(innerKey.(string), innerValue)

		for idx, path := range paths {
			paths[idx] = k + "." + path
		}

		res = append(res, paths...)
	}

	return
}
