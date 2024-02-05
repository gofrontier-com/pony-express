package adf

import (
	json "github.com/json-iterator/go"
)

type Change struct {
	Name  string
	Type  string
	Patch string
}

func (c *Change) addJsonPatch(patch interface{}) error {
	js, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	c.Patch = string(js)
	return nil
}
