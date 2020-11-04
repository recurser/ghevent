package ghevent

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filter Suite")
}

var _ = Describe("ghevent", func() {
	XIt("should test something", func() {
		// Todo.
	})
})
