package util

import (
	"strings"
)

type SplitInclude int

const (
	NEVER SplitInclude = iota
	BEFORE
	AFTER
)

func SplitAt(full string, stub string, includeRemoved SplitInclude) (string, string) {
	s := strings.Split(full, stub)
	start, end := s[0], s[1]

	switch includeRemoved {
	case 0:
		return start, end
	case 1:
		return start + stub, end
	case 2:
		return start, stub + end
	}

	return "", ""
}

func ExpandPath(path string) (string, string) {
	if strings.Contains(path, ":") {
		namespace, parent := SplitAt(path, ":", NEVER)

		return namespace, parent
	}

	return "minecraft", path
}

func SExpandPath(path string) string {
	namespace, parent := ExpandPath(path)
	return namespace + "/models/" + parent
}

func SExpandModel(namespace string, parent string) string {
	return namespace + "/models/" + parent
}
