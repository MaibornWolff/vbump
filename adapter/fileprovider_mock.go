package adapter

import (
	"maibornwolff/vbump/model"
)

// FileProviderMock for testing
type FileProviderMock struct {
	version       model.Version
	project       string
	VersionStored bool
}

// NewMock constructs a new mock for a file provider
func NewMock(version model.Version, project string) StorageProvider {
	return &FileProviderMock{
		version: version,
		project: project,
	}
}

// ReadVersion returns the current version from FileProviderMock
func (provider *FileProviderMock) ReadVersion(project string) (model.Version, error) {
	if provider.project == project {
		return provider.version, nil
	}

	return model.Version{}, nil
}

// StoreVersion sets VersionStored to true
func (provider *FileProviderMock) StoreVersion(string, model.Version) error {
	provider.VersionStored = true
	return nil
}
