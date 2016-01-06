package goci_test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/cloudfoundry-incubator/goci"
	"github.com/cloudfoundry-incubator/goci/specs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Saving", func() {
	var (
		tmp   string
		bndle *goci.Bndl
	)

	BeforeEach(func() {
		var err error
		tmp, err = ioutil.TempDir("", "gocitest")
		Expect(err).NotTo(HaveOccurred())

		bndle = &goci.Bndl{
			Spec: specs.LinuxSpec{
				Spec: specs.Spec{Version: "abcd"},
			},
			RuntimeSpec: specs.LinuxRuntimeSpec{
				RuntimeSpec: specs.RuntimeSpec{
					Mounts: map[string]specs.Mount{
						"foo": specs.Mount{},
					},
				},
			},
		}

		Expect(bndle.Save(tmp)).To(Succeed())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(tmp)).To(Succeed())
	})

	It("serializes the spec to spec.json", func() {
		var configJson map[string]interface{}
		Expect(json.NewDecoder(mustOpen(path.Join(tmp, "config.json"))).Decode(&configJson)).To(Succeed())

		Expect(configJson).To(HaveKeyWithValue("version", Equal("abcd")))
	})

	It("serializes the runtime spec to runtime.json", func() {
		Expect(path.Join(tmp, "runtime.json")).To(BeAnExistingFile())

		var runtimeJson map[string]interface{}
		Expect(json.NewDecoder(mustOpen(path.Join(tmp, "runtime.json"))).Decode(&runtimeJson)).To(Succeed())

		Expect(runtimeJson).To(HaveKeyWithValue("mounts", HaveKey("foo")))
	})

	Describe("Load", func() {
		Context("when config.json and runtime.json exist", func() {
			It("loads the bundle from runtime.json and config.json", func() {
				bundleLoader := &goci.BndlLoader{}
				loadedBundle, _ := bundleLoader.Load(tmp)
				Expect(loadedBundle).To(Equal(bndle))
			})
		})

		Context("when config.json does not exist", func() {
			It("returns an error", func() {
				Expect(os.Remove(path.Join(tmp, "config.json"))).To(Succeed())
				bundleLoader := &goci.BndlLoader{}
				_, err := bundleLoader.Load(tmp)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when runtime.json does not exist", func() {
			It("returns an error", func() {
				Expect(os.Remove(path.Join(tmp, "runtime.json"))).To(Succeed())
				bundleLoader := &goci.BndlLoader{}
				_, err := bundleLoader.Load(tmp)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func mustOpen(path string) io.Reader {
	r, err := os.Open(path)
	Expect(err).NotTo(HaveOccurred())

	return r
}
