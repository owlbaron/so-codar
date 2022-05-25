package nl_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nl Suite")
}
