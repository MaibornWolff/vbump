package adapter

// FileProviderMock for testing
type FileProviderMock struct {
	version       string
	project       string
	VersionStored bool
}

// NewMock constructs a new FileProvider Mock
func NewMock(version string, project string) IFileProvider {
	return &FileProviderMock{
		version: version,
		project: project,
	}
}

// ReadVersion return the current version from providermock
func (provider *FileProviderMock) ReadVersion(project string) (string, error) {
	if provider.project == project {
		return provider.version, nil
	}

	return "", nil
}

//StoreVersion sets versionStored to true, if this funcion is called
func (provider *FileProviderMock) StoreVersion(project string, version string) error {
	provider.VersionStored = true
	return nil
}
