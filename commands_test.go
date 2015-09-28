package goci_test

import (
	"github.com/cloudfoundry-incubator/goci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	BeforeEach(func() {
		goci.DefaultRuncBinary = "funC"
	})

	Describe("StartCommand", func() {
		It("creates an *exec.Cmd to start a bundle", func() {
			cmd := goci.StartCommand("my-bundle-path", "my-bundle-id")
			Expect(cmd.Args).To(Equal([]string{"funC", "--id", "my-bundle-id", "start", "my-bundle-path"}))
		})
	})

	Describe("ExecCommand", func() {
		It("creates an *exec.Cmd to exec a process in a bundle", func() {
			cmd := goci.ExecCommand("my-bundle-id", "my-process-json.json")
			Expect(cmd.Args).To(Equal([]string{"funC", "--id", "my-bundle-id", "exec", "my-process-json.json"}))
		})
	})
})
