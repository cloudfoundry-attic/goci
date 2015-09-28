package goci_test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/cloudfoundry-incubator/goci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Saving", func() {
	var tmp string

	BeforeEach(func() {
		var err error
		tmp, err = ioutil.TempDir("", "gocitest")
		Expect(err).NotTo(HaveOccurred())

		bndle := &goci.Bndle{
			Spec: goci.Spec{
				Version: "abcd",
			},
			RuntimeSpec: goci.RuntimeSpec{
				Mounts: map[string]goci.RuntimeSpecMount{
					"foo": goci.RuntimeSpecMount{},
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
})

func read(path string) []byte {
	b, err := ioutil.ReadAll(mustOpen(path))
	Expect(err).NotTo(HaveOccurred())

	return b
}

func mustOpen(path string) io.Reader {
	r, err := os.Open(path)
	Expect(err).NotTo(HaveOccurred())

	return r
}
