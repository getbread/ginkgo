package A_test

import (
	. "github.com/getbread/ginkgo/integration/_fixtures/watch_fixtures/A"

	. "github.com/getbread/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
