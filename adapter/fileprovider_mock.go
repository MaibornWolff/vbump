package adapter

type FileProviderMock struct {
	version string
	project string
}

func NewMock(version string, project string) IFileProvider {
	return &FileProviderMock{
		version: version,
		project: project,
	}
}

func (provider *FileProviderMock) ReadVersion(project string) (string, error) {
	if provider.project == project {
		return provider.version, nil
	}

	return "", nil
}

func (provider *FileProviderMock) StoreVersion(project string, version string) error {
	return nil
}
