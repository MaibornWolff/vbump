package service

import (
	"testing"

	. "github.com/onsi/gomega"
	"maibornwolff/vbump/adapter"
	"maibornwolff/vbump/model"
)

func TestBumbMajorVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	actual, _ := versionManager.BumpMajor("A")

	Ω.Expect(actual.String()).To(Equal("2.0.0"))
}

func TestBumbMinorVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	actual, _ := versionManager.BumpMinor("A")

	Ω.Expect(actual.String()).To(Equal("1.1.0"))
}

func TestBumbPatchVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	actual, _ := versionManager.BumpPatch("A")

	Ω.Expect(actual.String()).To(Equal("1.0.1"))
}

func TestSetVersionExplicitWithMajorMinorPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", "1.0.0")

	Ω.Expect(err).To(BeNil())
}

func TestSetVersionExplicitWithMajorMinor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", "1.0")

	Ω.Expect(err).To(BeNil())
}

func TestSetVersionExplicitWithMajor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", "1")

	Ω.Expect(err).To(BeNil())
}

func TestSetVersionExplicitWithInvalidVersionAsText(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", "aaa")

	Ω.Expect(err).NotTo(BeNil())
}

func TestSetVersionExplicitWithInvalidVersionWithOnlyDot(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", "1.")

	Ω.Expect(err).NotTo(BeNil())
}

func TestSetVersionExplicitWithInvalidVersionWithDotInFront(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", ".1")

	Ω.Expect(err).NotTo(BeNil())
}

func TestSetVersionExplicitWithInvalidVersionWithSurroundingChars(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.SetVersion("A", "aaa1.1.1aaa")

	Ω.Expect(err).NotTo(BeNil())
}

func TestGetVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	actual, _ := versionManager.GetVersion("A")

	Ω.Expect(actual.String()).To(Equal("1.0"))
}

func TestBumpTransientPatchVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	actual, _ := versionManager.BumpTransientPatch("1.0")

	Ω.Expect(actual.String()).To(Equal("1.0.1"))
	Ω.Expect(providerMock.(*adapter.FileProviderMock).VersionStored).To(Equal(false))
}

func TestBumpTransientMinorVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	actual, _ := versionManager.BumpTransientMinor("1.0")

	Ω.Expect(actual.String()).To(Equal("1.1"))
	Ω.Expect(providerMock.(*adapter.FileProviderMock).VersionStored).To(Equal(false))
}

func TestBumpTransientPatchVersionWithInvalidGivenVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.BumpTransientPatch("a1.0a")

	Ω.Expect(err).NotTo(BeNil())
}

func TestBumpTransientMinorVersionWithInvalidGivenVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	version, _ := model.FromVersionString("1.0")
	providerMock := adapter.NewMock(version, "A")
	versionManager := NewVersionManager(providerMock)
	_, err := versionManager.BumpTransientMinor("a1.0a")

	Ω.Expect(err).NotTo(BeNil())
}
