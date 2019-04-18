package adapter

import (
	"log"
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_Store_And_Read_File(t *testing.T) {
	Ω := NewGomegaWithT(t)
	filename := "store_project"
	removeFile(filename)
	provider := New("../adapter")

	provider.StoreVersion(filename, "1.0.0")
	actual, _ := provider.ReadVersion(filename)

	Ω.Expect(actual).To(Equal("1.0.0"))
}

func Test_Read_Without_Store_File(t *testing.T) {
	Ω := NewGomegaWithT(t)
	filename := "store_project"
	removeFile(filename)
	provider := New("../adapter")

	actual, _ := provider.ReadVersion(filename)

	Ω.Expect(actual).To(Equal(""))
}

func Test_Correct_Version_Scope(t *testing.T) {
	Ω := NewGomegaWithT(t)
	filename1 := "store_project1"
	filename2 := "store_project2"
	removeFile(filename1)
	removeFile(filename2)
	provider := New("../adapter")

	provider.StoreVersion(filename1, "1.0")
	provider.StoreVersion(filename2, "2.0")
	actual1, _ := provider.ReadVersion(filename1)
	actual2, _ := provider.ReadVersion(filename2)

	Ω.Expect(actual1).To(Equal("1.0"))
	Ω.Expect(actual2).To(Equal("2.0"))
}

func removeFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		var err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}
