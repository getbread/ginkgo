package writer_test

import (
	. "github.com/getbread/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Writer Suite")
}
