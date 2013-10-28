package main_test

import (
	. "github.com/lsimoneau/rosalind"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Revp", func() {
  var (
    input string
  )

  BeforeEach(func() {
    input = ">Rosalind_24\nTCAATGCATGCGGGTCTATATGCAT"
  })

  Describe("Finding reverse complements", func() {
    It("returns the position and length of reverse complements between 4 and 12 bases", func() {
      Expect(Revp(input)).To(Equal("4 6\n5 4\n6 6\n7 4\n17 4\n18 4\n20 6\n21 4\n"))
    })
  })

})
