package adapter

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

type IFileProvider interface {
	ReadVersion(project string) (string, error)
	StoreVersion(project string, version string) error
}

type FileProvider struct {
	basePath string
}

func New(basePath string) IFileProvider {
	return &FileProvider{basePath: basePath}
}

func (provider *FileProvider) ReadVersion(project string) (string, error) {
	filename := path.Join(provider.basePath, project)

	if _, err := os.Stat(provider.basePath); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "Basedirectory %v not exist", provider.basePath)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", nil
	}

	versiontext, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrapf(err, "Read version from file %v failed", filename)
	}

	return string(versiontext[:]), nil
}

func (provider *FileProvider) StoreVersion(project string, version string) error {
	text := []byte(version)
	filename := path.Join(provider.basePath, project)
	err := ioutil.WriteFile(filename, text, 0644)
	if err != nil {
		return errors.Wrap(err, "Store version in file failed")
	}

	return nil
}
