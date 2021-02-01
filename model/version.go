package model

import (
	"regexp"
	"strconv"
	"strings"
)

const separator = "."

// Version represents a version consisting of a major, minor and patch part
type Version struct {
	major versionPart
	minor versionPart
	patch versionPart
}

type versionPart struct {
	number    int
	isPresent bool
}

// NewVersion constructs a new version
func NewVersion(major int, minor int, patch int) Version {
	return Version{
		versionPart{major, true},
		versionPart{minor, true},
		versionPart{patch, true},
	}
}

// FromVersionString constructs a new version from a given version string
func FromVersionString(versionString string) (version Version, err error) {
	versionParts := strings.Split(versionString, separator)
	var major, minor, patch int

	if len(versionParts) > 0 {
		major, err = strconv.Atoi(versionParts[0])
		if err != nil {
			return
		}
		version.major = versionPart{major, true}
	}

	if len(versionParts) > 1 {
		minor, err = strconv.Atoi(versionParts[1])
		if err != nil {
			return
		}
		version.minor = versionPart{minor, true}
	}

	if len(versionParts) > 2 {
		patch, err = strconv.Atoi(versionParts[2])
		if err != nil {
			return
		}
		version.patch = versionPart{patch, true}
	}

	return
}

// ValidateVersionString checks if a given version string conforms to our version syntax
func ValidateVersionString(versionString string) bool {
	return regexp.MustCompile("^([0-9]+)(\\" + separator + "[0-9]+)?(\\" + separator + "[0-9]+)?$").MatchString(versionString)
}

// String returns the version's string representation
func (version Version) String() string {
	versionParts := make([]string, 0)

	if version.major.isPresent {
		versionParts = append(versionParts, strconv.Itoa(version.major.number))
	}

	if version.minor.isPresent {
		versionParts = append(versionParts, strconv.Itoa(version.minor.number))
	}

	if version.patch.isPresent {
		versionParts = append(versionParts, strconv.Itoa(version.patch.number))
	}

	return strings.Join(versionParts, separator)
}

// BumpMajor bumps the version's major part
func (version Version) BumpMajor() Version {
	return Version{
		versionPart{version.major.number + 1, true},
		versionPart{0, version.minor.isPresent},
		versionPart{0, version.patch.isPresent},
	}
}

// BumpMinor bumps the version's minor part
func (version Version) BumpMinor() Version {
	return Version{
		versionPart{version.major.number, true},
		versionPart{version.minor.number + 1, true},
		versionPart{0, version.patch.isPresent},
	}
}

// BumpPatch bumps the version's patch part
func (version Version) BumpPatch() Version {
	return Version{
		versionPart{version.major.number, true},
		versionPart{version.minor.number, true},
		versionPart{version.patch.number + 1, true},
	}
}
