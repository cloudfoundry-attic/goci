A Golang API for working with OCI containers. Go OCI. GOCI. 

# THIS LIBRARY HAS MERGED INTO GUARDIAN AND WILL NO LONGER BE MAINTAINED AT THIS LOCATION

It can be found at in the [guardian repository](https://github.com/cloudfoundry-incubator/guardian/tree/master/rundmc/goci)

# API

Here's what we're going for:

## Bundle Creation Helpers

To create an OCI bundle, create a `goci.Bndle` and issue `Save()`.

~~~~
// struct combining the various required spec structures
bundle := goci.Bndle{
  RootfsPath: "some-path", // absolute path, will become the 'rootfs' subdirectory of saved bundle
  Spec: &specs.Spec{},
  RuntimeSpec: &specs.Spec{},
}

// some helpers to work with the bundle fields
pathToBundle, err := goci.Bundle("/bin/echo", "foo", "bar")
          .WithRootfs("someRootfsPath")
          .WithMounts(mounts.New("proc", "/proc")) // added to both spec and runtime spec for you
          .WithUserNamespace("", []bundle.UidRange{})
          .WithMountNamespace()
          .Save("") // empty string to save to temporary directory, non-empty to specify path
~~~~

## Container Helpers

~~~~
cmd := goci.StartCommand(pathToBundle, id)
cmd.Run() // just a regular exec.Cmd..
cmd.Start()

// goci also supports some non-standard runC features
cmd := goci.ExecCommand(id, &specs.Process{ User: "foo", Args: []string{ "echo", "hi" }})
cmd.Run() // just a regular exec.Cmd..
~~~~
