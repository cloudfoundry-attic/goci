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

	Describe("SetProcess", func() {
		It("adds the process to the bundle", func() {
			returnedBundle := initialBundle.SetProcess(goci.Process("echo", "foo"))
			Expect(returnedBundle.Spec.Process).To(Equal(specs.Process{Args: []string{"echo", "foo"}}))
		})

		It("returns the modified bundle", func() {
			returnedBundle := initialBundle.SetProcess(goci.Process("echo", "foo"))
			Expect(returnedBundle).To(Equal(initialBundle))
		})
	})

	Describe("AddMounts", func() {
		BeforeEach(func() {
			returnedBundle = initialBundle.AddMounts(
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
			Expect(initialBundle.Spec.Mounts).To(ContainElement(specs.MountPoint{Name: "banana", Path: "/banana"}))
			Expect(initialBundle.Spec.Mounts).To(ContainElement(specs.MountPoint{Name: "apple", Path: "/apple"}))
		})

		It("returns a bundle with the mounts mapped in the runtime spec", func() {
			Expect(initialBundle.RuntimeSpec.Mounts).To(HaveKeyWithValue("banana", specs.Mount{
				Type:    "banana_fs",
				Source:  "banana_device",
				Options: []string{"yellow", "fresh"},
			}))

			Expect(initialBundle.RuntimeSpec.Mounts).To(HaveKeyWithValue("apple", specs.Mount{
				Type:    "apple_fs",
				Source:  "iDevice",
				Options: []string{"healthy", "shiny"},
			}))
		})

		It("returns the modified bundle", func() {
			Expect(returnedBundle).To(Equal(initialBundle))
		})
	})
})
