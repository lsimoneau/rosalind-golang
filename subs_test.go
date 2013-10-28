package main_test

import (
	. "github.com/lsimoneau/rosalind"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subs", func() {
  It("Finds the locations of each occurrence of the pattern", func() {
    haystack := "GATATATGCATATACTT"
    needle   := "ATAT"

    Expect(Subs(needle, haystack)).To(Equal([]int{2,4,10}))
  })

})
