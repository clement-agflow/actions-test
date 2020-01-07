package version

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func New(s string) (Version, error) {
	if !regexp.MustCompile(`^\d{1,2}\.\d{1,2}\.\d{1,2}$`).MatchString(s) {
		return Version{}, errors.New("invalid session string")
	}

	v := strings.Split(string(s), ".")
	major, _ := strconv.Atoi(v[0])
	minor, _ := strconv.Atoi(v[1])
	patch, _ := strconv.Atoi(v[2])

	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
