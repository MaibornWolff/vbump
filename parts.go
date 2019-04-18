package main

import "strings"

func concatVersionPart(part string, withDot bool) string {
	version := ""
	if part == "0" || part != "" {
		if withDot {
			version = "."
		}
		version += part
	}

	return version
}

func resetPart(part string) string {
	if part != "" {
		return "0"
	}

	return ""
}

func initEmptyPartToZero(part string) string {
	if part == "" {
		return "0"
	}

	return part
}

func extractVersionParts(version string) (string, string, string) {
	major, minor, patch := "", "", ""

	parts := strings.Split(version, ".")
	major, parts = parts[0], parts[1:]

	if len(parts) > 0 {
		minor, parts = parts[0], parts[1:]
	}

	if len(parts) > 0 {
		patch, parts = parts[0], parts[1:]
	}

	return major, minor, patch
}
