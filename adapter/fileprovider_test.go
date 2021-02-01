package adapter

import (
	"log"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"maibornwolff/vbump/model"
)

func TestStoreAndReadFile(t *testing.T) {
	Ω := NewGomegaWithT(t)

	filename := "store_project"
	removeFile(filename)
	provider := NewFileProvider("../adapter")

	_ = provider.StoreVersion(filename, model.NewVersion(1, 0, 0))
	actual, _ := provider.ReadVersion(filename)

	removeFile(filename)

	Ω.Expect(actual).To(Equal(model.NewVersion(1, 0, 0)))
}

func TestReadWithoutStoreFile(t *testing.T) {
	Ω := NewGomegaWithT(t)

	filename := "store_project"
	removeFile(filename)
	provider := NewFileProvider("../adapter")

	actual, _ := provider.ReadVersion(filename)

	Ω.Expect(actual.String()).To(Equal(""))
}

func TestCorrectVersionScope(t *testing.T) {
	Ω := NewGomegaWithT(t)

	filename1 := "store_project1"
	filename2 := "store_project2"
	removeFile(filename1)
	removeFile(filename2)
	provider := NewFileProvider("../adapter")

	_ = provider.StoreVersion(filename1, model.NewVersion(1, 0, 0))
	_ = provider.StoreVersion(filename2, model.NewVersion(2, 0, 0))
	actual1, _ := provider.ReadVersion(filename1)
	actual2, _ := provider.ReadVersion(filename2)

	removeFile(filename1)
	removeFile(filename2)

	Ω.Expect(actual1).To(Equal(model.NewVersion(1, 0, 0)))
	Ω.Expect(actual2).To(Equal(model.NewVersion(2, 0, 0)))
}

func removeFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		var err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}
