/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 23.11.2013 17:18
 * description: Mostly 1:1 bindings to the functions defined in rpmts.h.
 */

package rpm

/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmlib.h>
#include <rpm/rpmts.h>

rpmdbMatchIterator rpmtsInitIterator_grapper(const rpmts ts, int rpmtag, const void * keyp, size_t keylen) {
    return rpmtsInitIterator(ts, rpmtag, keyp, keylen);
}

rpmdbMatchIterator rpmtsInitIterator_named_grapper(const rpmts ts, int rpmtag, char * keyp, size_t keylen) {
    return rpmtsInitIterator(ts, rpmtag, (const void *)keyp, keylen);
}

*/
import "C"
import (
	"unicode/utf8"
)

type RpmTs struct {
	c_ts C.rpmts
}

// RpmTsCreate (rpmtsCreate in RPM) creates an empty transaction set.
func RpmTsCreate() *RpmTs {
	return &RpmTs{c_ts: C.rpmtsCreate()}
}

// Free (rpmtsFree in RPM) destroys transaction set and closes the database.
func (ts *RpmTs) Free() {
	C.rpmtsFree(ts.c_ts)
}

// RpmTsInitIteratorSeq (rpmtsInitIterator in RPM) creates an interator over a transaction set
func (ts *RpmTs) RpmTsInitIteratorSeq(tag RpmTag) *RpmDbMatchIterator {

	crdmi := C.rpmtsInitIterator_grapper(ts.c_ts, C.int(tag), nil, 0)

	if crdmi == nil {
		return nil
	}

	return &RpmDbMatchIterator{c_rpmdbMatchIterator: crdmi}
}

// RpmTsInitIteratorNamed (rpmtsInitIterator in RPM) creates an interator over a transaction set
func (ts *RpmTs) RpmTsInitIteratorNamed(key string) *RpmDbMatchIterator {
	if utf8.RuneCountInString(key) > 0 {
		cstr := C.CString(key)
		crdmi := C.rpmtsInitIterator_named_grapper(ts.c_ts, C.int(rpm.RPMTAG_NAME), cstr, 0)

		if crdmi == nil {
			return nil
		}

		return &RpmDbMatchIterator{c_rpmdbMatchIterator: crdmi}
	}

	return nil
}
