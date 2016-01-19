package goci_test

import (
	"github.com/cloudfoundry-incubator/goci"
	"github.com/cloudfoundry-incubator/goci/specs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bundle", func() {
	var initialBundle *goci.Bndl
	var returnedBundle *goci.Bndl

	BeforeEach(func() {
		initialBundle = goci.Bundle()
	})

	It("specifies the correct version", func() {
		Expect(initialBundle.Spec.Version).To(Equal("0.2.0"))
	})

	Describe("WithCapabilities", func() {
		It("adds capabilities to the bundle", func() {
			returnedBundle := initialBundle.WithCapabilities("growtulips", "waterspuds")
			Expect(returnedBundle.Spec.Linux.Capabilities).To(ContainElement("growtulips"))
			Expect(returnedBundle.Spec.Linux.Capabilities).To(ContainElement("waterspuds"))
		})

		It("does not modify the initial bundle", func() {
			returnedBundle := initialBundle.WithCapabilities("growtulips", "waterspuds")
			Expect(returnedBundle).NotTo(Equal(initialBundle))
		})
	})

	Describe("WithProcess", func() {
		It("adds the process to the bundle", func() {
			returnedBundle := initialBundle.WithProcess(goci.Process("echo", "foo"))
			Expect(returnedBundle.Spec.Process).To(Equal(specs.Process{Args: []string{"echo", "foo"}}))
		})

		It("does not modify the initial bundle", func() {
			returnedBundle := initialBundle.WithProcess(goci.Process("echo", "foo"))
			Expect(returnedBundle).NotTo(Equal(initialBundle))
		})
	})

	Describe("WithRootFS", func() {
		It("sets the rootfs path in the spec", func() {
			returnedBundle := initialBundle.WithRootFS("/foo/bar/baz")
			Expect(returnedBundle.Spec.Root.Path).To(Equal("/foo/bar/baz"))
		})
	})

	Describe("WithPrestartHooks", func() {
		It("adds the hook to the runtime spec", func() {
			returnedBundle := initialBundle.WithPrestartHooks(specs.Hook{
				Path: "foo",
				Args: []string{"bar", "baz"},
			})

			Expect(returnedBundle.RuntimeSpec.Hooks.Prestart).To(Equal([]specs.Hook{{
				Path: "foo",
				Args: []string{"bar", "baz"},
			}}))
		})
	})

	Describe("WithMounts", func() {
		BeforeEach(func() {
			returnedBundle = initialBundle.WithMounts(
				goci.Mount{
					Name:        "apple",
					Type:        "apple_fs",
					Source:      "iDevice",
					Destination: "/apple",
					Options: []string{
						"healthy",
						"shiny",
					},
				},
				goci.Mount{
					Name:        "banana",
					Type:        "banana_fs",
					Source:      "banana_device",
					Destination: "/banana",
					Options: []string{
						"yellow",
						"fresh",
					},
				})
		})

		It("returns a bundle with the mounts added to the spec", func() {
			Expect(returnedBundle.Spec.Mounts).To(ContainElement(specs.MountPoint{Name: "banana", Path: "/banana"}))
			Expect(returnedBundle.Spec.Mounts).To(ContainElement(specs.MountPoint{Name: "apple", Path: "/apple"}))
		})

		It("returns a bundle with the mounts mapped in the runtime spec", func() {
			Expect(returnedBundle.RuntimeSpec.Mounts).To(HaveKeyWithValue("banana", specs.Mount{
				Type:    "banana_fs",
				Source:  "banana_device",
				Options: []string{"yellow", "fresh"},
			}))

			Expect(returnedBundle.RuntimeSpec.Mounts).To(HaveKeyWithValue("apple", specs.Mount{
				Type:    "apple_fs",
				Source:  "iDevice",
				Options: []string{"healthy", "shiny"},
			}))
		})

		It("does not modify the original bundle", func() {
			Expect(returnedBundle).NotTo(Equal(initialBundle))
			Expect(initialBundle.Spec.Mounts).To(HaveLen(0))
		})
	})

	Describe("WithResources", func() {
		BeforeEach(func() {
			returnedBundle = initialBundle.WithResources(&specs.Resources{DisableOOMKiller: true})
		})

		It("returns a bundle with the resources added to the runtime spec", func() {
			Expect(returnedBundle.RuntimeSpec.Linux.Resources).To(Equal(&specs.Resources{DisableOOMKiller: true}))
		})

		It("does not modify the original bundle", func() {
			Expect(returnedBundle).NotTo(Equal(initialBundle))
			Expect(initialBundle.RuntimeSpec.Linux.Resources).To(BeNil())
		})
	})

	Describe("WithNamespace", func() {
		It("does not change any namespaces other than the one with the given type", func() {
			colin := specs.Namespace{Type: "colin", Path: ""}
			potato := specs.Namespace{Type: "potato", Path: "pan"}

			initialBundle = initialBundle.WithNamespace(colin)
			returnedBundle = initialBundle.WithNamespace(potato)
			Expect(returnedBundle.RuntimeSpec.Linux.Namespaces).To(ConsistOf(colin, potato))
		})

		Context("when the namespace isnt already in the spec", func() {
			It("adds the namespace", func() {
				ns := specs.Namespace{Type: "potato", Path: "pan"}
				returnedBundle = initialBundle.WithNamespace(ns)
				Expect(returnedBundle.RuntimeSpec.Linux.Namespaces).To(ConsistOf(ns))
			})
		})

		Context("when the namespace is already in the spec", func() {
			It("overrides the namespace", func() {
				initialBundle = initialBundle.WithNamespace(specs.Namespace{Type: "potato", Path: "should-be-overridden"})
				ns := specs.Namespace{Type: "potato", Path: "pan"}
				returnedBundle = initialBundle.WithNamespace(ns)
				Expect(returnedBundle.RuntimeSpec.Linux.Namespaces).To(ConsistOf(ns))
			})
		})
	})

	Describe("WithNamespaces", func() {
		BeforeEach(func() {
			returnedBundle = initialBundle.WithNamespaces(specs.Namespace{Type: specs.NetworkNamespace})
		})

		It("returns a bundle with the namespaces added to the runtime spec", func() {
			Expect(returnedBundle.RuntimeSpec.Linux.Namespaces).To(ConsistOf(specs.Namespace{Type: specs.NetworkNamespace}))
		})

		Context("when the spec already contains namespaces", func() {
			It("replaces them", func() {
				overridenBundle := returnedBundle.WithNamespaces(specs.Namespace{Type: "mynamespace"})
				Expect(overridenBundle.RuntimeSpec.Linux.Namespaces).To(ConsistOf(specs.Namespace{Type: "mynamespace"}))
			})
		})
	})

	Describe("WithUIDMappings", func() {
		It("returns a bundle with the provided uid mappings added to the runtime spec", func() {
			uidMappings := []specs.IDMapping{
				{
					HostID:      40000,
					ContainerID: 0,
					Size:        1,
				},
				{
					HostID:      1,
					ContainerID: 1,
					Size:        39999,
				},
			}
			returnedBundle := initialBundle.WithUIDMappings(uidMappings...)

			Expect(returnedBundle.RuntimeSpec.Linux.UIDMappings).To(Equal(uidMappings))
		})
	})

	Describe("WithGIDMappings", func() {
		It("returns a bundle with the provided gid mappings added to the runtime spec", func() {
			gidMappings := []specs.IDMapping{
				{
					HostID:      40000,
					ContainerID: 0,
					Size:        1,
				},
				{
					HostID:      1,
					ContainerID: 1,
					Size:        39999,
				},
			}
			returnedBundle := initialBundle.WithGIDMappings(gidMappings...)

			Expect(returnedBundle.RuntimeSpec.Linux.GIDMappings).To(Equal(gidMappings))
		})
	})

	Describe("WithDevices", func() {
		BeforeEach(func() {
			returnedBundle = initialBundle.WithDevices(specs.Device{Path: "test/path"})
		})

		It("returns a bundle with the namespaces added to the runtime spec", func() {
			Expect(returnedBundle.RuntimeSpec.Linux.Devices).To(ConsistOf(specs.Device{Path: "test/path"}))
		})

		Context("when the spec already contains namespaces", func() {
			It("replaces them", func() {
				overridenBundle := returnedBundle.WithDevices(specs.Device{Path: "new-device"})
				Expect(overridenBundle.RuntimeSpec.Linux.Devices).To(ConsistOf(specs.Device{Path: "new-device"}))
			})
		})
	})

	Describe("GetRootfsPath", func() {
		It("returns the rootfs path", func() {
			bndl := &goci.Bndl{}
			bndl.Spec.Spec.Root.Path = "/path/to/mordor"
			Expect(bndl.GetRootfsPath()).To(Equal("/path/to/mordor"))
		})
	})

	Describe("NamespaceSlice", func() {
		Context("when the namespace isnt already in the slice", func() {
			It("adds the namespace", func() {
				ns := specs.Namespace{Type: "potato", Path: "pan"}
				nsl := goci.NamespaceSlice{}
				nsl = nsl.Set(ns)
				Expect(nsl).To(ConsistOf(ns))
			})
		})

		Context("when the namespace is already in the slice", func() {
			It("overrides the namespace", func() {
				ns := specs.Namespace{Type: "potato", Path: "pan"}
				nsl := goci.NamespaceSlice{specs.Namespace{Type: "potato", Path: "chips"}}
				nsl = nsl.Set(ns)
				Expect(nsl).To(ConsistOf(ns))
			})
		})
	})
})
