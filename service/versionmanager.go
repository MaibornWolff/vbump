package service

import (
	"github.com/pkg/errors"
	"maibornwolff/vbump/adapter"
	"maibornwolff/vbump/model"
)

// VersionManager bumps major, minor, patch part of a given project
type VersionManager struct {
	storageProvider adapter.StorageProvider
}

// NewVersionManager constructs a new version manager
func NewVersionManager(provider adapter.StorageProvider) *VersionManager {
	return &VersionManager{
		storageProvider: provider,
	}
}

// BumpMajor bumps major version for given project
func (vm *VersionManager) BumpMajor(project string) (model.Version, error) {
	currentVersion, err := vm.storageProvider.ReadVersion(project)
	if err != nil {
		return currentVersion, err
	}

	newVersion := currentVersion.BumpMajor()

	err = vm.storageProvider.StoreVersion(project, newVersion)
	if err != nil {
		return currentVersion, err
	}

	return newVersion, nil
}

// BumpMinor bumps minor version for given project
func (vm *VersionManager) BumpMinor(project string) (model.Version, error) {
	currentVersion, err := vm.storageProvider.ReadVersion(project)
	if err != nil {
		return currentVersion, err
	}

	newVersion := currentVersion.BumpMinor()

	err = vm.storageProvider.StoreVersion(project, newVersion)
	if err != nil {
		return currentVersion, err
	}

	return newVersion, nil
}

// BumpPatch bumps patch version for given project
func (vm *VersionManager) BumpPatch(project string) (model.Version, error) {
	currentVersion, err := vm.storageProvider.ReadVersion(project)
	if err != nil {
		return currentVersion, err
	}

	newVersion := currentVersion.BumpPatch()

	err = vm.storageProvider.StoreVersion(project, newVersion)
	if err != nil {
		return currentVersion, err
	}

	return newVersion, err
}

// SetVersion sets the current given version for the given project
func (vm *VersionManager) SetVersion(project string, versionString string) (model.Version, error) {
	isValidated := model.ValidateVersionString(versionString)
	if !isValidated {
		return model.Version{}, errors.Errorf("%v is not a valid version", versionString)
	}

	version, err := model.FromVersionString(versionString)
	if err != nil {
		return version, errors.Wrapf(err, "Failed to convert version %v", versionString)
	}

	err = vm.storageProvider.StoreVersion(project, version)
	if err != nil {
		return version, errors.Wrapf(err, "Failed to set version %v for project %v", versionString, project)
	}

	return version, err
}

// GetVersion returns current version for given project
func (vm *VersionManager) GetVersion(project string) (model.Version, error) {
	version, err := vm.storageProvider.ReadVersion(project)
	if err != nil {
		return version, errors.Wrapf(err, "Failed to get version for project %v", project)
	}

	return version, nil
}

// BumpTransientPatch bumps only the patch part on given version without change any project
func (vm *VersionManager) BumpTransientPatch(versionString string) (model.Version, error) {
	isValidated := model.ValidateVersionString(versionString)
	if !isValidated {
		return model.Version{}, errors.Errorf("%v is not a valid version", versionString)
	}

	version, err := model.FromVersionString(versionString)
	if err != nil {
		return version, errors.Wrapf(err, "Failed to convert version %v", versionString)
	}

	newVersion := version.BumpPatch()

	return newVersion, nil
}

// BumpTransientMinor bumps only the minor part on given version without change any project
func (vm *VersionManager) BumpTransientMinor(versionString string) (model.Version, error) {
	isValidated := model.ValidateVersionString(versionString)
	if !isValidated {
		return model.Version{}, errors.Errorf("%v is not a valid version", versionString)
	}

	version, err := model.FromVersionString(versionString)
	if err != nil {
		return model.Version{}, errors.Wrapf(err, "Failed to convert version %v", versionString)
	}

	newVersion := version.BumpMinor()

	return newVersion, nil
}
