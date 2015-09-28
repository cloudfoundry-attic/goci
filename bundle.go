package goci

import "github.com/cloudfoundry-incubator/goci/specs"

// Bndl represents an in-memory OCI bundle
type Bndl struct {
	Spec        specs.Spec
	RuntimeSpec specs.LinuxRuntimeSpec
}

// Bundle creates a Bndl
func Bundle() *Bndl {
	return &Bndl{}
}

// SetProcess sets the bundle's process to the given process and returns the modified bundle.
func (b *Bndl) SetProcess(process specs.Process) *Bndl {
	b.Spec.Process = process
	return b
}

// AddMounts adds the given mounts to the existing mounts and returns the modified bundle.
func (b *Bndl) AddMounts(mounts ...Mount) *Bndl {
	if b.RuntimeSpec.Mounts == nil {
		b.RuntimeSpec.Mounts = make(map[string]specs.Mount)
	}

	for _, m := range mounts {
		b.Spec.Mounts = append(b.Spec.Mounts, specs.MountPoint{Name: m.Name, Path: m.Destination})
		b.RuntimeSpec.Mounts[m.Name] = specs.Mount{
			Source:  m.Source,
			Type:    m.Type,
			Options: m.Options,
		}
	}

	return b
}

// Process returns an OCI Process struct with the given args.
func Process(args ...string) specs.Process {
	return specs.Process{Args: args}
}

type Mount struct {
	Name        string
	Source      string
	Destination string
	Type        string
	Options     []string
}
