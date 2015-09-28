package goci

// Bndle represents an in-memory OCI bundle
type Bndle struct {
	Spec        Spec
	RuntimeSpec RuntimeSpec
}

// Bundle creats a bndle
func Bundle() Bndle {
	return Bndle{}
}

// WithProcess sets the bundle's process to the given process
func (b Bndle) WithProcess(process SpecProcess) Bndle {
	b.Spec.Process = process
	return b
}

// WithMounts returns a bundle with the given mounts added
func (b Bndle) WithMounts(mounts ...Mount) Bndle {
	if b.RuntimeSpec.Mounts == nil {
		b.RuntimeSpec.Mounts = make(map[string]RuntimeSpecMount)
	}

	for _, m := range mounts {
		b.Spec.Mounts = append(b.Spec.Mounts, SpecMount{Name: m.Name, Path: m.Destination})
		b.RuntimeSpec.Mounts[m.Name] = RuntimeSpecMount{
			Source:  m.Source,
			Type:    m.Type,
			Options: m.Options,
		}
	}

	return b
}

// Process returns an OCI Process struct with the given args
func Process(args ...string) SpecProcess {
	return SpecProcess{Args: args}
}

// Spec mirrors the github.com/opencontainers/specs.Spec struct but allows it to compile
// on non-linux systems.
type Spec struct {
	Version string      `json:"version"`
	Mounts  []SpecMount `json:"mounts"`
	Process SpecProcess `json:"process"`
}

// RuntimeSpec mirrors the github.com/opencontainers/specs.RuntimeSpec struct but allows it to compile
// on non-linux systems.
type RuntimeSpec struct {
	Mounts map[string]RuntimeSpecMount `json:"mounts"`
}

// Process mirrors the "github.com/opencontainers/specs".Process struct
type SpecProcess struct {
	Args []string `json:"args"`
}

type SpecMount struct {
	Name string
	Path string
}

type RuntimeSpecMount struct {
	Type    string
	Source  string
	Options []string
}

type Mount struct {
	Name        string
	Source      string
	Destination string
	Type        string
	Options     []string
}
