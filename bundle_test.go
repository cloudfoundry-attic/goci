package goci_test

import (
	"github.com/cloudfoundry-incubator/goci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bundle", func() {
	var (
		initialBundle goci.Bndle
	)

	BeforeEach(func() {
		initialBundle = goci.Bundle()
	})

	Describe("WithProcess", func() {
		It("adds the process to the bundle", func() {
			bundle := goci.Bundle().WithProcess(goci.Process("echo", "foo"))
			Expect(bundle.Spec.Process).To(Equal(goci.SpecProcess{Args: []string{"echo", "foo"}}))
		})
	})

	Describe("WithMounts", func() {
		var bundle goci.Bndle

		BeforeEach(func() {
			bundle = goci.Bundle().WithMounts(
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
			Expect(bundle.Spec.Mounts).To(ContainElement(goci.SpecMount{Name: "banana", Path: "/banana"}))
			Expect(bundle.Spec.Mounts).To(ContainElement(goci.SpecMount{Name: "apple", Path: "/apple"}))
		})

		It("returns a bundle with the mounts mapped in the runtime spec", func() {
			Expect(bundle.RuntimeSpec.Mounts).To(HaveKeyWithValue("banana", goci.RuntimeSpecMount{
				Type:    "banana_fs",
				Source:  "banana_device",
				Options: []string{"yellow", "fresh"},
			}))

			Expect(bundle.RuntimeSpec.Mounts).To(HaveKeyWithValue("apple", goci.RuntimeSpecMount{
				Type:    "apple_fs",
				Source:  "iDevice",
				Options: []string{"healthy", "shiny"},
			}))
		})
	})
})
