package config

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	GitCommit   string
	Version     = "1.4.1-unity"
	VersionDesc string
)

func init() {
	gitCommit := ""
	if len(GitCommit) > 0 {
		gitCommit = "; git " + GitCommit
	}
	VersionDesc = fmt.Sprintf("git-lfs/%s (GitHub; %s %s; go %s%s; Unity Collaborate Enabled)",
		Version,
		runtime.GOOS,
		runtime.GOARCH,
		strings.Replace(runtime.Version(), "go", "", 1),
		gitCommit,
	)

}
