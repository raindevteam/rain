package rain

// I took this code verbatim from: https://github.com/derekparker/delve/blob/master/version/version.go
// I thought it was a pretty neat way to handle versioning, as it makes everything pretty explicit.

// However, it wasn't exactly semantic, as "alpha" should be pre-release, not metadata. So I've changed that
// here appropriately.

// In their defense, I'm not sure if they are actually doing semver.

import "fmt"

// The Version struct is used to construct a version string.
type Version struct {
	Major      string
	Minor      string
	Patch      string
	PreRelease string
	Build      string
	MetaData   string
}

// rainVersion is the current version of Rain.
var rainVersion = Version{Major: "0", Minor: "7", Patch: "0", PreRelease: "alpha", Build: "1", MetaData: "LifeBloom"}

// Version returns a formated version string.
func (v Version) Version() string {
	return fmt.Sprintf("Version: %s.%s.%s-%s.%s+%s", v.Major, v.Minor, v.Patch, v.PreRelease, v.Build, v.MetaData)
}
