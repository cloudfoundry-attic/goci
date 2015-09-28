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

// WithProcess returns a bundle with the process replaced with the given Process. The original bundle is not modified.
func (b Bndl) WithProcess(process specs.Process) *Bndl {
	b.Spec.Process = process
	return &b
}

// WithMounts returns a bundle with the given mounts added. The original bundle is not modified.
func (b Bndl) WithMounts(mounts ...Mount) *Bndl {
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

	return &b
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
