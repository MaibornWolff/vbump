package adapter

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
	"maibornwolff/vbump/model"
)

// FileProvider reads and writes version data from/to files
type FileProvider struct {
	basePath string
}

// NewFileProvider constructs a new file provider
func NewFileProvider(basePath string) StorageProvider {
	return &FileProvider{basePath: basePath}
}

// ReadVersion reads the given project's version from a file
func (provider *FileProvider) ReadVersion(project string) (version model.Version, err error) {
	filename := path.Join(provider.basePath, project)

	if _, err = os.Stat(provider.basePath); os.IsNotExist(err) {
		return version, errors.Wrapf(err, "Base directory %v does not exist", provider.basePath)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return version, errors.Wrapf(err, "File %v does not exist", filename)
	}

	versionData, err := ioutil.ReadFile(filename)
	if err != nil {
		return version, errors.Wrapf(err, "Failed to read version from file %v", filename)
	}

	version, err = model.FromVersionString(string(versionData))
	if err != nil {
		return version, errors.Wrapf(err, "Failed to convert version")
	}

	return
}

// StoreVersion writes the given project's version to a file
func (provider *FileProvider) StoreVersion(project string, version model.Version) error {
	filename := path.Join(provider.basePath, project)

	versionString := version.String()
	versionBytes := []byte(versionString)

	err := ioutil.WriteFile(filename, versionBytes, 0644)
	if err != nil {
		return errors.Wrap(err, "Failed to store version in file")
	}

	return nil
}
