package testsuite_test

import (
	. "github.com/getbread/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestsuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsuite Suite")
}
