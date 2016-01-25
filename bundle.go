package goci

import (
	"github.com/cloudfoundry-incubator/goci/specs"
)

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
				Version: "0.2.0",
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

// WithNamespace returns a bundle with the given namespace in the list of namespaces. The bundle is not modified, but any
// existing namespace of this type will be replaced.
func (b Bndl) WithNamespace(ns specs.Namespace) *Bndl {
	slice := NamespaceSlice(b.RuntimeSpec.Linux.Namespaces)
	b.RuntimeSpec.Linux.Namespaces = []specs.Namespace(slice.Set(ns))
	return &b
}

func (b Bndl) WithUIDMappings(mappings ...specs.IDMapping) *Bndl {
	b.RuntimeSpec.Linux.UIDMappings = mappings
	return &b
}

func (b Bndl) WithGIDMappings(mappings ...specs.IDMapping) *Bndl {
	b.RuntimeSpec.Linux.GIDMappings = mappings
	return &b
}

func (b Bndl) WithPrestartHooks(hook ...specs.Hook) *Bndl {
	b.RuntimeSpec.Hooks.Prestart = hook
	return &b
}

func (b Bndl) WithPoststopHooks(hook ...specs.Hook) *Bndl {
	b.RuntimeSpec.Hooks.Poststop = hook
	return &b
}

// WithNamespaces returns a bundle with the given namespaces. The original bundle is not modified, but the original
// set of namespaces is replaced in the returned bundle.
func (b Bndl) WithNamespaces(namespaces ...specs.Namespace) *Bndl {
	b.RuntimeSpec.Linux.Namespaces = namespaces
	return &b
}

// WithDevices returns a bundle with the given devices added. The original bundle is not modified.
func (b Bndl) WithDevices(devices ...specs.Device) *Bndl {
	b.RuntimeSpec.Linux.Devices = devices
	return &b
}

// WithCapabilities returns a bundle with the given capabilities added. The original bundle is not modified.
func (b Bndl) WithCapabilities(capabilities ...string) *Bndl {
	b.Spec.Linux.Capabilities = capabilities
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

// GetRootfsPath returns the path to the rootfs of this bundle. Nothing is modified
func (b Bndl) GetRootfsPath() string {
	return b.Spec.Spec.Root.Path
}

type NamespaceSlice []specs.Namespace

func (slice NamespaceSlice) Set(ns specs.Namespace) NamespaceSlice {
	for i, namespace := range slice {
		if namespace.Type == ns.Type {
			slice[i] = ns
			return slice
		}
	}

	return append(slice, ns)
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
