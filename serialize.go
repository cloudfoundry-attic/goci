package goci

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (b *Bndle) Save(path string) error {
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
