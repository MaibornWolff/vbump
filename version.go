package main

import (
	"regexp"
	"strconv"

	"github.com/maibornwolff/vbump/adapter"
	"github.com/pkg/errors"
)

type Version struct {
	fileProvider adapter.IFileProvider
}

func NewVersion(provider adapter.IFileProvider) *Version {
	return &Version{
		fileProvider: provider,
	}
}

func (v *Version) BumpMajor(project string) (string, error) {
	currentVersion, err := v.fileProvider.ReadVersion(project)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot bump major version on project %v", project)
	}

	major, minor, patch := extractVersionParts(currentVersion)
	newMajor := convertAndInc(major)
	minor = resetPart(minor)
	patch = resetPart(patch)
	newVersion := formatVersion(newMajor, minor, patch)
	v.fileProvider.StoreVersion(project, newVersion)

	return newVersion, nil
}

func (v *Version) BumpMinor(project string) (string, error) {
	currentVersion, err := v.fileProvider.ReadVersion(project)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot bump minor version on project %v", project)
	}

	major, minor, patch := extractVersionParts(currentVersion)
	newMinor := convertAndInc(minor)
	major = initEmptyPartToZero(major)
	patch = resetPart(patch)
	newVersion := formatVersion(major, newMinor, patch)
	err = v.fileProvider.StoreVersion(project, newVersion)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot bump minor version on project %v", project)
	}

	return newVersion, nil
}

func (v *Version) BumpPatch(project string) (string, error) {
	currentVersion, err := v.fileProvider.ReadVersion(project)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot bump patch version on project %v", project)
	}

	major, minor, patch := extractVersionParts(currentVersion)
	newPatch := convertAndInc(patch)
	major = initEmptyPartToZero(major)
	minor = initEmptyPartToZero(minor)
	newVersion := formatVersion(major, minor, newPatch)
	v.fileProvider.StoreVersion(project, newVersion)
	err = v.fileProvider.StoreVersion(project, newVersion)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot bump patch version on project %v", project)
	}

	return newVersion, err
}

func (v *Version) SetVersion(project string, version string) (string, error) {
	isValidated := validateVersion(version)
	if !isValidated {
		return "", errors.Errorf("%v is not a valid version", version)
	}

	err := v.fileProvider.StoreVersion(project, version)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot set version %v for project %v", version, project)
	}

	return version, err
}

func (v *Version) GetVersion(project string) (string, error) {
	version, err := v.fileProvider.ReadVersion(project)
	if err != nil {
		return "", errors.Wrapf(err, "Cannot get version for project %v", project)
	}

	return version, err
}

func validateVersion(version string) bool {
	ex3 := regexp.MustCompile("^([0-9]+)\\.([0-9]+)\\.([0-9]+)$")
	ex2 := regexp.MustCompile("^([0-9]+)\\.([0-9]+)$")
	ex1 := regexp.MustCompile("^([0-9]+)$")
	return ex3.MatchString(version) || ex2.MatchString(version) || ex1.MatchString(version)
}

func convertAndInc(version string) string {
	versionToInc, _ := strconv.Atoi(version)
	newVersion := strconv.Itoa(versionToInc + 1)

	return newVersion
}

func formatVersion(major string, minor string, patch string) string {
	version := concatVersionPart(major, false)
	version += concatVersionPart(minor, true)
	version += concatVersionPart(patch, true)

	return version
}
