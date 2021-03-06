/* -*- mode:go; coding:utf-8; -*-
 * author: Aaron Eppert <aeppert@gmail.com>
 * created: 2016-08-26
 * description: Attempted 1:1 bindings for rpmdb.h
 */

package rpm

/*
#cgo LDFLAGS: -lrpm -lrpmio
#include <rpm/rpmtypes.h>
#include <rpm/rpmsw.h>
#include <rpm/rpmdb.h>
#include <stdlib.h>
*/
import "C"

type RpmMireMode uint32

const (
	RPMMIRE_DEFAULT = RpmMireMode(C.RPMMIRE_DEFAULT) // !< regex with \., .* and ^...$ added
	RPMMIRE_STRCMP  = RpmMireMode(C.RPMMIRE_STRCMP)  // !< strings  using strcmp(3) */
	RPMMIRE_REGEX   = RpmMireMode(C.RPMMIRE_REGEX)   // !< regex(7) patterns through regcomp(3)
	RPMMIRE_GLOB    = RpmMireMode(C.RPMMIRE_GLOB)    // !< glob(7) patterns through fnmatch(3)
)

type RpmDbOpX uint32

const (
	RPMDB_OP_DBGET = RpmDbOpX(C.RPMDB_OP_DBGET)
	RPMDB_OP_DBPUT = RpmDbOpX(C.RPMDB_OP_DBPUT)
	RPMDB_OP_DBDEL = RpmDbOpX(C.RPMDB_OP_DBDEL)
	RPMDB_OP_MAX   = RpmDbOpX(C.RPMDB_OP_MAX)
)

type RpmDbMatchIterator struct {
	c_rpmdbMatchIterator C.rpmdbMatchIterator
}

// RpmDbNextIterator (rpmdbNextIterator in C)
func (mi *RpmDbMatchIterator) RpmDbNextIterator() *Header {
	c_header := C.rpmdbNextIterator(mi.c_rpmdbMatchIterator)
	if c_header == nil {
		return nil
	}

	return &Header{c_header: c_header}
}

// Free (rpmdbFreeIterator in C)
func (mi *RpmDbMatchIterator) Free() {
	C.rpmdbFreeIterator(mi.c_rpmdbMatchIterator)
}
