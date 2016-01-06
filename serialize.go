package goci

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type BndlLoader struct {
}

func (b *BndlLoader) Load(path string) (*Bndl, error) {
	bundle := Bndl{}

	err := readJsonInto(filepath.Join(path, "runtime.json"), &bundle.RuntimeSpec)
	if err != nil {
		return nil, err
	}
	err = readJsonInto(filepath.Join(path, "config.json"), &bundle.Spec)
	if err != nil {
		return nil, err
	}

	return &bundle, nil
}

func (b *Bndl) Save(path string) error {
	if err := save(b.Spec, filepath.Join(path, "config.json")); err != nil {
		return err
	}

	return save(b.RuntimeSpec, filepath.Join(path, "runtime.json"))
}

func save(value interface{}, path string) error {
	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(value)
}

func readJsonInto(path string, object interface{}) error {
	runtimeContents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	json.Unmarshal(runtimeContents, object)
	return nil
}
