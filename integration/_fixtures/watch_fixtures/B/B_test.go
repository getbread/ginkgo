package B_test

import (
	. "github.com/getbread/ginkgo/integration/_fixtures/watch_fixtures/B"

	. "github.com/getbread/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
