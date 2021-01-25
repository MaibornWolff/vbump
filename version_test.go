package main

import (
	"testing"

	"maibornwolff/vbump/adapter"

	. "github.com/onsi/gomega"
)

func Test_Bumb_Major_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("2.0.0"))
}

func Test_Bumb_Major_Version_Without_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("2.0"))
}

func Test_Bumb_Major_Version_Without_Minor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("2"))
}

func Test_Bumb_Minor_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMinor("1")

	Ω.Expect(actual).To(Equal("1.1.0"))
}

func Test_Bumb_Minor_Version_Without_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMinor("1")

	Ω.Expect(actual).To(Equal("1.1"))
}

func Test_Bumb_Patch_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpPatch("1")

	Ω.Expect(actual).To(Equal("1.0.1"))
}

func Test_Bumb_Patch_On_Major_Minor_Without_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpPatch("1")

	Ω.Expect(actual).To(Equal("1.0.1"))
}

func Test_Bumb_Patch_On_Major_Without_Minor_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpPatch("1")

	Ω.Expect(actual).To(Equal("1.0.1"))
}

func Test_Bumb_Minor_On_Major_Without_Minor_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMinor("1")

	Ω.Expect(actual).To(Equal("1.1"))
}

func Test_Bumb_Major_And_Reset_Minor_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.1.1", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("2.0.0"))
}

func Test_Bumb_Major_And_Reset_Minor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.1.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("2.0.0"))
}

func Test_Bumb_Major_And_Reset_Minor_Without_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.1", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("2.0"))
}

func Test_Bumb_Minor_And_Reset_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.1.1", "1")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMinor("1")

	Ω.Expect(actual).To(Equal("1.2.0"))
}

func Test_Bumb_Major_On_Empty_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("", "")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMajor("1")

	Ω.Expect(actual).To(Equal("1"))
}

func Test_Bumb_Minor_On_Empty_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("", "")
	version := NewVersion(providerMock)
	actual, _ := version.BumpMinor("1")

	Ω.Expect(actual).To(Equal("0.1"))
}

func Test_Bumb_Patch_On_Empty_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("", "")
	version := NewVersion(providerMock)
	actual, _ := version.BumpPatch("1")

	Ω.Expect(actual).To(Equal("0.0.1"))
}

func Test_Set_Version_Explicit_With_Major_Minor_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1.0.0")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", "1.0.0")

	Ω.Expect(err).To(BeNil())
}

func Test_Set_Version_Explicit_With_Major_Minor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1.0")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", "1.0")

	Ω.Expect(err).To(BeNil())
}

func Test_Set_Version_Explicit_With_Major(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", "1")

	Ω.Expect(err).To(BeNil())
}

func Test_Set_Version_Explicit_With_Invalid_Version_As_Text(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "aaa")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", "aaa")

	Ω.Expect(err).NotTo(BeNil())
}

func Test_Set_Version_Explicit_With_Invalid_Version_With_Only_Dot(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "1.")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", "1.")

	Ω.Expect(err).NotTo(BeNil())
}

func Test_Set_Version_Explicit_With_Invalid_Version_With_Dot_In_Front(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", ".1")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", ".1")

	Ω.Expect(err).NotTo(BeNil())
}

func Test_Set_Version_Explicit_With_Invalid_Version_With_surrounding_Chars(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1", "aaa1.1.1aaa")
	version := NewVersion(providerMock)
	_, err := version.SetVersion("1", "aaa1.1.1aaa")

	Ω.Expect(err).NotTo(BeNil())
}

func Test_Get_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)

	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)
	actual, _ := version.GetVersion("1")

	Ω.Expect(actual).To(Equal("1.0"))
}

func Test_Bump_Transient_Patch_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)
	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)

	actual, _ := version.BumpTransientPatch("1.0")

	Ω.Expect(actual).To(Equal("1.0.1"))
	Ω.Expect(providerMock.(*adapter.FileProviderMock).VersionStored).To(Equal(false))
}

func Test_Bump_Transient_Minor_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)
	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)

	actual, _ := version.BumpTransientMinor("1.0")

	Ω.Expect(actual).To(Equal("1.1"))
	Ω.Expect(providerMock.(*adapter.FileProviderMock).VersionStored).To(Equal(false))
}

func Test_Bump_Transient_Patch_Version_With_Invalid_Given_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)
	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)

	_, err := version.BumpTransientPatch("a1.0a")

	Ω.Expect(err).NotTo(BeNil())
}

func Test_Bump_Transient_Minor_Version_With_Invalid_Given_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)
	providerMock := adapter.NewMock("1.0", "1")
	version := NewVersion(providerMock)

	_, err := version.BumpTransientMinor("a1.0a")

	Ω.Expect(err).NotTo(BeNil())
}
