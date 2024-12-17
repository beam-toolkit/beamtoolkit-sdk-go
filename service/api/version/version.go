package version

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

const Version = "0.0.2"

const Prerelease = ""

var SemVar *version.Version

func init() {
	SemVar = version.Must(version.NewVersion(Version))
}

func String() string {
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s", Version, Prerelease)
	}

	return Version
}
