package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRosalind(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rosalind Suite")
}
