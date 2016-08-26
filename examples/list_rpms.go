//
// Aaron Eppert - 2016
//
package main

import (
	"fmt"

	"github.com/aeppert/gorpm"
)

func main() {
	ts := rpm.RpmTsCreate()
	rpm.ReadConfigFiles(nil, nil)

	mi := ts.RpmTsInitIterator(rpm.RPMTAG_DBI_PACKAGES)

	hdr := mi.RpmDbNextIterator()
	for hdr != nil {
		hdr.Link()
		name, _ := hdr.GetString(rpm.RPMTAG_NAME)
		version, _ := hdr.GetString(rpm.RPMTAG_VERSION)
		release, _ := hdr.GetString(rpm.RPMTAG_RELEASE)
		arch, _ := hdr.GetString(rpm.RPMTAG_ARCH)
		summary, _ := hdr.GetString(rpm.RPMTAG_SUMMARY)

		fmt.Printf("%s-%s-%s.%s - %s\n", string(name), string(version),
			string(release), string(arch), string(summary))
		hdr.Free()
		hdr = mi.RpmDbNextIterator()
	}

	mi.Free()
	ts.Free()
}
