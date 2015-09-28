A Golang API for working with OCI containers. Go OCI. GOCI. 

# Work in progress

NOTE: this is a work in progress golang wrapper of a work in progress reference implementation
of a work in progress specification. Probably don't use in production k?

# API

Here's what we're going for:

## Bundle Creation Helpers

~~~~
// struct combining the various required spec structurse
bundle := goci.Bndle{
  RootfsPath: "some-path", // hard-linked inside the bundle directory on save
  Spec: &specs.Spec,
  RuntimeSpec: &specs.Spec,
}

// some helpers to work with the bundle fields
bundle := goci.Bundle("/bin/echo", "foo", "bar")
bundle = bundle.WithRootfs("someRootfsPath")
bundle = bundle.WithMounts(mounts.New("proc", "/proc")) // added to both spec and runtime spec for you
bundle = bundle.WithUserNamespace("", []bundle.UidRange{})
bundle = bundle.WithMountNamespace()
pathToBundle, err  := bundle.Save("")
~~~~

## Container Helpers

~~~~
cmd := goci.StartCommand(pathToBundle, id)
cmd.Run() // just a regular exec.Cmd..
cmd.Start()

cmd := goci.ExecCommand(id, &specs.Process{ User: "foo", Args: []string{ "echo", "hi" }})
cmd.Run() // just a regular exec.Cmd..
~~~~
