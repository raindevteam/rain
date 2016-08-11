// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rain

// I took this code verbatim from: https://github.com/derekparker/delve/blob/master/version/version.go
// I thought it was a pretty neat way to handle versioning, as it makes everything pretty explicit.

// However, it wasn't exactly semantic, as "alpha" should be pre-release, not metadata. So I've changed that
// here appropriately.

// In their defense, I'm not sure if they are actually doing semver.

import "fmt"

// The Version struct is used to construct a version string.
type version struct {
	Major      string
	Minor      string
	Patch      string
	PreRelease string
	Build      string
	MetaData   string
}

// v is the current version of Rain.
var v = version{Major: "0", Minor: "7", Patch: "0", PreRelease: "", Build: "", MetaData: "LifeBloom"}

// Version returns the current version of Rain.
func Version() string {
	fmtstr := "%s.%s.%s"

	if v.PreRelease != "" {
		fmtstr = fmtstr + "-"
	}
	fmtstr = fmtstr + "%s"

	if v.Build != "" {
		fmtstr = fmtstr + "."
	}
	fmtstr = fmtstr + "%s"

	if v.MetaData != "" {
		fmtstr = fmtstr + "+"
	}
	fmtstr = fmtstr + "%s"

	return fmt.Sprintf(fmtstr, v.Major, v.Minor, v.Patch, v.PreRelease, v.Build, v.MetaData)
}
