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

var (
	NetworkNamespace = specs.Namespace{Type: specs.NetworkNamespace}
	UserNamespace    = specs.Namespace{Type: specs.UserNamespace}
	PIDNamespace     = specs.Namespace{Type: specs.PIDNamespace}
	IPCNamespace     = specs.Namespace{Type: specs.IPCNamespace}
	UTSNamespace     = specs.Namespace{Type: specs.UTSNamespace}
	MountNamespace   = specs.Namespace{Type: specs.MountNamespace}
)

// WithProcess returns a bundle with the process replaced with the given Process. The original bundle is not modified.
func (b Bndl) WithProcess(process specs.Process) *Bndl {
	b.Spec.Process = process
	return &b
}

func (b Bndl) WithRootFS(absolutePath string) *Bndl {
	b.Spec.Root = specs.Root{Path: absolutePath}
	return &b
}

// WithResources returns a bundle with the resources replaced with the given Resources. The original bundle is not modified.
func (b Bndl) WithResources(resources *specs.Resources) *Bndl {
	b.RuntimeSpec.Linux.Resources = resources
	return &b
}

// WithNamespaces returns a bundle with the given namespaces added. The original bundle is not modified.
func (b Bndl) WithNamespaces(namespaces ...specs.Namespace) *Bndl {
	for _, namespace := range namespaces {
		b.RuntimeSpec.Linux.Namespaces = append(b.RuntimeSpec.Linux.Namespaces, namespace)
	}

	return &b
}

// WithDevices returns a bundle with the given devices added. The original bundle is not modified.
func (b Bndl) WithDevices(devices ...specs.Device) *Bndl {
	for _, device := range devices {
		b.RuntimeSpec.Linux.Devices = append(b.RuntimeSpec.Linux.Devices, device)
	}

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
