package goci

import "github.com/cloudfoundry-incubator/goci/specs"

// Bndl represents an in-memory OCI bundle
type Bndl struct {
	Spec        specs.LinuxSpec
	RuntimeSpec specs.LinuxRuntimeSpec
}

// Bundle creates a Bndl
func Bundle() *Bndl {
	return &Bndl{
		Spec: specs.LinuxSpec{
			Spec: specs.Spec{
				Version: "0.1.0",
			},
		},
	}
}

// WithProcess returns a bundle with the process replaced with the given Process. The original bundle is not modified.
func (b Bndl) WithProcess(process specs.Process) *Bndl {
	b.Spec.Process = process
	return &b
}

// WithResources returns a bundle with the resources replaced with the given Resources. The original bundle is not modified.
func (b Bndl) WithResources(resources *specs.Resources) *Bndl {
	b.RuntimeSpec.Linux.Resources = resources
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
