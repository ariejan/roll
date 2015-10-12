package core

import "fmt"

type version struct {
	Major, Minor, Patch int
	Label               string
}

var Version = version{1, 2, 3, "dev"}

var Build string

func (v version) String() string {
	if v.Label != "" {
		return fmt.Sprintf("Roll version %d.%d.%d-%s\nGit commit hash: %s", v.Major, v.Minor, v.Patch, v.Label, Build)
	} else {
		return fmt.Sprintf("Roll version %d.%d.%d\nGit commit hash: %s", v.Major, v.Minor, v.Patch, Build)
	}
}
