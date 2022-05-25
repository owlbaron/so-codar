package grouper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGrouper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grouper Suite")
}
