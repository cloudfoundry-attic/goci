package goci_test

import (
	"github.com/cloudfoundry-incubator/goci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	BeforeEach(func() {
		goci.DefaultRuncBinary = goci.NewBinary("funC")
	})

	Describe("StartCommand", func() {
		It("creates an *exec.Cmd to start a bundle", func() {
			cmd := goci.WithLogFile("mylog.file").StartCommand("my-bundle-path", "my-bundle-id", false)
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "start", "my-bundle-id"}))
			Expect(cmd.Dir).To(Equal("my-bundle-path"))
		})

		It("turns on debug logging", func() {
			cmd := goci.WithLogFile("mylog.file").StartCommand("my-bundle-path", "my-bundle-id", true)
			Expect(cmd.Args).To(ContainElement("--debug"))
		})

		It("passes the detach flag if requested", func() {
			cmd := goci.StartCommand("my-bundle-path", "my-bundle-id", true)
			Expect(cmd.Args).To(ContainElement("-d"))
		})
	})

	Describe("ExecCommand", func() {
		It("creates an *exec.Cmd to exec a process in a bundle", func() {
			cmd := goci.WithLogFile("mylog.file").ExecCommand("my-bundle-id", "my-process-json.json", "some-pid-file")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "exec", "my-bundle-id", "--pid-file", "some-pid-file", "-p", "my-process-json.json"}))
		})
	})

	Describe("EventsCommand", func() {
		It("creates an *exec.Cmd to get events for a bundle", func() {
			cmd := goci.WithLogFile("mylog.file").EventsCommand("my-bundle-id")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "events", "my-bundle-id"}))
		})
	})

	Describe("KillCommand", func() {
		It("creates an *exec.Cmd to signal the bundle", func() {
			cmd := goci.WithLogFile("mylog.file").KillCommand("my-bundle-id", "TERM")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "kill", "my-bundle-id", "TERM"}))
		})
	})

	Describe("StateCommand", func() {
		It("creates an *exec.Cmd to get the state of the bundle", func() {
			cmd := goci.WithLogFile("mylog.file").StateCommand("my-bundle-id")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "state", "my-bundle-id"}))
		})
	})

	Describe("StatsCommand", func() {
		It("creates an *exec.Cmd to get the state of the bundle", func() {
			cmd := goci.WithLogFile("mylog.file").StatsCommand("my-bundle-id")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "events", "--stats", "my-bundle-id"}))
		})
	})

	Describe("DeleteCommand", func() {
		It("creates an *exec.Cmd to delete the bundle", func() {
			cmd := goci.WithLogFile("mylog.file").DeleteCommand("my-bundle-id")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "delete", "my-bundle-id"}))
		})
	})
})
