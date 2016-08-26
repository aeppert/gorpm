//
// Aaron Eppert - 2016
//
package main

import (
	"fmt"

	"github.com/aeppert/gorpm"
)

func dump_rpm_hdr(hdr *rpm.Header) {
	name, _ := hdr.GetString(rpm.RPMTAG_NAME)
	version, _ := hdr.GetString(rpm.RPMTAG_VERSION)
	release, _ := hdr.GetString(rpm.RPMTAG_RELEASE)
	arch, _ := hdr.GetString(rpm.RPMTAG_ARCH)
	summary, _ := hdr.GetString(rpm.RPMTAG_SUMMARY)
	fmt.Printf("%s-%s-%s.%s - %s\n", string(name), string(version),
		string(release), string(arch), string(summary))
}

func dump_all_rpms() {
	ts := rpm.RpmTsCreate()
	rpm.ReadConfigFiles(nil, nil)

	mi := ts.RpmTsInitIteratorSeq(rpm.RPMTAG_DBI_PACKAGES)

	hdr := mi.RpmDbNextIterator()

	for hdr != nil {
		hdr.Link()
		dump_rpm_hdr(hdr)
		hdr.Free()
		hdr = mi.RpmDbNextIterator()
	}

	mi.Free()
	ts.Free()
}

func dump_named_rpm(roi string) {
	ts := rpm.RpmTsCreate()
	rpm.ReadConfigFiles(nil, nil)

	mi := ts.RpmTsInitIteratorNamed(roi)
	hdr := mi.RpmDbNextIterator()
	for hdr != nil {
		hdr.Link()
		dump_rpm_hdr(hdr)
		hdr.Free()
		hdr = mi.RpmDbNextIterator()
	}

	mi.Free()
	ts.Free()
}

func main() {
	dump_all_rpms()
	dump_named_rpm("kernel-devel")
}
