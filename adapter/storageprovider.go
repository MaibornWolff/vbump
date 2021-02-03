package adapter

import (
	"maibornwolff/vbump/model"
)

// StorageProvider allows to read and write version data from/to a storage
type StorageProvider interface {
	ReadVersion(project string) (model.Version, error)
	StoreVersion(project string, version model.Version) error
}
