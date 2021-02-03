package model

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestBumbMajorVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.0.0")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("2.0.0"))
}

func TestBumbMajorVersionWithoutPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.0")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("2.0"))
}

func TestBumbMajorVersionWithoutMinor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("2"))
}

func TestBumbMinorVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.0.0")
	actual := version.BumpMinor()

	Ω.Expect(actual.String()).To(Equal("1.1.0"))
}

func TestBumbMinorVersionWithoutPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.0")
	actual := version.BumpMinor()

	Ω.Expect(actual.String()).To(Equal("1.1"))
}

func TestBumbPatchVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.0.0")
	actual := version.BumpPatch()

	Ω.Expect(actual.String()).To(Equal("1.0.1"))
}

func TestBumbPatchOnMajorMinorWithoutPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.0")
	actual := version.BumpPatch()

	Ω.Expect(actual.String()).To(Equal("1.0.1"))
}

func TestBumbPatchOnMajorWithoutMinorPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1")
	actual := version.BumpPatch()

	Ω.Expect(actual.String()).To(Equal("1.0.1"))
}

func TestBumbMinorOnMajorWithoutMinorPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1")
	actual := version.BumpMinor()

	Ω.Expect(actual.String()).To(Equal("1.1"))
}

func TestBumbMajorAndResetMinorPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.1.1")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("2.0.0"))
}

func TestBumbMajorAndResetMinor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.1.0")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("2.0.0"))
}

func TestBumbMajorAndResetMinorWithoutPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.1")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("2.0"))
}

func TestBumbMinorAndResetPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("1.1.1")
	actual := version.BumpMinor()

	Ω.Expect(actual.String()).To(Equal("1.2.0"))
}

func TestBumbMajorOnEmptyVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("")
	actual := version.BumpMajor()

	Ω.Expect(actual.String()).To(Equal("1"))
}

func TestBumbMinorOnEmptyVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("")
	actual := version.BumpMinor()

	Ω.Expect(actual.String()).To(Equal("0.1"))
}

func TestBumbPatchOnEmptyVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := FromVersionString("")
	actual := version.BumpPatch()

	Ω.Expect(actual.String()).To(Equal("0.0.1"))
}
